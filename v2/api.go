// Package v2 contains the public v2 plugin host API.
//
// It provides clean DTO names and domain-scoped API clients.
//
// This package is pre-stable until the first public v2 release tag. Field names
// and DTO shapes may still change while the host and common plugin package are
// being aligned.
package v2

import (
	"errors"

	common "github.com/pikselisbusiness/common-plugin/common"
)

var ErrUnsupported = errors.New("host does not expose plugin API v2")

type RequestContext = common.RequestContext

type API interface {
	Invoices() InvoicesAPI
	JournalEntries() JournalEntriesAPI
	Companies() CompaniesAPI
	Products() ProductsAPI
	Stocks() StocksAPI
	Orders() OrdersAPI
	OrderPayments() OrderPaymentsAPI
	ReferenceData() ReferenceDataAPI
	CommerceChannels() CommerceChannelsAPI
	Sync() SyncAPI
	CustomFields() CustomFieldsAPI
	Attachments() AttachmentsAPI
	Email() EmailAPI
	Issues() IssuesAPI
	Extensions() ExtensionsAPI
}

type VersionedAPI interface {
	V2() API
}

type InvoicesAPI interface {
	GetInvoices(context RequestContext, request InvoicesRequest) (InvoicesListResponse, error)
	GetInvoiceLines(context RequestContext, invoiceId uint) ([]InvoiceProduct, error)
	GetInvoiceReferences(context RequestContext, invoiceId uint) ([]InvoiceReference, error)
	CreateInvoice(context RequestContext, request InvoiceCreateUpdateRequest) (uint, error, InvoiceErrorResponse)
	InvoiceExistsByDocument(context RequestContext, document string) (bool, error)
	GetInvoicePayments(context RequestContext, request InvoicePaymentsRequest) (InvoicePaymentsResponse, error)
	PatchInvoice(context RequestContext, invoiceId uint, request map[string]interface{}) error
	CreateInvoiceReference(context RequestContext, request InvoiceReference) error
}

type JournalEntriesAPI interface {
	GetJournalEntries(context RequestContext, request JournalEntriesRequest) (JournalEntriesResponse, error)
	GetJournalEntry(context RequestContext, journalEntryId uint) (JournalEntry, error)
}

type CompaniesAPI interface {
	GetCompany(context RequestContext, companyId uint) (Company, error)
	GetCompanies(context RequestContext, request CompaniesRequest) (CompaniesResponse, error)
	GetCompaniesByID(context RequestContext, companyIds []uint) (map[uint]Company, error)
	CreateCompany(context RequestContext, request CompanyCreateUpdateRequest) (CompanyMutateResponse, error)
	UpdateCompany(context RequestContext, companyId uint, request CompanyCreateUpdateRequest) (CompanyMutateResponse, error)
	PatchCompany(context RequestContext, companyId uint, request CompanyPatchRequest) (CompanyMutateResponse, error)
}

type ProductsAPI interface {
	GetProductFeatureValues(context RequestContext, productId uint) ([]ProductFeatureValue, error)
	GetProduct(context RequestContext, productId uint) (Product, error)
	GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error)
	GetProductsByID(context RequestContext, productIds []uint) (map[uint]Product, error)
	GetProductCategories(context RequestContext, request ProductCategoriesRequest) (ProductCategoriesResponse, error)
	CreateProduct(context RequestContext, request ProductCreateEditRequest) (ProductMutateResponse, error)
	UpdateProduct(context RequestContext, productId uint, request ProductCreateEditRequest) (ProductMutateResponse, error)
	PatchProduct(context RequestContext, productId uint, request ProductPatchRequest) (ProductMutateResponse, error)

	// UpsertFeature idempotently creates or updates a feature (attribute definition)
	// by name and returns its id. ReplaceFeatures matches products by SKU and replaces
	// each matched product's assignments per feature.
	UpsertFeature(context RequestContext, request UpsertFeatureRequest) (UpsertFeatureResponse, error)
	ReplaceFeatures(context RequestContext, request ReplaceProductFeaturesRequest) (ReplaceProductFeaturesResult, error)
}

type StocksAPI interface {
	GetProductStocks(context RequestContext, request ProductStocksRequest) (ProductStocksResponse, error)
	GetProductStocksWarehouses(context RequestContext, productId uint, filterAvailable bool) (ProductWarehousesQuantitiesResponse, error)
	GetAvailableProductQuantity(context RequestContext, productId uint, warehouses []string) (float64, error)
	SetProductStocks(context RequestContext, request ProductStockSetRequest) (ProductStockAdjustmentResponse, error)
	AdjustProductStocks(context RequestContext, request ProductStockAdjustRequest) (ProductStockAdjustmentResponse, error)
}

