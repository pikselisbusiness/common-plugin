package shared

import (
	"errors"

	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

// GetInvoicesV2
type GetInvoicesV2Request struct {
	Context RequestContext
	Request v2.InvoicesRequest
}

// GetInvoicesV2
type GetInvoicesV2Response struct {
	Response v2.InvoicesListResponse
	Error    error
}

type GetCompanyV2Request struct {
	Context   RequestContext
	CompanyId uint
}

type GetCompanyV2Response struct {
	Company v2.Company
	Error   error
}

type GetCompaniesV2Request struct {
	Context RequestContext
	Request v2.CompaniesRequest
}

type GetCompaniesV2Response struct {
	Response v2.CompaniesResponse
	Error    error
}

type CreateCompanyV2Request struct {
	Context RequestContext
	Request v2.CompanyCreateUpdateRequest
}

type CreateCompanyV2Response struct {
	Response v2.CompanyMutateResponse
	Error    error
}

type UpdateCompanyV2Request struct {
	Context   RequestContext
	CompanyId uint
	Request   v2.CompanyCreateUpdateRequest
}

type UpdateCompanyV2Response struct {
	Response v2.CompanyMutateResponse
	Error    error
}

type PatchCompanyV2Request struct {
	Context   RequestContext
	CompanyId uint
	Request   v2.CompanyPatchRequest
}

type PatchCompanyV2Response struct {
	Response v2.CompanyMutateResponse
	Error    error
}

// GetProductV2
type GetProductV2Request struct {
	Context   RequestContext
	ProductId uint
}

// GetProductV2
type GetProductV2Response struct {
	Product v2.Product
	Error   error
}

// GetProductsV2
type GetProductsV2Request struct {
	Context RequestContext
	Request v2.ProductsRequest
}

// GetProductsV2
type GetProductsV2Response struct {
	Response v2.ProductsResponse
	Error    error
}

// CreateProductV2
type CreateProductV2Request struct {
	Context RequestContext
	Request v2.ProductCreateEditRequest
}

// CreateProductV2
type CreateProductV2Response struct {
	Response v2.ProductMutateResponse
	Error    error
}

// UpdateProductV2
type UpdateProductV2Request struct {
	Context   RequestContext
	ProductId uint
	Request   v2.ProductCreateEditRequest
}

// UpdateProductV2
type UpdateProductV2Response struct {
	Response v2.ProductMutateResponse
	Error    error
}

// PatchProductV2
type PatchProductV2Request struct {
	Context   RequestContext
	ProductId uint
	Request   v2.ProductPatchRequest
}

// PatchProductV2
type PatchProductV2Response struct {
	Response v2.ProductMutateResponse
	Error    error
}

type GetProductStocksV2Request struct {
	Context RequestContext
	Request v2.ProductStocksRequest
}

type GetProductStocksV2Response struct {
	Response v2.ProductStocksResponse
	Error    error
}

type GetProductStocksWarehousesRequest struct {
	Context         RequestContext
	ProductId       uint
	FilterAvailable bool
}

type GetProductStocksWarehousesResponse struct {
	Response v2.ProductWarehousesQuantitiesResponse
	Error    error
}

type SetProductStocksRequest struct {
	Context RequestContext
	Request v2.ProductStockSetRequest
}

type SetProductStocksResponse struct {
	Response v2.ProductStockAdjustmentResponse
	Error    error
}

type AdjustProductStocksRequest struct {
	Context RequestContext
	Request v2.ProductStockAdjustRequest
}

type AdjustProductStocksResponse struct {
	Response v2.ProductStockAdjustmentResponse
	Error    error
}

// GetOrderV2
type GetOrderV2Request struct {
	Context RequestContext
	OrderId uint
}

// GetOrderV2
type GetOrderV2Response struct {
	Order v2.OrderDetail
	Error error
}

// GetOrdersV2
type GetOrdersV2Request struct {
	Context RequestContext
	Request v2.OrdersRequest
}

// GetOrdersV2
type GetOrdersV2Response struct {
	Response v2.OrdersResponse
	Error    error
}

// CreateOrderV2
type CreateOrderV2Request struct {
	Context RequestContext
	Request v2.OrderCreateRequest
}

// CreateOrderV2
type CreateOrderV2Response struct {
	Response      v2.OrderMutateResponse
	Error         error
	ErrorResponse OrderErrorResponse
}

// UpdateOrderV2
type UpdateOrderV2Request struct {
	Context RequestContext
	OrderId uint
	Request v2.OrderUpdateRequest
}

// UpdateOrderV2
type UpdateOrderV2Response struct {
	Response      v2.OrderMutateResponse
	Error         error
	ErrorResponse OrderErrorResponse
}

// PatchOrderV2
type PatchOrderV2Request struct {
	Context RequestContext
	OrderId uint
	Request v2.OrderPatchRequest
}

// PatchOrderV2
type PatchOrderV2Response struct {
	Response v2.OrderMutateResponse
	Error    error
}

// CreateInvoiceFromOrderV2
type CreateInvoiceFromOrderV2Request struct {
	Context RequestContext
	OrderId uint
	Request v2.OrderInvoiceCreateRequest
}

// CreateInvoiceFromOrderV2
type CreateInvoiceFromOrderV2Response struct {
	Response      v2.OrderInvoiceCreateResponse
	Error         error
	ErrorResponse InvoiceErrorResponse
}

// GetOrderPaymentV2
type GetOrderPaymentV2Request struct {
	Context   RequestContext
	PaymentId uint
}

// GetOrderPaymentV2
type GetOrderPaymentV2Response struct {
	Payment v2.OrderPayment
	Error   error
}

// GetOrderPaymentsV2
type GetOrderPaymentsV2Request struct {
	Context RequestContext
	Request v2.OrderPaymentsRequest
}

// GetOrderPaymentsV2
type GetOrderPaymentsV2Response struct {
	Response v2.OrderPaymentsResponse
	Error    error
}

// CreateOrderPaymentV2
type CreateOrderPaymentV2Request struct {
	Context RequestContext
	OrderId uint
	Request v2.OrderPaymentCreateRequest
}

// CreateOrderPaymentV2
type CreateOrderPaymentV2Response struct {
	Response v2.OrderPaymentMutateResponse
	Error    error
}

// UpdateOrderPaymentV2
type UpdateOrderPaymentV2Request struct {
	Context   RequestContext
	PaymentId uint
	Request   v2.OrderPaymentUpdateRequest
}

// UpdateOrderPaymentV2
type UpdateOrderPaymentV2Response struct {
	Response v2.OrderPaymentMutateResponse
	Error    error
}

// PatchOrderPaymentV2
type PatchOrderPaymentV2Request struct {
	Context   RequestContext
	PaymentId uint
	Request   v2.OrderPaymentUpdateRequest
}

// PatchOrderPaymentV2
type PatchOrderPaymentV2Response struct {
	Response v2.OrderPaymentMutateResponse
	Error    error
}

func (m *apiRPCClient) V2() v2.API {
	return newRPCV2API(m)
}

func (m *apiRPCServer) GetInvoicesV2(req GetInvoicesV2Request, resp *GetInvoicesV2Response) error {
	if impl, ok := m.impl.(v2.VersionedAPI); ok {
		data, err := impl.V2().Invoices().GetInvoices(req.Context, req.Request)
		resp.Response = data
		resp.Error = encodableError(err)
		return nil
	}
	legacyResponse, err := m.impl.GetInvoices(invoicesRequestV2ToLegacy(req.Request))
	resp.Response = invoicesListResponseToV2(legacyResponse)
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetInvoicesV2(context RequestContext, request v2.InvoicesRequest) (v2.InvoicesListResponse, error) {
	var reply GetInvoicesV2Response
	err := m.client.Call("Plugin.GetInvoicesV2", GetInvoicesV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.InvoicesListResponse{}, err
	}

	return reply.Response, decodableError(reply.Error)
}

func invoicesRequestV2ToLegacy(request v2.InvoicesRequest) InvoicesRequest {
	legacy := InvoicesRequest{
		UnlimitPerPage:          request.UnlimitPerPage,
		PerPage:                 request.PerPage,
		Page:                    request.Page,
		DateFrom:                request.DateFrom,
		DateTo:                  request.DateTo,
		SearchQuery:             request.SearchQuery,
		SelectOnlySales:         request.SelectOnlySales,
		SelectOnlyPurchases:     request.SelectOnlyPurchases,
		SelectOnlyWarehouseMove: request.SelectOnlyWarehouseMove,
		SelectOnlyDebtInvoices:  request.SelectOnlyDebtInvoices,
		FromWarehouses:          request.FromWarehouses,
		ToWarehouses:            request.ToWarehouses,
		Operations:              request.Operations,
		InvoiceId:               request.InvoiceIds,
		CompanyIds:              request.CompanyIds,
		ExcludeCompanyIds:       request.ExcludeCompanyIds,
		OrderWay:                request.OrderWay,
		Statuses:                request.Statuses,
	}
	if len(legacy.FromWarehouses) == 0 {
		legacy.FromWarehouses = request.Warehouses
	}
	switch request.ListKind {
	case InvoiceListKindSale:
		legacy.SelectOnlySales = true
	case InvoiceListKindPurchase:
		legacy.SelectOnlyPurchases = true
	case InvoiceListKindWarehouseMove:
		legacy.SelectOnlyWarehouseMove = true
	}
	return legacy
}

func invoicesListResponseToV2(response InvoicesListResponse) v2.InvoicesListResponse {
	invoices := make([]v2.Invoice, 0, len(response.Invoices))
	for _, invoice := range response.Invoices {
		invoices = append(invoices, invoiceFullToV2(invoice))
	}
	return v2.InvoicesListResponse{
		Invoices:   invoices,
		TotalCount: response.InvoicesCount,
	}
}

func invoiceFullToV2(invoice InvoiceFull) v2.Invoice {
	invoiceType := InvoiceTypeSale
	if invoice.IsPurchase {
		invoiceType = InvoiceTypePurchase
	}
	return v2.Invoice{
		InvoiceId:         invoice.InvoiceId,
		Type:              invoiceType,
		Operation:         invoice.Operation,
		CompanyName:       invoice.CompanyName,
		CompanyId:         invoice.CompanyId,
		InvoicePrefix:     invoice.DocSeries,
		InvoiceNumber:     invoice.Document,
		InvoiceNumber2:    invoice.Document2,
		InternalNumber:    invoice.InnerDocument,
		Date:              invoice.InvoiceDate,
		TotalTax:          invoice.TotalTax,
		TotalExclTax:      invoice.TotalExclTax,
		TotalInclTax:      invoice.TotalInclTax,
		Currency:          invoice.TotalCurrency,
		IsInvoicePaid:     invoice.InvoicePaid,
		TotalDebt:         invoice.DebtSum,
		TotalPaid:         invoice.PaidSum,
		PaidLastDate:      invoice.PaidLastDate,
		DueDate:           invoice.PayUntilDate,
		ReceiveDate:       invoice.ReceiveDate,
		TotalCost:         invoice.TotalCostSum,
		TotalMargin:       invoice.TotalGainSum,
		FromWarehouse:     invoice.FromWarehouse,
		ToWarehouse:       invoice.ToWarehouse,
		Comment:           invoice.Comment,
		TotalCompensation: invoice.TotalCompensation,
		CreatedAt:         invoice.InsertDatetime,
		UpdatedAt:         invoice.ModifyDatetime,
		CreatedUserInfo:   invoice.InsertUserInfo,
		UpdatedUserInfo:   invoice.UpdateUserInfo,
		IsVerified:        invoice.IsVerified,
		Status:            invoice.Status,
		CustomFields:      invoice.CustomFields,
	}
}

func (m *apiRPCServer) GetCompanyV2(req GetCompanyV2Request, resp *GetCompanyV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetCompanyV2"))
		return nil
	}
	company, err := impl.V2().Companies().GetCompany(req.Context, req.CompanyId)
	resp.Company = company
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetCompanyV2(context RequestContext, companyId uint) (v2.Company, error) {
	var reply GetCompanyV2Response
	err := m.client.Call("Plugin.GetCompanyV2", GetCompanyV2Request{
		Context:   context,
		CompanyId: companyId,
	}, &reply)
	if err != nil {
		return v2.Company{}, err
	}
	return reply.Company, decodableError(reply.Error)
}

func (m *apiRPCServer) GetCompaniesV2(req GetCompaniesV2Request, resp *GetCompaniesV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetCompaniesV2"))
		return nil
	}
	response, err := impl.V2().Companies().GetCompanies(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetCompaniesV2(context RequestContext, request v2.CompaniesRequest) (v2.CompaniesResponse, error) {
	var reply GetCompaniesV2Response
	err := m.client.Call("Plugin.GetCompaniesV2", GetCompaniesV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.CompaniesResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) CreateCompanyV2(req CreateCompanyV2Request, resp *CreateCompanyV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: CreateCompanyV2"))
		return nil
	}
	response, err := impl.V2().Companies().CreateCompany(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) CreateCompanyV2(context RequestContext, request v2.CompanyCreateUpdateRequest) (v2.CompanyMutateResponse, error) {
	var reply CreateCompanyV2Response
	err := m.client.Call("Plugin.CreateCompanyV2", CreateCompanyV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.CompanyMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) UpdateCompanyV2(req UpdateCompanyV2Request, resp *UpdateCompanyV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: UpdateCompanyV2"))
		return nil
	}
	response, err := impl.V2().Companies().UpdateCompany(req.Context, req.CompanyId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) UpdateCompanyV2(context RequestContext, companyId uint, request v2.CompanyCreateUpdateRequest) (v2.CompanyMutateResponse, error) {
	var reply UpdateCompanyV2Response
	err := m.client.Call("Plugin.UpdateCompanyV2", UpdateCompanyV2Request{
		Context:   context,
		CompanyId: companyId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.CompanyMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) PatchCompanyV2(req PatchCompanyV2Request, resp *PatchCompanyV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: PatchCompanyV2"))
		return nil
	}
	response, err := impl.V2().Companies().PatchCompany(req.Context, req.CompanyId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) PatchCompanyV2(context RequestContext, companyId uint, request v2.CompanyPatchRequest) (v2.CompanyMutateResponse, error) {
	var reply PatchCompanyV2Response
	err := m.client.Call("Plugin.PatchCompanyV2", PatchCompanyV2Request{
		Context:   context,
		CompanyId: companyId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.CompanyMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) GetProductV2(req GetProductV2Request, resp *GetProductV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetProductV2"))
		return nil
	}
	product, err := impl.V2().Products().GetProduct(req.Context, req.ProductId)
	resp.Product = product
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetProductV2(context RequestContext, productId uint) (v2.Product, error) {
	var reply GetProductV2Response
	err := m.client.Call("Plugin.GetProductV2", GetProductV2Request{
		Context:   context,
		ProductId: productId,
	}, &reply)
	if err != nil {
		return v2.Product{}, err
	}
	return reply.Product, decodableError(reply.Error)
}

func (m *apiRPCServer) GetProductsV2(req GetProductsV2Request, resp *GetProductsV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetProductsV2"))
		return nil
	}
	response, err := impl.V2().Products().GetProducts(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetProductsV2(context RequestContext, request v2.ProductsRequest) (v2.ProductsResponse, error) {
	var reply GetProductsV2Response
	err := m.client.Call("Plugin.GetProductsV2", GetProductsV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.ProductsResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) CreateProductV2(req CreateProductV2Request, resp *CreateProductV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: CreateProductV2"))
		return nil
	}
	response, err := impl.V2().Products().CreateProduct(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) CreateProductV2(context RequestContext, request v2.ProductCreateEditRequest) (v2.ProductMutateResponse, error) {
	var reply CreateProductV2Response
	err := m.client.Call("Plugin.CreateProductV2", CreateProductV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.ProductMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) UpdateProductV2(req UpdateProductV2Request, resp *UpdateProductV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: UpdateProductV2"))
		return nil
	}
	response, err := impl.V2().Products().UpdateProduct(req.Context, req.ProductId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) UpdateProductV2(context RequestContext, productId uint, request v2.ProductCreateEditRequest) (v2.ProductMutateResponse, error) {
	var reply UpdateProductV2Response
	err := m.client.Call("Plugin.UpdateProductV2", UpdateProductV2Request{
		Context:   context,
		ProductId: productId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.ProductMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) PatchProductV2(req PatchProductV2Request, resp *PatchProductV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: PatchProductV2"))
		return nil
	}
	response, err := impl.V2().Products().PatchProduct(req.Context, req.ProductId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) PatchProductV2(context RequestContext, productId uint, request v2.ProductPatchRequest) (v2.ProductMutateResponse, error) {
	var reply PatchProductV2Response
	err := m.client.Call("Plugin.PatchProductV2", PatchProductV2Request{
		Context:   context,
		ProductId: productId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.ProductMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) GetProductStocksV2(req GetProductStocksV2Request, resp *GetProductStocksV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetProductStocksV2"))
		return nil
	}
	response, err := impl.V2().Stocks().GetProductStocks(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetProductStocksV2(context RequestContext, request v2.ProductStocksRequest) (v2.ProductStocksResponse, error) {
	var reply GetProductStocksV2Response
	err := m.client.Call("Plugin.GetProductStocksV2", GetProductStocksV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.ProductStocksResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) GetProductStocksWarehouses(req GetProductStocksWarehousesRequest, resp *GetProductStocksWarehousesResponse) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetProductStocksWarehouses"))
		return nil
	}
	response, err := impl.V2().Stocks().GetProductStocksWarehouses(req.Context, req.ProductId, req.FilterAvailable)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetProductStocksWarehouses(context RequestContext, productId uint, filterAvailable bool) (v2.ProductWarehousesQuantitiesResponse, error) {
	var reply GetProductStocksWarehousesResponse
	err := m.client.Call("Plugin.GetProductStocksWarehouses", GetProductStocksWarehousesRequest{
		Context:         context,
		ProductId:       productId,
		FilterAvailable: filterAvailable,
	}, &reply)
	if err != nil {
		return v2.ProductWarehousesQuantitiesResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) SetProductStocks(req SetProductStocksRequest, resp *SetProductStocksResponse) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: SetProductStocks"))
		return nil
	}
	response, err := impl.V2().Stocks().SetProductStocks(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) SetProductStocks(context RequestContext, request v2.ProductStockSetRequest) (v2.ProductStockAdjustmentResponse, error) {
	var reply SetProductStocksResponse
	err := m.client.Call("Plugin.SetProductStocks", SetProductStocksRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.ProductStockAdjustmentResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) AdjustProductStocks(req AdjustProductStocksRequest, resp *AdjustProductStocksResponse) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: AdjustProductStocks"))
		return nil
	}
	response, err := impl.V2().Stocks().AdjustProductStocks(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) AdjustProductStocks(context RequestContext, request v2.ProductStockAdjustRequest) (v2.ProductStockAdjustmentResponse, error) {
	var reply AdjustProductStocksResponse
	err := m.client.Call("Plugin.AdjustProductStocks", AdjustProductStocksRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.ProductStockAdjustmentResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) GetOrderV2(req GetOrderV2Request, resp *GetOrderV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetOrderV2"))
		return nil
	}
	order, err := impl.V2().Orders().GetOrder(req.Context, req.OrderId)
	resp.Order = order
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetOrderV2(context RequestContext, orderId uint) (v2.OrderDetail, error) {
	var reply GetOrderV2Response
	err := m.client.Call("Plugin.GetOrderV2", GetOrderV2Request{
		Context: context,
		OrderId: orderId,
	}, &reply)
	if err != nil {
		return v2.OrderDetail{}, err
	}
	return reply.Order, decodableError(reply.Error)
}

func (m *apiRPCServer) GetOrdersV2(req GetOrdersV2Request, resp *GetOrdersV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetOrdersV2"))
		return nil
	}
	response, err := impl.V2().Orders().GetOrders(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetOrdersV2(context RequestContext, request v2.OrdersRequest) (v2.OrdersResponse, error) {
	var reply GetOrdersV2Response
	err := m.client.Call("Plugin.GetOrdersV2", GetOrdersV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrdersResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) CreateOrderV2(req CreateOrderV2Request, resp *CreateOrderV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: CreateOrderV2"))
		return nil
	}
	response, err, errResponse := impl.V2().Orders().CreateOrder(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	resp.ErrorResponse = errResponse
	return nil
}

func (m *apiRPCClient) CreateOrderV2(context RequestContext, request v2.OrderCreateRequest) (v2.OrderMutateResponse, error, OrderErrorResponse) {
	var reply CreateOrderV2Response
	err := m.client.Call("Plugin.CreateOrderV2", CreateOrderV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderMutateResponse{}, err, OrderErrorResponse{}
	}
	return reply.Response, decodableError(reply.Error), reply.ErrorResponse
}

func (m *apiRPCServer) UpdateOrderV2(req UpdateOrderV2Request, resp *UpdateOrderV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: UpdateOrderV2"))
		return nil
	}
	response, err, errResponse := impl.V2().Orders().UpdateOrder(req.Context, req.OrderId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	resp.ErrorResponse = errResponse
	return nil
}

func (m *apiRPCClient) UpdateOrderV2(context RequestContext, orderId uint, request v2.OrderUpdateRequest) (v2.OrderMutateResponse, error, OrderErrorResponse) {
	var reply UpdateOrderV2Response
	err := m.client.Call("Plugin.UpdateOrderV2", UpdateOrderV2Request{
		Context: context,
		OrderId: orderId,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderMutateResponse{}, err, OrderErrorResponse{}
	}
	return reply.Response, decodableError(reply.Error), reply.ErrorResponse
}

func (m *apiRPCServer) PatchOrderV2(req PatchOrderV2Request, resp *PatchOrderV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: PatchOrderV2"))
		return nil
	}
	response, err := impl.V2().Orders().PatchOrder(req.Context, req.OrderId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) PatchOrderV2(context RequestContext, orderId uint, request v2.OrderPatchRequest) (v2.OrderMutateResponse, error) {
	var reply PatchOrderV2Response
	err := m.client.Call("Plugin.PatchOrderV2", PatchOrderV2Request{
		Context: context,
		OrderId: orderId,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) CreateInvoiceFromOrderV2(req CreateInvoiceFromOrderV2Request, resp *CreateInvoiceFromOrderV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: CreateInvoiceFromOrderV2"))
		return nil
	}
	response, err, errResponse := impl.V2().Orders().CreateInvoiceFromOrder(req.Context, req.OrderId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	resp.ErrorResponse = errResponse
	return nil
}

func (m *apiRPCClient) CreateInvoiceFromOrderV2(context RequestContext, orderId uint, request v2.OrderInvoiceCreateRequest) (v2.OrderInvoiceCreateResponse, error, InvoiceErrorResponse) {
	var reply CreateInvoiceFromOrderV2Response
	err := m.client.Call("Plugin.CreateInvoiceFromOrderV2", CreateInvoiceFromOrderV2Request{
		Context: context,
		OrderId: orderId,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderInvoiceCreateResponse{}, err, InvoiceErrorResponse{}
	}
	return reply.Response, decodableError(reply.Error), reply.ErrorResponse
}

func (m *apiRPCServer) GetOrderPaymentV2(req GetOrderPaymentV2Request, resp *GetOrderPaymentV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetOrderPaymentV2"))
		return nil
	}
	payment, err := impl.V2().OrderPayments().GetOrderPayment(req.Context, req.PaymentId)
	resp.Payment = payment
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetOrderPaymentV2(context RequestContext, paymentId uint) (v2.OrderPayment, error) {
	var reply GetOrderPaymentV2Response
	err := m.client.Call("Plugin.GetOrderPaymentV2", GetOrderPaymentV2Request{
		Context:   context,
		PaymentId: paymentId,
	}, &reply)
	if err != nil {
		return v2.OrderPayment{}, err
	}
	return reply.Payment, decodableError(reply.Error)
}

func (m *apiRPCServer) GetOrderPaymentsV2(req GetOrderPaymentsV2Request, resp *GetOrderPaymentsV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: GetOrderPaymentsV2"))
		return nil
	}
	response, err := impl.V2().OrderPayments().GetOrderPayments(req.Context, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetOrderPaymentsV2(context RequestContext, request v2.OrderPaymentsRequest) (v2.OrderPaymentsResponse, error) {
	var reply GetOrderPaymentsV2Response
	err := m.client.Call("Plugin.GetOrderPaymentsV2", GetOrderPaymentsV2Request{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderPaymentsResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) CreateOrderPaymentV2(req CreateOrderPaymentV2Request, resp *CreateOrderPaymentV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: CreateOrderPaymentV2"))
		return nil
	}
	response, err := impl.V2().OrderPayments().CreateOrderPayment(req.Context, req.OrderId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) CreateOrderPaymentV2(context RequestContext, orderId uint, request v2.OrderPaymentCreateRequest) (v2.OrderPaymentMutateResponse, error) {
	var reply CreateOrderPaymentV2Response
	err := m.client.Call("Plugin.CreateOrderPaymentV2", CreateOrderPaymentV2Request{
		Context: context,
		OrderId: orderId,
		Request: request,
	}, &reply)
	if err != nil {
		return v2.OrderPaymentMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) UpdateOrderPaymentV2(req UpdateOrderPaymentV2Request, resp *UpdateOrderPaymentV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: UpdateOrderPaymentV2"))
		return nil
	}
	response, err := impl.V2().OrderPayments().UpdateOrderPayment(req.Context, req.PaymentId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) UpdateOrderPaymentV2(context RequestContext, paymentId uint, request v2.OrderPaymentUpdateRequest) (v2.OrderPaymentMutateResponse, error) {
	var reply UpdateOrderPaymentV2Response
	err := m.client.Call("Plugin.UpdateOrderPaymentV2", UpdateOrderPaymentV2Request{
		Context:   context,
		PaymentId: paymentId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.OrderPaymentMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}

func (m *apiRPCServer) PatchOrderPaymentV2(req PatchOrderPaymentV2Request, resp *PatchOrderPaymentV2Response) error {
	impl, ok := m.impl.(v2.VersionedAPI)
	if !ok {
		resp.Error = encodableError(errors.New("host does not implement plugin API v2: PatchOrderPaymentV2"))
		return nil
	}
	response, err := impl.V2().OrderPayments().PatchOrderPayment(req.Context, req.PaymentId, req.Request)
	resp.Response = response
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) PatchOrderPaymentV2(context RequestContext, paymentId uint, request v2.OrderPaymentUpdateRequest) (v2.OrderPaymentMutateResponse, error) {
	var reply PatchOrderPaymentV2Response
	err := m.client.Call("Plugin.PatchOrderPaymentV2", PatchOrderPaymentV2Request{
		Context:   context,
		PaymentId: paymentId,
		Request:   request,
	}, &reply)
	if err != nil {
		return v2.OrderPaymentMutateResponse{}, err
	}
	return reply.Response, decodableError(reply.Error)
}
