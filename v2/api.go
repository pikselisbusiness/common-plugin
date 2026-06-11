// Package v2 contains the public v2 plugin host API.
//
// It provides clean DTO names and domain-scoped API clients.
package v2

import (
	"errors"

	common "github.com/pikselisbusiness/common-plugin/common"
)

var ErrUnsupported = errors.New("host does not expose plugin API v2")

type RequestContext = common.RequestContext

type API interface {
	Invoices() InvoicesAPI
	Companies() CompaniesAPI
	Products() ProductsAPI
	Stocks() StocksAPI
	Orders() OrdersAPI
	OrderPayments() OrderPaymentsAPI
}

type VersionedAPI interface {
	V2() API
}

type InvoicesAPI interface {
	GetInvoices(context RequestContext, request InvoicesRequest) (InvoicesListResponse, error)
}

type CompaniesAPI interface {
	GetCompany(context RequestContext, companyId uint) (Company, error)
	GetCompanies(context RequestContext, request CompaniesRequest) (CompaniesResponse, error)
	CreateCompany(context RequestContext, request CompanyCreateUpdateRequest) (CompanyMutateResponse, error)
	UpdateCompany(context RequestContext, companyId uint, request CompanyCreateUpdateRequest) (CompanyMutateResponse, error)
	PatchCompany(context RequestContext, companyId uint, request CompanyPatchRequest) (CompanyMutateResponse, error)
}

type ProductsAPI interface {
	GetProduct(context RequestContext, productId uint) (Product, error)
	GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error)
	CreateProduct(context RequestContext, request ProductCreateEditRequest) (ProductMutateResponse, error)
	UpdateProduct(context RequestContext, productId uint, request ProductCreateEditRequest) (ProductMutateResponse, error)
	PatchProduct(context RequestContext, productId uint, request ProductPatchRequest) (ProductMutateResponse, error)
}

type StocksAPI interface {
	GetProductStocks(context RequestContext, request ProductStocksRequest) (ProductStocksResponse, error)
	GetProductStocksWarehouses(context RequestContext, productId uint, filterAvailable bool) (ProductWarehousesQuantitiesResponse, error)
	SetProductStocks(context RequestContext, request ProductStockSetRequest) (ProductStockAdjustmentResponse, error)
	AdjustProductStocks(context RequestContext, request ProductStockAdjustRequest) (ProductStockAdjustmentResponse, error)
}

type OrdersAPI interface {
	GetOrder(context RequestContext, orderId uint) (OrderDetail, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
	CreateOrder(context RequestContext, request OrderCreateRequest) (OrderMutateResponse, error, OrderErrorResponse)
	UpdateOrder(context RequestContext, orderId uint, request OrderUpdateRequest) (OrderMutateResponse, error, OrderErrorResponse)
	PatchOrder(context RequestContext, orderId uint, request OrderPatchRequest) (OrderMutateResponse, error)
	CreateInvoiceFromOrder(context RequestContext, orderId uint, request OrderInvoiceCreateRequest) (OrderInvoiceCreateResponse, error, InvoiceErrorResponse)
}

type OrderPaymentsAPI interface {
	GetOrderPayment(context RequestContext, paymentId uint) (OrderPayment, error)
	GetOrderPayments(context RequestContext, request OrderPaymentsRequest) (OrderPaymentsResponse, error)
	CreateOrderPayment(context RequestContext, orderId uint, request OrderPaymentCreateRequest) (OrderPaymentMutateResponse, error)
	UpdateOrderPayment(context RequestContext, paymentId uint, request OrderPaymentUpdateRequest) (OrderPaymentMutateResponse, error)
	PatchOrderPayment(context RequestContext, paymentId uint, request OrderPaymentUpdateRequest) (OrderPaymentMutateResponse, error)
}

// From returns the v2 domain API from a plugin host API instance.
func From(api any) (API, bool) {
	if versioned, ok := api.(VersionedAPI); ok {
		return versioned.V2(), true
	}
	return nil, false
}

func MustFrom(api any) (API, error) {
	v2API, ok := From(api)
	if !ok {
		return nil, ErrUnsupported
	}
	return v2API, nil
}

type UserInfo = common.UserInfo
type ItemFilter = common.ItemFilter
type PriceRange = common.PriceRange

type InvoiceListKind = common.InvoiceListKind

const (
	InvoiceListKindSale          = common.InvoiceListKindSale
	InvoiceListKindPurchase      = common.InvoiceListKindPurchase
	InvoiceListKindWarehouseMove = common.InvoiceListKindWarehouseMove
	InvoiceListKindAll           = common.InvoiceListKindAll
)

type OrderErrorResponse = common.OrderErrorResponse
type InvoiceErrorResponse = common.InvoiceErrorResponse
