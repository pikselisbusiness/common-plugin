package shared

import (
	"net/rpc"

	"github.com/pikselisbusiness/go-plugin"
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

// GetCompanyByCode
type GetCompanyByCodeRequest struct {
	Context     RequestContext
	CompanyCode string
}
type GetCompanyByCodeResponse struct {
	Company Company
	Error   error
}

// GetCompanyByVatCode
type GetCompanyByVatCodeRequest struct {
	Context        RequestContext
	CompanyVatCode string
}
type GetCompanyByVatCodeResponse struct {
	Company Company
	Error   error
}

// CreateCompany
type CreateCompanyRequest struct {
	Context RequestContext
	Company Company
}
type CreateCompanyResponse struct {
	CompanyId uint
	Error     error
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

// GetProducts
type GetProductsRequest struct {
	Context RequestContext
	Request ProductsRequest
}
type GetProductsResponse struct {
	Response ProductsResponse
	Error    error
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

// GetProducts
type GetProductCategoriesRequest struct {
	Context RequestContext
	Request ProductCategoriesRequest
}
type GetProductCategoriesResponse struct {
	Response ProductCategoriesResponse
	Error    error
}

// GetProductById
type GetProductByIdRequest struct {
	RequestContext RequestContext
	ProductId      uint
}
type GetProductByIdResponse struct {
	Product Product
	Error   error
}

// GetProductByAnyField
type GetProductByAnyFieldRequest struct {
	RequestContext RequestContext
	FieldName      string
	FieldValue     any
}
type GetProductByAnyFieldResponse struct {
	Product Product
	Error   error
}

// CreateProduct
type CreateProductRequest struct {
	RequestContext RequestContext
	Request        ProductCreateEditRequest
}
type CreateProductResponse struct {
	ProductId uint
	Error     error
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

// GetOrders
type GetOrdersRequest struct {
	Context RequestContext
	Request OrdersRequest
}
type GetOrdersResponse struct {
	OrdersResponse OrdersResponse
	Error          error
}

// GetCountryByName
type GetCountryByNameRequest struct {
	Context RequestContext
	Name    string
}
type GetCountryByNameResponse struct {
	Country Country
	Error   error
}

// GetAllCountries
type GetAllCountriesRequest struct {
	Context RequestContext
}
type GetAllCountriesResponse struct {
	Countries []Country
	Error     error
}

// CreateOrder
type CreateOrderRequest struct {
	Context RequestContext
	Request OrderCreateRequest
}
type CreateOrderResponse struct {
	OrderId       uint
	Error         error
	ErrorResponse OrderErrorResponse
}

// CreateInvoice
type CreateInvoiceRequest struct {
	Context RequestContext
	Request InvoiceCreateUpdateRequest
}
type CreateInvoiceResponse struct {
	InvoiceId     uint
	Error         error
	ErrorResponse InvoiceErrorResponse
}

// GetInvoiceExistsByDocument
type GetInvoiceExistsByDocumentRequest struct {
	Context  RequestContext
	Document string
}
type GetInvoiceExistsByDocumentResponse struct {
	Exists bool
	Error  error
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

func (m *apiRPCServer) GetCompanyByCode(req GetCompanyByCodeRequest, resp *GetCompanyByCodeResponse) error {
	company, err := m.impl.GetCompanyByCode(req.Context, req.CompanyCode)

	resp.Company = company
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetCompanyByCode(context RequestContext, companyCode string) (Company, error) {

	var reply GetCompanyByCodeResponse
	err := m.client.Call("Plugin.GetCompanyByCode", GetCompanyByCodeRequest{
		Context:     context,
		CompanyCode: companyCode,
	}, &reply)
	if err != nil {
		return Company{}, err
	}

	return reply.Company, reply.Error
}

func (m *apiRPCServer) GetCompanyByVatCode(req GetCompanyByVatCodeRequest, resp *GetCompanyByVatCodeResponse) error {
	company, err := m.impl.GetCompanyByVatCode(req.Context, req.CompanyVatCode)

	resp.Company = company
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetCompanyByVatCode(context RequestContext, companyVatCode string) (Company, error) {

	var reply GetCompanyByVatCodeResponse
	err := m.client.Call("Plugin.GetCompanyByVatCode", GetCompanyByVatCodeRequest{
		Context:        context,
		CompanyVatCode: companyVatCode,
	}, &reply)
	if err != nil {
		return Company{}, err
	}

	return reply.Company, reply.Error
}

func (m *apiRPCServer) CreateCompany(req CreateCompanyRequest, resp *CreateCompanyResponse) error {
	companyId, err := m.impl.CreateCompany(req.Context, req.Company)

	resp.CompanyId = companyId
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) CreateCompany(context RequestContext, company Company) (uint, error) {

	var reply CreateCompanyResponse
	err := m.client.Call("Plugin.CreateCompany", CreateCompanyRequest{
		Context: context,
		Company: company,
	}, &reply)
	if err != nil {
		return 0, err
	}

	return reply.CompanyId, reply.Error
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
func (m *apiRPCServer) GetProducts(req GetProductsRequest, resp *GetProductsResponse) error {
	response, err := m.impl.GetProducts(req.Context, req.Request)

	resp.Response = response
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error) {

	var reply GetProductsResponse
	err := m.client.Call("Plugin.GetProducts", GetProductsRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return ProductsResponse{}, err
	}

	return reply.Response, reply.Error
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

func (m *apiRPCServer) GetProductByAnyField(req GetProductByAnyFieldRequest, resp *GetProductByAnyFieldResponse) error {
	product, err := m.impl.GetProductByAnyField(req.RequestContext, req.FieldName, req.FieldValue)
	resp.Product = product
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetProductByAnyField(rc RequestContext, fieldName string, fieldValue any) (Product, error) {

	var reply GetProductByAnyFieldResponse
	err := m.client.Call("Plugin.GetProductByAnyField", GetProductByAnyFieldRequest{
		RequestContext: rc,
		FieldName:      fieldName,
		FieldValue:     fieldValue,
	}, &reply)
	if err != nil {
		return Product{}, err
	}

	return reply.Product, reply.Error
}

func (m *apiRPCServer) CreateProduct(req CreateProductRequest, resp *CreateProductResponse) error {
	productId, err := m.impl.CreateProduct(req.RequestContext, req.Request)
	resp.ProductId = productId
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) CreateProduct(rc RequestContext, request ProductCreateEditRequest) (uint, error) {

	var reply CreateProductResponse
	err := m.client.Call("Plugin.CreateProduct", CreateProductRequest{
		RequestContext: rc,
		Request:        request,
	}, &reply)
	if err != nil {
		return 0, err
	}

	return reply.ProductId, reply.Error
}

func (m *apiRPCServer) GetProductCategories(req GetProductCategoriesRequest, resp *GetProductCategoriesResponse) error {
	response, err := m.impl.GetProductCategories(req.Context, req.Request)

	resp.Response = response
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetProductCategories(context RequestContext, request ProductCategoriesRequest) (ProductCategoriesResponse, error) {

	var reply GetProductCategoriesResponse
	err := m.client.Call("Plugin.GetProductCategories", GetProductCategoriesRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return ProductCategoriesResponse{}, err
	}

	return reply.Response, reply.Error
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

func (m *apiRPCServer) CreateOrder(req CreateOrderRequest, resp *CreateOrderResponse) error {
	orderId, err, errResponse := m.impl.CreateOrder(req.Context, req.Request)

	resp.OrderId = orderId
	resp.Error = encodableError(err)
	resp.ErrorResponse = errResponse

	return nil
}
func (m *apiRPCClient) CreateOrder(context RequestContext, request OrderCreateRequest) (uint, error, OrderErrorResponse) {

	var reply CreateOrderResponse
	err := m.client.Call("Plugin.CreateOrder", CreateOrderRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return 0, err, OrderErrorResponse{}
	}

	return reply.OrderId, reply.Error, reply.ErrorResponse
}

func (m *apiRPCServer) GetCountryByName(req GetCountryByNameRequest, resp *GetCountryByNameResponse) error {
	data, err := m.impl.GetCountryByName(req.Context, req.Name)

	resp.Country = data
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetCountryByName(context RequestContext, name string) (Country, error) {

	var reply GetCountryByNameResponse
	err := m.client.Call("Plugin.GetCountryByName", GetCountryByNameRequest{
		Context: context,
		Name:    name,
	}, &reply)
	if err != nil {
		return Country{}, err
	}

	return reply.Country, reply.Error
}

func (m *apiRPCServer) GetAllCountries(req GetAllCountriesRequest, resp *GetAllCountriesResponse) error {
	data, err := m.impl.GetAllCountries(req.Context)

	resp.Countries = data
	resp.Error = encodableError(err)

	return nil
}
func (m *apiRPCClient) GetAllCountries(context RequestContext) ([]Country, error) {

	var reply GetAllCountriesResponse
	err := m.client.Call("Plugin.GetAllCountries", GetAllCountriesRequest{
		Context: context,
	}, &reply)
	if err != nil {
		return []Country{}, err
	}

	return reply.Countries, reply.Error
}

func (m *apiRPCServer) CreateInvoice(req CreateInvoiceRequest, resp *CreateInvoiceResponse) error {
	invoiceId, error, errorResponse := m.impl.CreateInvoice(req.Context, req.Request)

	resp.InvoiceId = invoiceId
	resp.Error = encodableError(error)
	resp.ErrorResponse = errorResponse
	return nil
}
func (m *apiRPCClient) CreateInvoice(context RequestContext, request InvoiceCreateUpdateRequest) (uint, error, InvoiceErrorResponse) {

	var reply CreateInvoiceResponse
	err := m.client.Call("Plugin.CreateInvoice", CreateInvoiceRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return 0, err, InvoiceErrorResponse{}
	}

	return reply.InvoiceId, reply.Error, reply.ErrorResponse
}
func (m *apiRPCServer) GetInvoiceExistsByDocument(req GetInvoiceExistsByDocumentRequest, resp *GetInvoiceExistsByDocumentResponse) error {
	exists, error := m.impl.GetInvoiceExistsByDocument(req.Context, req.Document)

	resp.Exists = exists
	resp.Error = encodableError(error)
	return nil
}
func (m *apiRPCClient) GetInvoiceExistsByDocument(context RequestContext, document string) (bool, error) {

	var reply GetInvoiceExistsByDocumentResponse
	err := m.client.Call("Plugin.GetInvoiceExistsByDocument", GetInvoiceExistsByDocumentRequest{
		Context:  context,
		Document: document,
	}, &reply)
	if err != nil {
		return false, err
	}

	return reply.Exists, reply.Error
}
