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

func (a rpcV2API) JournalEntries() v2.JournalEntriesAPI {
	return rpcJournalEntriesV2API{api: a.api}
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

func (a rpcV2API) ReferenceData() v2.ReferenceDataAPI {
	return rpcReferenceDataV2API{api: a.api}
}

func (a rpcV2API) CommerceChannels() v2.CommerceChannelsAPI {
	return rpcCommerceChannelsV2API{api: a.api}
}

func (a rpcV2API) Sync() v2.SyncAPI {
	return rpcSyncV2API{api: a.api}
}

func (a rpcV2API) CustomFields() v2.CustomFieldsAPI {
	return rpcCustomFieldsV2API{api: a.api}
}

func (a rpcV2API) Attachments() v2.AttachmentsAPI {
	return rpcAttachmentsV2API{api: a.api}
}

func (a rpcV2API) Email() v2.EmailAPI {
	return rpcEmailV2API{api: a.api}
}

func (a rpcV2API) Issues() v2.IssuesAPI {
	return rpcIssuesV2API{api: a.api}
}

func (a rpcV2API) Extensions() v2.ExtensionsAPI {
	return rpcExtensionsV2API{api: a.api}
}

func (a rpcV2API) Repairs() v2.RepairsAPI {
	return rpcRepairsV2API{api: a.api}
}

type rpcInvoicesV2API struct {
	api *apiRPCClient
}

func (a rpcInvoicesV2API) GetInvoices(context RequestContext, request v2.InvoicesRequest) (v2.InvoicesListResponse, error) {
	return a.api.GetInvoicesV2(context, request)
}

func (a rpcInvoicesV2API) GetInvoiceLines(context RequestContext, invoiceId uint) ([]v2.InvoiceProduct, error) {
	return a.api.GetInvoiceProducts(context, invoiceId)
}

func (a rpcInvoicesV2API) GetInvoiceReferences(context RequestContext, invoiceId uint) ([]v2.InvoiceReference, error) {
	return a.api.GetInvoiceReferences(context, invoiceId)
}

func (a rpcInvoicesV2API) CreateInvoice(context RequestContext, request v2.InvoiceCreateUpdateRequest) (uint, error, v2.InvoiceErrorResponse) {
	return a.api.CreateInvoice(context, request)
}

func (a rpcInvoicesV2API) InvoiceExistsByDocument(context RequestContext, document string) (bool, error) {
	return a.api.GetInvoiceExistsByDocument(context, document)
}

func (a rpcInvoicesV2API) GetInvoicePayments(context RequestContext, request v2.InvoicePaymentsRequest) (v2.InvoicePaymentsResponse, error) {
	return a.api.GetInvoicePayments(context, request)
}

func (a rpcInvoicesV2API) PatchInvoice(context RequestContext, invoiceId uint, request map[string]interface{}) error {
	return a.api.PatchUpdateInvoice(context, invoiceId, request)
}

func (a rpcInvoicesV2API) CreateInvoiceReference(context RequestContext, request v2.InvoiceReference) error {
	return a.api.CreateInvoiceReference(context, request)
}

type rpcJournalEntriesV2API struct {
	api *apiRPCClient
}

func (a rpcJournalEntriesV2API) GetJournalEntries(context RequestContext, request v2.JournalEntriesRequest) (v2.JournalEntriesResponse, error) {
	return a.api.GetJournalEntries(context, request)
}

func (a rpcJournalEntriesV2API) GetJournalEntry(context RequestContext, journalEntryId uint) (v2.JournalEntry, error) {
	return a.api.GetJournalEntryById(context, journalEntryId)
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

func (a rpcCompaniesV2API) GetCompaniesByID(context RequestContext, companyIds []uint) (map[uint]v2.Company, error) {
	companies := make(map[uint]v2.Company, len(companyIds))
	for _, companyId := range companyIds {
		company, err := a.api.GetCompanyV2(context, companyId)
		if err != nil {
			return nil, err
		}
		companies[companyId] = company
	}
	return companies, nil
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

func (a rpcProductsV2API) GetProductsByID(context RequestContext, productIds []uint) (map[uint]v2.Product, error) {
	products := make(map[uint]v2.Product, len(productIds))
	for _, productId := range productIds {
		product, err := a.api.GetProductV2(context, productId)
		if err != nil {
			return nil, err
		}
		products[productId] = product
	}
	return products, nil
}

func (a rpcProductsV2API) GetProductCategories(context RequestContext, request v2.ProductCategoriesRequest) (v2.ProductCategoriesResponse, error) {
	return a.api.GetProductCategories(context, request)
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

// Feature methods reuse the v1 RPC transport — the DTOs are identical (aliases to
// the common package), so no dedicated *V2 RPC stub is needed.
func (a rpcProductsV2API) UpsertFeature(context RequestContext, request v2.UpsertFeatureRequest) (v2.UpsertFeatureResponse, error) {
	return a.api.UpsertProductFeature(context, request)
}
func (a rpcProductsV2API) ReplaceFeatures(context RequestContext, request v2.ReplaceProductFeaturesRequest) (v2.ReplaceProductFeaturesResult, error) {
	return a.api.ReplaceProductFeatures(context, request)
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

func (a rpcStocksV2API) GetAvailableProductQuantity(context RequestContext, productId uint, warehouses []string) (float64, error) {
	return a.api.GetAvailableProductQuantityByWarehouses(context, productId, warehouses)
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

func (a rpcOrdersV2API) GetOrderStatuses(context RequestContext) ([]v2.OrderStatus, error) {
	response, err := a.api.GetOrderStatuses(context, OrderStatusesRequest{})
	if err != nil {
		return nil, err
	}
	statuses := make([]v2.OrderStatus, 0, len(response.Statuses))
	for _, status := range response.Statuses {
		statuses = append(statuses, orderStatusV2FromLegacy(status))
	}
	return statuses, nil
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

func orderStatusV2FromLegacy(status OrderStatus) v2.OrderStatus {
	return v2.OrderStatus{
		StatusId:              uint(status.StatusId),
		Code:                  status.Code,
		Name:                  status.Name,
		FormInvoiceOnChange:   status.FormInvoiceOnChange,
		IsInvoiceDownloadable: status.IsInvoiceDownloadable,
		ShowAsPaid:            status.ShowAsPaid,
		IsWaitingForPayment:   status.IsWaitingForPayment,
		IsCancelled:           status.IsCancelled,
		IsOrderCompleted:      status.IsOrderCompleted,
		IsActive:              status.IsActive,
	}
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

type rpcReferenceDataV2API struct {
	api *apiRPCClient
}

func (a rpcReferenceDataV2API) GetDocumentOperations(context RequestContext, request v2.DocumentOperationsRequest) (v2.DocumentOperationsResponse, error) {
	return a.api.GetDocumentOperations(context, request)
}

func (a rpcReferenceDataV2API) GetDivisions(context RequestContext, request v2.DivisionsRequest) ([]v2.Division, error) {
	return a.api.GetDivisions(context, request)
}

func (a rpcReferenceDataV2API) GetCountryByName(context RequestContext, name string) (v2.Country, error) {
	return a.api.GetCountryByName(context, name)
}

func (a rpcReferenceDataV2API) GetAllCountries(context RequestContext) ([]v2.Country, error) {
	return a.api.GetAllCountries(context)
}

type rpcCommerceChannelsV2API struct {
	api *apiRPCClient
}

func (a rpcCommerceChannelsV2API) GetCommerceChannelByUUID(context RequestContext, channelUUID string) (v2.CommerceChannel, error) {
	return a.api.GetCommerceChannelByUUID(context, channelUUID)
}

func (a rpcCommerceChannelsV2API) CreateCommerceChannel(context RequestContext, channel v2.CommerceChannel) (uint, error) {
	return a.api.CreateCommerceChannel(context, channel)
}

type rpcSyncV2API struct {
	api *apiRPCClient
}

func (a rpcSyncV2API) CreateIntegrationSyncRecord(context RequestContext, sync v2.IntegrationSyncRecord) (uint, error) {
	return a.api.CreateIntegrationSyncRecord(context, sync)
}

func (a rpcSyncV2API) DeleteIntegrationSyncRecordById(context RequestContext, syncRecordId uint) error {
	return a.api.DeleteIntegrationSyncRecordById(context, syncRecordId)
}

func (a rpcSyncV2API) GetIntegrationSyncRecords(context RequestContext, request v2.IntegrationSyncRecordsRequest) ([]v2.IntegrationSyncRecord, error) {
	return a.api.GetIntegrationSyncRecords(context, request)
}

func (a rpcSyncV2API) UpsertPluginSyncState(context RequestContext, state v2.PluginSyncState) error {
	return a.api.UpsertPluginSyncState(context, state)
}

func (a rpcSyncV2API) ListPluginSyncStates(context RequestContext, request v2.PluginSyncStateRequest) (v2.PluginSyncStateResponse, error) {
	return a.api.ListPluginSyncStates(context, request)
}

func (a rpcSyncV2API) CountPluginSyncStates(context RequestContext, request v2.PluginSyncStateRequest) (int64, error) {
	return a.api.CountPluginSyncStates(context, request)
}

type rpcCustomFieldsV2API struct {
	api *apiRPCClient
}

func (a rpcCustomFieldsV2API) GetCustomFields(context RequestContext, request v2.CustomFieldsRequest) (v2.CustomFieldsResponse, error) {
	return a.api.GetCustomFields(context, request)
}

func (a rpcCustomFieldsV2API) CreateCustomField(context RequestContext, request v2.CustomField) (uint, error) {
	return a.api.CreateCustomField(context, request)
}

type rpcAttachmentsV2API struct {
	api *apiRPCClient
}

func (a rpcAttachmentsV2API) CreateAttachment(context RequestContext, request v2.AttachmentCreateRequest) (v2.AttachmentCreateResponse, error) {
	return a.api.UploadFile(context, request)
}

func (a rpcAttachmentsV2API) DeleteAttachment(context RequestContext, attachmentId uint) error {
	return a.api.DeleteFile(context, attachmentId)
}

type rpcEmailV2API struct {
	api *apiRPCClient
}

func (a rpcEmailV2API) SendEmail(context RequestContext, request v2.EmailRequest) error {
	return a.api.SendEmail(context, request)
}

func (a rpcProductsV2API) GetProductFeatureValues(context RequestContext, productId uint) ([]v2.ProductFeatureValue, error) {
	return a.api.GetProductFeatureValuesV2(context, productId)
}
