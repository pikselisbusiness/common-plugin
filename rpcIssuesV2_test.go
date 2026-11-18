package shared

import (
	"encoding/json"
	"errors"
	"testing"

	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

// issuesV2Host is a host exposing v2 Issues and Extensions.
type issuesV2Host struct {
	API
	organizationId    uint
	issueId, statusId uint
	entityType        string
	entityId          uint
	data              any
}

func (h *issuesV2Host) V2() v2.API { return issuesV2{host: h} }

type issuesV2 struct {
	v2.API
	host *issuesV2Host
}

func (s issuesV2) Issues() v2.IssuesAPI         { return issuesV2Issues{host: s.host} }
func (s issuesV2) Extensions() v2.ExtensionsAPI { return issuesV2Extensions{host: s.host} }

type issuesV2Issues struct{ host *issuesV2Host }

func (s issuesV2Issues) SetIssueStatus(context RequestContext, issueId uint, statusId uint) error {
	if statusId == 4 {
		return errors.New("a plugin cannot close an issue")
	}
	s.host.organizationId, s.host.issueId, s.host.statusId = context.OrganizationId, issueId, statusId
	return nil
}

func (s issuesV2Issues) GetIssueStatuses(context RequestContext) ([]v2.IssueStatus, error) {
	s.host.organizationId = context.OrganizationId
	return []v2.IssueStatus{{StatusId: 6, Name: "Review", Category: v2.IssueStatusCategoryWaiting, Active: true}}, nil
}

type issuesV2Extensions struct{ host *issuesV2Host }

func (s issuesV2Extensions) SetEntityExtensionData(context RequestContext, entityType string, entityId uint, data any) error {
	s.host.organizationId, s.host.entityType, s.host.entityId, s.host.data = context.OrganizationId, entityType, entityId, data
	return nil
}

func TestIssuesV2OverRPC(t *testing.T) {
	host := &issuesV2Host{}
	client := featureReadClient(t, host)
	// A plugin that is not organization-aware names the organization per call.
	context := RequestContext{OrganizationId: 7}

	if err := client.V2().Issues().SetIssueStatus(context, 6047, 6); err != nil || host.organizationId != 7 || host.issueId != 6047 || host.statusId != 6 {
		t.Fatalf("status: %v %+v", err, host)
	}
	if err := client.V2().Issues().SetIssueStatus(context, 6047, 4); err == nil || err.Error() != "a plugin cannot close an issue" {
		t.Fatalf("error lost: %v", err)
	}
	statuses, err := client.V2().Issues().GetIssueStatuses(context)
	if err != nil || len(statuses) != 1 || statuses[0].Category != v2.IssueStatusCategoryWaiting {
		t.Fatalf("statuses: %v %v", statuses, err)
	}
}

func TestExtensionsV2OverRPC(t *testing.T) {
	host := &issuesV2Host{}
	client := featureReadClient(t, host)
	context := RequestContext{OrganizationId: 7}

	if err := client.V2().Extensions().SetEntityExtensionData(context, v2.EntityTypeIssue, 6046, map[string]int{"open": 1}); err != nil {
		t.Fatal(err)
	}
	raw, ok := host.data.(json.RawMessage)
	if !ok || string(raw) != `{"open":1}` || host.organizationId != 7 || host.entityType != "issue" || host.entityId != 6046 {
		t.Fatalf("stored %+v", host)
	}
	if err := client.V2().Extensions().SetEntityExtensionData(context, v2.EntityTypeIssue, 6046, nil); err != nil || host.data != nil {
		t.Fatalf("nil must remove: %v %v", err, host.data)
	}
}

// v1Host has no V2(): an older host.
type v1Host struct{ API }

func TestIssuesV2OnHostWithoutV2(t *testing.T) {
	client := featureReadClient(t, &v1Host{})
	context := RequestContext{OrganizationId: 7}
	if err := client.V2().Issues().SetIssueStatus(context, 1, 2); err == nil || err.Error() != v2.ErrUnsupported.Error() {
		t.Fatalf("status: %v", err)
	}
	if _, err := client.V2().Issues().GetIssueStatuses(context); err == nil {
		t.Fatal("statuses must fail")
	}
	if err := client.V2().Extensions().SetEntityExtensionData(context, v2.EntityTypeIssue, 1, nil); err == nil {
		t.Fatal("extension data must fail")
	}
}
