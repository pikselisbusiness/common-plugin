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

// GetInvoiceProducts
type GetInvoiceProductsRequest struct {
	Context   RequestContext
	InvoiceId uint
}
type GetInvoiceProductsResponse struct {
	Products []InvoiceProduct
	Error    error
}

// GetInvoiceReferences
type GetInvoiceReferencesRequest struct {
	Context   RequestContext
	InvoiceId uint
}
type GetInvoiceReferencesResponse struct {
	References []InvoiceReference
	Error      error
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

// GetProductsMap
type GetProductsMapRequest struct {
	Context    RequestContext
	ProductIds []uint
}
type GetProductsMapResponse struct {
	Products map[uint]Product
	Error    error
}

// GetProductsMap
type GetProductByIdRequest struct {
	RequestContext RequestContext
	ProductId      uint
}
type GetProductByIdResponse struct {
	Product Product
	Error   error
}

// GetOrderById
type GetOrderByIdRequest struct {
	RequestContext RequestContext
	OrderId        uint
}
type GetOrderByIdResponse struct {
	Order Order
	Error error
}

// GetDivisions
type GetOrdersRequest struct {
	Context RequestContext
	Request OrdersRequest
}
type GetOrdersResponse struct {
	OrdersResponse OrdersResponse
	Error          error
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
	resp.Error = encodableError(err)

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
	resp.Error = encodableError(err)

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

func (m *apiRPCServer) GetInvoiceProducts(req GetInvoiceProductsRequest, resp *GetInvoiceProductsResponse) error {
	data, err := m.impl.GetInvoiceProducts(req.Context, req.InvoiceId)

	resp.Products = data
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetInvoiceProducts(context RequestContext, invoiceId uint) ([]InvoiceProduct, error) {

	var reply GetInvoiceProductsResponse
	err := m.client.Call("Plugin.GetInvoiceProducts", GetInvoiceProductsRequest{
		Context:   context,
		InvoiceId: invoiceId,
	}, &reply)
	if err != nil {
		return []InvoiceProduct{}, err
	}

	return reply.Products, reply.Error
}

func (m *apiRPCServer) GetInvoiceReferences(req GetInvoiceReferencesRequest, resp *GetInvoiceReferencesResponse) error {
	data, err := m.impl.GetInvoiceReferences(req.Context, req.InvoiceId)

	resp.References = data
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetInvoiceReferences(context RequestContext, invoiceId uint) ([]InvoiceReference, error) {

	var reply GetInvoiceReferencesResponse
	err := m.client.Call("Plugin.GetInvoiceReferences", GetInvoiceReferencesRequest{
		Context:   context,
		InvoiceId: invoiceId,
	}, &reply)
	if err != nil {
		return []InvoiceReference{}, err
	}

	return reply.References, reply.Error
}

func (m *apiRPCServer) GetDivisions(req GetDivisionsRequest, resp *GetDivisionsResponse) error {
	data, err := m.impl.GetDivisions(req.Context, req.Request)

	resp.Divisions = data
	resp.Error = encodableError(err)

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
	resp.Error = encodableError(err)

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

func (m *apiRPCServer) GetProductsMap(req GetProductsMapRequest, resp *GetProductsMapResponse) error {
	products, err := m.impl.GetProductsMap(req.Context, req.ProductIds)

	resp.Products = products
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetProductsMap(context RequestContext, productIds []uint) (map[uint]Product, error) {

	var reply GetProductsMapResponse
	err := m.client.Call("Plugin.GetProductsMap", GetProductsMapRequest{
		Context:    context,
		ProductIds: productIds,
	}, &reply)
	if err != nil {
		return map[uint]Product{}, err
	}

	return reply.Products, reply.Error
}

func (m *apiRPCServer) GetProductById(req GetProductByIdRequest, resp *GetProductByIdResponse) error {
	product, err := m.impl.GetProductById(req.RequestContext, req.ProductId)
	resp.Product = product
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetProductById(rc RequestContext, productId uint) (Product, error) {

	var reply GetProductByIdResponse
	err := m.client.Call("Plugin.GetProductById", GetProductByIdRequest{
		RequestContext: rc,
		ProductId:      productId,
	}, &reply)
	if err != nil {
		return Product{}, err
	}

	return reply.Product, reply.Error
}

func (m *apiRPCServer) GetOrderById(req GetOrderByIdRequest, resp *GetOrderByIdResponse) error {
	order, err := m.impl.GetOrderById(req.RequestContext, req.OrderId)
	resp.Order = order
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetOrderById(rc RequestContext, orderId uint) (Order, error) {

	var reply GetOrderByIdResponse
	err := m.client.Call("Plugin.GetOrderById", GetOrderByIdRequest{
		RequestContext: rc,
		OrderId:        orderId,
	}, &reply)
	if err != nil {
		return Order{}, err
	}

	return reply.Order, reply.Error
}

func (m *apiRPCServer) GetOrders(req GetOrdersRequest, resp *GetOrdersResponse) error {
	data, err := m.impl.GetOrders(req.Context, req.Request)

	resp.OrdersResponse = data
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error) {

	var reply GetOrdersResponse
	err := m.client.Call("Plugin.GetOrders", GetOrdersRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return OrdersResponse{}, err
	}

	return reply.OrdersResponse, reply.Error
}