type OrdersAPI interface {
	GetOrder(context RequestContext, orderId uint) (OrderDetail, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
	GetOrderStatuses(context RequestContext) ([]OrderStatus, error)
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

type ReferenceDataAPI interface {
	GetDocumentOperations(context RequestContext, request DocumentOperationsRequest) (DocumentOperationsResponse, error)
	GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error)
	GetCountryByName(context RequestContext, name string) (Country, error)
	GetAllCountries(context RequestContext) ([]Country, error)
}

type CommerceChannelsAPI interface {
	GetCommerceChannelByUUID(context RequestContext, channelUUID string) (CommerceChannel, error)
	CreateCommerceChannel(context RequestContext, channel CommerceChannel) (uint, error)
}

type SyncAPI interface {
	CreateIntegrationSyncRecord(context RequestContext, sync IntegrationSyncRecord) (uint, error)
	DeleteIntegrationSyncRecordById(context RequestContext, syncRecordId uint) error
	GetIntegrationSyncRecords(context RequestContext, request IntegrationSyncRecordsRequest) ([]IntegrationSyncRecord, error)
	UpsertPluginSyncState(context RequestContext, state PluginSyncState) error
	ListPluginSyncStates(context RequestContext, request PluginSyncStateRequest) (PluginSyncStateResponse, error)
	CountPluginSyncStates(context RequestContext, request PluginSyncStateRequest) (int64, error)
}

type CustomFieldsAPI interface {
	GetCustomFields(context RequestContext, request CustomFieldsRequest) (CustomFieldsResponse, error)
	CreateCustomField(context RequestContext, request CustomField) (uint, error)
}

type AttachmentsAPI interface {
	CreateAttachment(context RequestContext, request AttachmentCreateRequest) (AttachmentCreateResponse, error)
	DeleteAttachment(context RequestContext, attachmentId uint) error
}

type EmailAPI interface {
	SendEmail(context RequestContext, request EmailRequest) error
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
type DocumentOperation = common.DocumentOperation
type DocumentOperationsRequest = common.DocumentOperationsRequest
type DocumentOperationsResponse = common.DocumentOperationsResponse
type InvoiceProduct = common.InvoiceProduct
type InvoiceReference = common.InvoiceReference
type InvoicePaymentsRequest = common.InvoicePaymentsRequest
type InvoicePayment = common.InvoicePayment
type InvoicePaymentsResponse = common.InvoicePaymentsResponse
type InvoiceSend = common.InvoiceSend
type InvoiceCreateUpdateRequest = common.InvoiceCreateUpdateRequest
type ExtraInvoiceSettings = common.ExtraInvoiceSettings
type InvoiceCost = common.InvoiceCost
type InvoicePosData = common.InvoicePosData
type JournalEntriesRequest = common.JournalEntriesRequest
type JournalEntryLine = common.JournalEntryLine
type JournalEntry = common.JournalEntry
type JournalEntriesResponse = common.JournalEntriesResponse
type ProductCategoriesShort = common.ProductCategoriesShort
type ProductCategoriesRequest = common.ProductCategoriesRequest
type ProductCategoriesResponse = common.ProductCategoriesResponse

// Product feature sync DTOs (shared with the v1 API via the common package).
type UpsertFeatureRequest = common.UpsertFeatureRequest
type UpsertFeatureResponse = common.UpsertFeatureResponse
type FeatureValueItem = common.FeatureValueItem
type ProductFeatureSet = common.ProductFeatureSet
type ProductFeatureSyncItem = common.ProductFeatureSyncItem
type ReplaceProductFeaturesRequest = common.ReplaceProductFeaturesRequest
type ReplaceProductFeaturesResult = common.ReplaceProductFeaturesResult
type DivisionsRequest = common.DivisionsRequest
type Division = common.Division
type DivisionWarehouse = common.DivisionWarehouse
type Country = common.Country
type CommerceChannel = common.CommerceChannel
type IntegrationSyncRecord = common.IntegrationSyncRecord
type IntegrationSyncRecordsRequest = common.IntegrationSyncRecordsRequest
type PluginSyncState = common.PluginSyncState
type PluginSyncStateRequest = common.PluginSyncStateRequest
type PluginSyncStateResponse = common.PluginSyncStateResponse
type CustomFieldType = common.CustomFieldType
type CustomField = common.CustomField
type CustomFieldsRequest = common.CustomFieldsRequest
type CustomFieldsResponse = common.CustomFieldsResponse
type AttachmentCreateRequest = common.FileUploadRequest
type AttachmentCreateResponse = common.FileUploadResponse
type FileUploadRequest = common.FileUploadRequest
type FileUploadResponse = common.FileUploadResponse
type EmailSender = common.EmailSender
type EmailRequest = common.EmailRequest

var ValidCustomFieldTypes = common.ValidCustomFieldTypes

var (
	CustomFieldString   = common.CustomFieldString
	CustomFieldText     = common.CustomFieldText
	CustomFieldFloat    = common.CustomFieldFloat
	CustomFieldDate     = common.CustomFieldDate
	CustomFieldDatetime = common.CustomFieldDatetime
	CustomFieldBoolean  = common.CustomFieldBoolean
	CustomFieldSelect   = common.CustomFieldSelect
)

type InvoiceListKind = common.InvoiceListKind

const (
	InvoiceListKindSale          = common.InvoiceListKindSale
	InvoiceListKindPurchase      = common.InvoiceListKindPurchase
	InvoiceListKindWarehouseMove = common.InvoiceListKindWarehouseMove
	InvoiceListKindAll           = common.InvoiceListKindAll
)

type OrderErrorResponse = common.OrderErrorResponse
type InvoiceErrorResponse = common.InvoiceErrorResponse
