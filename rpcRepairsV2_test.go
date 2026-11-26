package shared

import (
	"errors"
	"testing"
	"time"

	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

// repairsV2Host is a host exposing v2 Repairs that records what it received.
type repairsV2Host struct {
	API
	organizationId uint
	created        v2.RepairCreateRequest
	patched        v2.RepairPatchRequest
	statusId       uint
	commentsAfter  uint
	commentText    string
	images         []v2.RepairImage
}

func (h *repairsV2Host) V2() v2.API { return repairsV2{host: h} }

type repairsV2 struct {
	v2.API
	host *repairsV2Host
}

func (s repairsV2) Repairs() v2.RepairsAPI { return repairsV2Repairs{host: s.host} }

type repairsV2Repairs struct{ host *repairsV2Host }

func (s repairsV2Repairs) GetRepairStatuses(context RequestContext) ([]v2.RepairStatus, error) {
	s.host.organizationId = context.OrganizationId
	return []v2.RepairStatus{{StatusId: 3, Name: "At technician"}}, nil
}

func (s repairsV2Repairs) GetRepair(context RequestContext, repairId uint) (v2.Repair, error) {
	if repairId == 404 {
		return v2.Repair{}, errors.New("repairNotFound")
	}
	return v2.Repair{RepairId: repairId, RepairStatusId: 3}, nil
}

func (s repairsV2Repairs) CreateRepair(context RequestContext, request v2.RepairCreateRequest) (v2.RepairMutateResponse, error) {
	s.host.organizationId, s.host.created = context.OrganizationId, request
	return v2.RepairMutateResponse{RepairId: 77}, nil
}

func (s repairsV2Repairs) PatchRepair(context RequestContext, repairId uint, request v2.RepairPatchRequest) (v2.RepairMutateResponse, error) {
	s.host.patched = request
	return v2.RepairMutateResponse{RepairId: repairId}, nil
}

func (s repairsV2Repairs) SetRepairStatus(context RequestContext, repairId uint, repairStatusId uint) error {
	if repairStatusId == 99 {
		return errors.New("repairStatusNotFound")
	}
	s.host.statusId = repairStatusId
	return nil
}

func (s repairsV2Repairs) GetRepairComments(context RequestContext, repairId uint, request v2.RepairCommentsRequest) ([]v2.RepairComment, error) {
	s.host.commentsAfter = request.AfterCommentId
	return []v2.RepairComment{{CommentId: 12, Text: "Need a fan", CreatedUserId: 5}}, nil
}

func (s repairsV2Repairs) CreateRepairComment(context RequestContext, repairId uint, request v2.RepairCommentCreateRequest) (v2.RepairCommentMutateResponse, error) {
	s.host.commentText = request.Text
	return v2.RepairCommentMutateResponse{CommentId: 13}, nil
}

func (s repairsV2Repairs) AddRepairImages(context RequestContext, repairId uint, images []v2.RepairImage) error {
	s.host.images = images
	return nil
}

func TestRepairsV2OverRPC(t *testing.T) {
	host := &repairsV2Host{}
	repairs := featureReadClient(t, host).V2().Repairs()
	context := RequestContext{OrganizationId: 7}

	statuses, err := repairs.GetRepairStatuses(context)
	if err != nil || len(statuses) != 1 || statuses[0].StatusId != 3 || host.organizationId != 7 {
		t.Fatalf("statuses: %v %v", statuses, err)
	}
	if _, err := repairs.GetRepair(context, 404); err == nil || err.Error() != "repairNotFound" {
		t.Fatalf("error lost: %v", err)
	}

	received := time.Date(2026, 11, 13, 10, 12, 44, 0, time.UTC)
	created, err := repairs.CreateRepair(context, v2.RepairCreateRequest{CompanyId: 4, RepairStatusId: 3, DeviceName: "Laptop", DevicePassword: "1998", RegistrationDatetime: received})
	if err != nil || created.RepairId != 77 || !host.created.RegistrationDatetime.Equal(received) || host.created.DevicePassword != "1998" || host.created.RepairStatusId != 3 {
		t.Fatalf("create: %+v %v %+v", created, err, host.created)
	}

	serial := "SN-1"
	patched, err := repairs.PatchRepair(context, 77, v2.RepairPatchRequest{SerialNumber: &serial})
	if err != nil || patched.RepairId != 77 || host.patched.SerialNumber == nil || *host.patched.SerialNumber != "SN-1" || host.patched.DeviceName != nil {
		t.Fatalf("patch: %v %+v", err, host.patched)
	}

	if err := repairs.SetRepairStatus(context, 77, 4); err != nil || host.statusId != 4 {
		t.Fatalf("status: %v", err)
	}
	if err := repairs.SetRepairStatus(context, 77, 99); err == nil || err.Error() != "repairStatusNotFound" {
		t.Fatalf("status error lost: %v", err)
	}

	comments, err := repairs.GetRepairComments(context, 77, v2.RepairCommentsRequest{AfterCommentId: 11})
	if err != nil || len(comments) != 1 || comments[0].CreatedUserId != 5 || host.commentsAfter != 11 {
		t.Fatalf("comments: %v %v", comments, err)
	}
	comment, err := repairs.CreateRepairComment(context, 77, v2.RepairCommentCreateRequest{Text: "Thanks"})
	if err != nil || comment.CommentId != 13 || host.commentText != "Thanks" {
		t.Fatalf("comment: %+v %v", comment, err)
	}

	if err := repairs.AddRepairImages(context, 77, []v2.RepairImage{{FileName: "1.jpg", FileBlob: []byte{0xff, 0xd8}}}); err != nil || len(host.images) != 1 || len(host.images[0].FileBlob) != 2 {
		t.Fatalf("images: %v %+v", err, host.images)
	}
}

// A host without v2 answers ErrUnsupported instead of failing the call.
func TestRepairsV2Unsupported(t *testing.T) {
	repairs := featureReadClient(t, &unsupportedFeatureProbe{}).V2().Repairs()
	if _, err := repairs.GetRepairStatuses(RequestContext{}); err == nil || err.Error() != v2.ErrUnsupported.Error() {
		t.Fatalf("want unsupported, got %v", err)
	}
}
