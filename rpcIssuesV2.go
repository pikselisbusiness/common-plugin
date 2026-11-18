package shared

import (
	"encoding/json"

	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

// Issues and Extensions of plugin API v2 (ERP-6046, ERP-6047). Served only by
// hosts that expose v2 (v2.VersionedAPI); others answer v2.ErrUnsupported.

type SetIssueStatusV2Request struct {
	Context  RequestContext
	IssueId  uint
	StatusId uint
}
type SetIssueStatusV2Response struct {
	Error error
}

type GetIssueStatusesV2Request struct {
	Context RequestContext
}
type GetIssueStatusesV2Response struct {
	Statuses []v2.IssueStatus
	Error    error
}

type SetEntityExtensionDataV2Request struct {
	Context    RequestContext
	EntityType string
	EntityId   uint
	Data       []byte // JSON; empty removes
}
type SetEntityExtensionDataV2Response struct {
	Error error
}

func (m *apiRPCServer) SetIssueStatusV2(req SetIssueStatusV2Request, resp *SetIssueStatusV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	resp.Error = encodableError(impl.V2().Issues().SetIssueStatus(req.Context, req.IssueId, req.StatusId))
	return nil
}

func (m *apiRPCServer) GetIssueStatusesV2(req GetIssueStatusesV2Request, resp *GetIssueStatusesV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	statuses, err := impl.V2().Issues().GetIssueStatuses(req.Context)
	resp.Statuses = statuses
	resp.Error = encodableError(err)
	return nil
}

// SetEntityExtensionDataV2 hands the host the JSON the plugin marshalled; a
// host's ExtensionsAPI receives it as json.RawMessage (nil when removed).
func (m *apiRPCServer) SetEntityExtensionDataV2(req SetEntityExtensionDataV2Request, resp *SetEntityExtensionDataV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	var data any
	if len(req.Data) > 0 {
		data = json.RawMessage(req.Data)
	}
	resp.Error = encodableError(impl.V2().Extensions().SetEntityExtensionData(req.Context, req.EntityType, req.EntityId, data))
	return nil
}

type rpcIssuesV2API struct {
	api *apiRPCClient
}

func (a rpcIssuesV2API) SetIssueStatus(context RequestContext, issueId uint, statusId uint) error {
	var reply SetIssueStatusV2Response
	request := SetIssueStatusV2Request{Context: context, IssueId: issueId, StatusId: statusId}
	if err := a.api.client.Call("Plugin.SetIssueStatusV2", request, &reply); err != nil {
		return err
	}
	return reply.Error
}

func (a rpcIssuesV2API) GetIssueStatuses(context RequestContext) ([]v2.IssueStatus, error) {
	var reply GetIssueStatusesV2Response
	if err := a.api.client.Call("Plugin.GetIssueStatusesV2", GetIssueStatusesV2Request{Context: context}, &reply); err != nil {
		return nil, err
	}
	return reply.Statuses, reply.Error
}

type rpcExtensionsV2API struct {
	api *apiRPCClient
}

func (a rpcExtensionsV2API) SetEntityExtensionData(context RequestContext, entityType string, entityId uint, data any) error {
	var raw []byte
	if data != nil {
		var err error
		if raw, err = json.Marshal(data); err != nil {
			return err
		}
	}
	var reply SetEntityExtensionDataV2Response
	request := SetEntityExtensionDataV2Request{Context: context, EntityType: entityType, EntityId: entityId, Data: raw}
	if err := a.api.client.Call("Plugin.SetEntityExtensionDataV2", request, &reply); err != nil {
		return err
	}
	return reply.Error
}
