package shared

import v2 "github.com/pikselisbusiness/common-plugin/v2"

func newRPCV2API(api *apiRPCClient) v2.API {
	return rpcV2API{api: api}
}

type rpcV2API struct {
	api *apiRPCClient
}

func (a rpcV2API) Invoices() v2.InvoicesAPI {
	return rpcInvoicesV2API{api: a.api}
}

func (a rpcV2API) Companies() v2.CompaniesAPI {
	return rpcCompaniesV2API{api: a.api}
}

func (a rpcV2API) Products() v2.ProductsAPI {
	return rpcProductsV2API{api: a.api}
}

func (a rpcV2API) Stocks() v2.StocksAPI {
	return rpcStocksV2API{api: a.api}
}

func (a rpcV2API) Orders() v2.OrdersAPI {
	return rpcOrdersV2API{api: a.api}
}

func (a rpcV2API) OrderPayments() v2.OrderPaymentsAPI {
	return rpcOrderPaymentsV2API{api: a.api}
}

type rpcInvoicesV2API struct {
	api *apiRPCClient
}

func (a rpcInvoicesV2API) GetInvoices(context RequestContext, request v2.InvoicesRequest) (v2.InvoicesListResponse, error) {
	return a.api.GetInvoicesV2(context, request)
}

type rpcCompaniesV2API struct {
	api *apiRPCClient
}

func (a rpcCompaniesV2API) GetCompany(context RequestContext, companyId uint) (v2.Company, error) {
	return a.api.GetCompanyV2(context, companyId)
}

func (a rpcCompaniesV2API) GetCompanies(context RequestContext, request v2.CompaniesRequest) (v2.CompaniesResponse, error) {
	return a.api.GetCompaniesV2(context, request)
}

func (a rpcCompaniesV2API) CreateCompany(context RequestContext, request v2.CompanyCreateUpdateRequest) (v2.CompanyMutateResponse, error) {
	return a.api.CreateCompanyV2(context, request)
}

func (a rpcCompaniesV2API) UpdateCompany(context RequestContext, companyId uint, request v2.CompanyCreateUpdateRequest) (v2.CompanyMutateResponse, error) {
	return a.api.UpdateCompanyV2(context, companyId, request)
}

func (a rpcCompaniesV2API) PatchCompany(context RequestContext, companyId uint, request v2.CompanyPatchRequest) (v2.CompanyMutateResponse, error) {
	return a.api.PatchCompanyV2(context, companyId, request)
}

type rpcProductsV2API struct {
	api *apiRPCClient
}

func (a rpcProductsV2API) GetProduct(context RequestContext, productId uint) (v2.Product, error) {
	return a.api.GetProductV2(context, productId)
}

func (a rpcProductsV2API) GetProducts(context RequestContext, request v2.ProductsRequest) (v2.ProductsResponse, error) {
	return a.api.GetProductsV2(context, request)
}

func (a rpcProductsV2API) CreateProduct(context RequestContext, request v2.ProductCreateEditRequest) (v2.ProductMutateResponse, error) {
	return a.api.CreateProductV2(context, request)
}

func (a rpcProductsV2API) UpdateProduct(context RequestContext, productId uint, request v2.ProductCreateEditRequest) (v2.ProductMutateResponse, error) {
	return a.api.UpdateProductV2(context, productId, request)
}

func (a rpcProductsV2API) PatchProduct(context RequestContext, productId uint, request v2.ProductPatchRequest) (v2.ProductMutateResponse, error) {
	return a.api.PatchProductV2(context, productId, request)
}

type rpcStocksV2API struct {
	api *apiRPCClient
}

func (a rpcStocksV2API) GetProductStocks(context RequestContext, request v2.ProductStocksRequest) (v2.ProductStocksResponse, error) {
	return a.api.GetProductStocksV2(context, request)
}

func (a rpcStocksV2API) GetProductStocksWarehouses(context RequestContext, productId uint, filterAvailable bool) (v2.ProductWarehousesQuantitiesResponse, error) {
	return a.api.GetProductStocksWarehouses(context, productId, filterAvailable)
}

func (a rpcStocksV2API) SetProductStocks(context RequestContext, request v2.ProductStockSetRequest) (v2.ProductStockAdjustmentResponse, error) {
	return a.api.SetProductStocks(context, request)
}

func (a rpcStocksV2API) AdjustProductStocks(context RequestContext, request v2.ProductStockAdjustRequest) (v2.ProductStockAdjustmentResponse, error) {
	return a.api.AdjustProductStocks(context, request)
}

type rpcOrdersV2API struct {
	api *apiRPCClient
}

func (a rpcOrdersV2API) GetOrder(context RequestContext, orderId uint) (v2.OrderDetail, error) {
	return a.api.GetOrderV2(context, orderId)
}

func (a rpcOrdersV2API) GetOrders(context RequestContext, request v2.OrdersRequest) (v2.OrdersResponse, error) {
	return a.api.GetOrdersV2(context, request)
}

func (a rpcOrdersV2API) CreateOrder(context RequestContext, request v2.OrderCreateRequest) (v2.OrderMutateResponse, error, OrderErrorResponse) {
	return a.api.CreateOrderV2(context, request)
}

func (a rpcOrdersV2API) UpdateOrder(context RequestContext, orderId uint, request v2.OrderUpdateRequest) (v2.OrderMutateResponse, error, OrderErrorResponse) {
	return a.api.UpdateOrderV2(context, orderId, request)
}

func (a rpcOrdersV2API) PatchOrder(context RequestContext, orderId uint, request v2.OrderPatchRequest) (v2.OrderMutateResponse, error) {
	return a.api.PatchOrderV2(context, orderId, request)
}

func (a rpcOrdersV2API) CreateInvoiceFromOrder(context RequestContext, orderId uint, request v2.OrderInvoiceCreateRequest) (v2.OrderInvoiceCreateResponse, error, InvoiceErrorResponse) {
	return a.api.CreateInvoiceFromOrderV2(context, orderId, request)
}

type rpcOrderPaymentsV2API struct {
	api *apiRPCClient
}

func (a rpcOrderPaymentsV2API) GetOrderPayment(context RequestContext, paymentId uint) (v2.OrderPayment, error) {
	return a.api.GetOrderPaymentV2(context, paymentId)
}

func (a rpcOrderPaymentsV2API) GetOrderPayments(context RequestContext, request v2.OrderPaymentsRequest) (v2.OrderPaymentsResponse, error) {
	return a.api.GetOrderPaymentsV2(context, request)
}

func (a rpcOrderPaymentsV2API) CreateOrderPayment(context RequestContext, orderId uint, request v2.OrderPaymentCreateRequest) (v2.OrderPaymentMutateResponse, error) {
	return a.api.CreateOrderPaymentV2(context, orderId, request)
}

func (a rpcOrderPaymentsV2API) UpdateOrderPayment(context RequestContext, paymentId uint, request v2.OrderPaymentUpdateRequest) (v2.OrderPaymentMutateResponse, error) {
	return a.api.UpdateOrderPaymentV2(context, paymentId, request)
}

func (a rpcOrderPaymentsV2API) PatchOrderPayment(context RequestContext, paymentId uint, request v2.OrderPaymentUpdateRequest) (v2.OrderPaymentMutateResponse, error) {
	return a.api.PatchOrderPaymentV2(context, paymentId, request)
}
