package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type apiRPCClient struct {
	client *rpc.Client
	broker *plugin.MuxBroker
}

type apiRPCServer struct {
	impl   API
	broker *plugin.MuxBroker
}

//RegisterCronJob

type RegisterCronJobRequest struct {
	Schedule string
}
type RegisterCronJobResponse struct {
}

//GetConfigVariable

type GetConfigVariableRequest struct {
	VariableName string
}
type GetConfigVariableResponse struct {
	VariableData string
	Error        error
}

//GetUserInfoForUserId

type GetUserInfoForUserIdRequest struct {
	UserId uint
}
type GetUserInfoForUserIdResponse struct {
	UserInfo UserInfo
}

// GetInvoices
type GetInvoicesRequest struct {
	Request InvoicesRequest
}
type GetInvoicesResponse struct {
	Response InvoicesListResponse
	Error    error
}

// GetDivisions
type GetDivisionsRequest struct {
	Context RequestContext
	Request DivisionsRequest
}
type GetDivisionsResponse struct {
	Divisions []Division
	Error     error
}

// GetCompanyById
type GetCompanyByIdRequest struct {
	Context   RequestContext
	CompanyId uint
}
type GetCompanyByIdResponse struct {
	Company Company
	Error   error
}

// GetCompaniesMap
type GetCompaniesMapRequest struct {
	Context    RequestContext
	CompanyIds []uint
}
type GetCompaniesMapResponse struct {
	Companies map[uint]Company
	Error     error
}

func (m *apiRPCServer) RegisterCronJob(req RegisterCronJobRequest, resp *RegisterCronJobResponse) error {
	m.impl.RegisterCronJob(req.Schedule)

	return nil
}
func (m *apiRPCClient) RegisterCronJob(schedule string) {

	var reply RegisterCronJobResponse
	err := m.client.Call("Plugin.RegisterCronJob", RegisterCronJobRequest{
		Schedule: schedule,
	}, &reply)
	if err != nil {
		// return err
	}

	// return nil
}

func (m *apiRPCServer) GetConfigVariable(req GetConfigVariableRequest, resp *GetConfigVariableResponse) error {
	data, err := m.impl.GetConfigVariable(req.VariableName)

	resp.VariableData = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetConfigVariable(variableName string) (string, error) {

	var reply GetConfigVariableResponse
	err := m.client.Call("Plugin.GetConfigVariable", GetConfigVariableRequest{
		VariableName: variableName,
	}, &reply)
	if err != nil {
		return "", err
	}

	return reply.VariableData, reply.Error
}

func (m *apiRPCServer) GetUserInfoForUserId(req GetUserInfoForUserIdRequest, resp *GetUserInfoForUserIdResponse) error {
	data := m.impl.GetUserInfoForUserId(req.UserId)

	resp.UserInfo = data

	return nil
}
func (m *apiRPCClient) GetUserInfoForUserId(userId uint) UserInfo {

	var reply GetUserInfoForUserIdResponse
	err := m.client.Call("Plugin.GetUserInfoForUserId", GetUserInfoForUserIdRequest{
		UserId: userId,
	}, &reply)
	if err != nil {
		return UserInfo{}
	}

	return reply.UserInfo
}

func (m *apiRPCServer) GetInvoices(req GetInvoicesRequest, resp *GetInvoicesResponse) error {
	data, err := m.impl.GetInvoices(req.Request)

	resp.Response = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetInvoices(request InvoicesRequest) (InvoicesListResponse, error) {

	var reply GetInvoicesResponse
	err := m.client.Call("Plugin.GetInvoices", GetInvoicesRequest{
		Request: request,
	}, &reply)
	if err != nil {
		return InvoicesListResponse{}, err
	}

	return reply.Response, reply.Error
}

func (m *apiRPCServer) GetDivisions(req GetDivisionsRequest, resp *GetDivisionsResponse) error {
	data, err := m.impl.GetDivisions(req.Context, req.Request)

	resp.Divisions = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error) {

	var reply GetDivisionsResponse
	err := m.client.Call("Plugin.GetDivisions", GetDivisionsRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return []Division{}, err
	}

	return reply.Divisions, reply.Error
}

func (m *apiRPCServer) GetCompanyById(req GetCompanyByIdRequest, resp *GetCompanyByIdResponse) error {
	company, err := m.impl.GetCompanyById(req.Context, req.CompanyId)

	resp.Company = company
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetCompanyById(context RequestContext, companyId uint) (Company, error) {

	var reply GetCompanyByIdResponse
	err := m.client.Call("Plugin.GetCompanyById", GetCompanyByIdRequest{
		Context:   context,
		CompanyId: companyId,
	}, &reply)
	if err != nil {
		return Company{}, err
	}

	return reply.Company, reply.Error
}

func (m *apiRPCServer) GetCompaniesMap(req GetCompaniesMapRequest, resp *GetCompaniesMapResponse) error {
	companies, err := m.impl.GetCompaniesMap(req.Context, req.CompanyIds)

	resp.Companies = companies
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetCompaniesMap(context RequestContext, companyIds []uint) (map[uint]Company, error) {

	var reply GetCompaniesMapResponse
	err := m.client.Call("Plugin.GetCompaniesMap", GetCompaniesMapRequest{
		Context:    context,
		CompanyIds: companyIds,
	}, &reply)
	if err != nil {
		return map[uint]Company{}, err
	}

	return reply.Companies, reply.Error
}
