package shared

import (
	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

// Repairs of plugin API v2. Served only by hosts that expose v2
// (v2.VersionedAPI); others answer v2.ErrUnsupported.

type GetRepairStatusesV2Request struct {
	Context RequestContext
}
type GetRepairStatusesV2Response struct {
	Statuses []v2.RepairStatus
	Error    error
}

type GetRepairV2Request struct {
	Context  RequestContext
	RepairId uint
}
type RepairV2Response struct {
	Repair v2.Repair
	Error  error
}

type CreateRepairV2Request struct {
	Context RequestContext
	Request v2.RepairCreateRequest
}

type RepairMutateV2Response struct {
	Response v2.RepairMutateResponse
	Error    error
}

type PatchRepairV2Request struct {
	Context  RequestContext
	RepairId uint
	Request  v2.RepairPatchRequest
}

type SetRepairStatusV2Request struct {
	Context        RequestContext
	RepairId       uint
	RepairStatusId uint
}

type RepairV2ErrorResponse struct {
	Error error
}

type GetRepairCommentsV2Request struct {
	Context  RequestContext
	RepairId uint
	Request  v2.RepairCommentsRequest
}
type GetRepairCommentsV2Response struct {
	Comments []v2.RepairComment
	Error    error
}

type CreateRepairCommentV2Request struct {
	Context  RequestContext
	RepairId uint
	Request  v2.RepairCommentCreateRequest
}
type CreateRepairCommentV2Response struct {
	Response v2.RepairCommentMutateResponse
	Error    error
}

type AddRepairImagesV2Request struct {
	Context  RequestContext
	RepairId uint
	Images   []v2.RepairImage
}

func (m *apiRPCServer) repairsV2() (v2.RepairsAPI, bool) {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		return nil, false
	}
	return impl.V2().Repairs(), true
}

func (m *apiRPCServer) GetRepairStatusesV2(req GetRepairStatusesV2Request, resp *GetRepairStatusesV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	statuses, err := repairs.GetRepairStatuses(req.Context)
	resp.Statuses = statuses
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) GetRepairV2(req GetRepairV2Request, resp *RepairV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	repair, err := repairs.GetRepair(req.Context, req.RepairId)
	resp.Repair = repair
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) CreateRepairV2(req CreateRepairV2Request, resp *RepairMutateV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	response, err := repairs.CreateRepair(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) PatchRepairV2(req PatchRepairV2Request, resp *RepairMutateV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	response, err := repairs.PatchRepair(req.Context, req.RepairId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) SetRepairStatusV2(req SetRepairStatusV2Request, resp *RepairV2ErrorResponse) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	resp.Error = encodableError(repairs.SetRepairStatus(req.Context, req.RepairId, req.RepairStatusId))
	return nil
}

func (m *apiRPCServer) GetRepairCommentsV2(req GetRepairCommentsV2Request, resp *GetRepairCommentsV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	comments, err := repairs.GetRepairComments(req.Context, req.RepairId, req.Request)
	resp.Comments = comments
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) CreateRepairCommentV2(req CreateRepairCommentV2Request, resp *CreateRepairCommentV2Response) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	response, err := repairs.CreateRepairComment(req.Context, req.RepairId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCServer) AddRepairImagesV2(req AddRepairImagesV2Request, resp *RepairV2ErrorResponse) error {
	repairs, ok := m.repairsV2()
	if !ok {
		resp.Error = encodableError(v2.ErrUnsupported)
		return nil
	}
	resp.Error = encodableError(repairs.AddRepairImages(req.Context, req.RepairId, req.Images))
	return nil
}

type rpcRepairsV2API struct {
	api *apiRPCClient
}

func (a rpcRepairsV2API) GetRepairStatuses(context RequestContext) ([]v2.RepairStatus, error) {
	var reply GetRepairStatusesV2Response
	if err := a.api.client.Call("Plugin.GetRepairStatusesV2", GetRepairStatusesV2Request{Context: context}, &reply); err != nil {
		return nil, err
	}
	return reply.Statuses, reply.Error
}

func (a rpcRepairsV2API) GetRepair(context RequestContext, repairId uint) (v2.Repair, error) {
	var reply RepairV2Response
	if err := a.api.client.Call("Plugin.GetRepairV2", GetRepairV2Request{Context: context, RepairId: repairId}, &reply); err != nil {
		return v2.Repair{}, err
	}
	return reply.Repair, reply.Error
}

func (a rpcRepairsV2API) CreateRepair(context RequestContext, request v2.RepairCreateRequest) (v2.RepairMutateResponse, error) {
	var reply RepairMutateV2Response
	if err := a.api.client.Call("Plugin.CreateRepairV2", CreateRepairV2Request{Context: context, Request: request}, &reply); err != nil {
		return v2.RepairMutateResponse{}, err
	}
	return reply.Response, reply.Error
}

func (a rpcRepairsV2API) PatchRepair(context RequestContext, repairId uint, request v2.RepairPatchRequest) (v2.RepairMutateResponse, error) {
	var reply RepairMutateV2Response
	if err := a.api.client.Call("Plugin.PatchRepairV2", PatchRepairV2Request{Context: context, RepairId: repairId, Request: request}, &reply); err != nil {
		return v2.RepairMutateResponse{}, err
	}
	return reply.Response, reply.Error
}

func (a rpcRepairsV2API) SetRepairStatus(context RequestContext, repairId uint, repairStatusId uint) error {
	var reply RepairV2ErrorResponse
	if err := a.api.client.Call("Plugin.SetRepairStatusV2", SetRepairStatusV2Request{Context: context, RepairId: repairId, RepairStatusId: repairStatusId}, &reply); err != nil {
		return err
	}
	return reply.Error
}

func (a rpcRepairsV2API) GetRepairComments(context RequestContext, repairId uint, request v2.RepairCommentsRequest) ([]v2.RepairComment, error) {
	var reply GetRepairCommentsV2Response
	if err := a.api.client.Call("Plugin.GetRepairCommentsV2", GetRepairCommentsV2Request{Context: context, RepairId: repairId, Request: request}, &reply); err != nil {
		return nil, err
	}
	return reply.Comments, reply.Error
}

func (a rpcRepairsV2API) CreateRepairComment(context RequestContext, repairId uint, request v2.RepairCommentCreateRequest) (v2.RepairCommentMutateResponse, error) {
	var reply CreateRepairCommentV2Response
	if err := a.api.client.Call("Plugin.CreateRepairCommentV2", CreateRepairCommentV2Request{Context: context, RepairId: repairId, Request: request}, &reply); err != nil {
		return v2.RepairCommentMutateResponse{}, err
	}
	return reply.Response, reply.Error
}

func (a rpcRepairsV2API) AddRepairImages(context RequestContext, repairId uint, images []v2.RepairImage) error {
	var reply RepairV2ErrorResponse
	if err := a.api.client.Call("Plugin.AddRepairImagesV2", AddRepairImagesV2Request{Context: context, RepairId: repairId, Images: images}, &reply); err != nil {
		return err
	}
	return reply.Error
}
