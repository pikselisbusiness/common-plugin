package shared

// directly.
type API interface {
	RegisterCronJob(schedule string)
	RegisterCronJobWithTag(tag string, schedule string)
	GetConfigVariable(name string) (string, error)
	GetUserInfoForUserId(userId uint) UserInfo
	GetInvoices(request InvoicesRequest) (InvoicesListResponse, error)
	GetInvoiceProducts(context RequestContext, invoiceId uint) ([]InvoiceProduct, error)
	GetInvoiceReferences(context RequestContext, invoiceId uint) ([]InvoiceReference, error)
	GetDocumentOperations(context RequestContext, request DocumentOperationsRequest) (DocumentOperationsResponse, error)
	GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error)

	GetCompanyById(context RequestContext, companyId uint) (Company, error)
	GetCompanies(context RequestContext, request CompaniesRequest) (CompaniesResponse, error)
	GetCompaniesMap(context RequestContext, companyIds []uint) (map[uint]Company, error)
	GetCompanyByCode(context RequestContext, companyCode string) (Company, error)
	GetCompanyByVatCode(context RequestContext, companyVatCode string) (Company, error)
	CreateCompany(context RequestContext, company Company) (uint, error)
	GetCommerceChannelByUUID(context RequestContext, channelUUID string) (CommerceChannel, error)
	CreateCommerceChannel(context RequestContext, channel CommerceChannel) (uint, error)

	GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error)
	GetProductsMap(context RequestContext, productsIds []uint) (map[uint]Product, error)
	GetProductById(context RequestContext, productId uint) (Product, error)
	GetProductByAnyField(context RequestContext, fieldName string, fieldValue any) (Product, error)
	GetProductCategories(context RequestContext, request ProductCategoriesRequest) (ProductCategoriesResponse, error)
	GetProductStocks(context RequestContext, request ProductStocksRequest) (ProductStocksResponse, error)
	GetAvailableProductQuantityByWarehouses(context RequestContext, productId uint, warehouses []string) (float64, error)
	CreateProduct(context RequestContext, request ProductCreateEditRequest) (uint, error)

	// Product features. UpsertProductFeature idempotently creates or updates a feature
	// (attribute definition) by name and returns its id. ReplaceProductFeatures matches
	// products by SKU and replaces each matched product's assignments per feature.
	//
	// TODO(v0.2.0 vocabulary): these bulk methods grew plugin-outward and their naming
	// is inconsistent with the core import engine (dto.ProductImportRequest). Target
	// for the next breaking bump: adopt the core dialect — `ProductIdentifierColumn`
	// (not MatchColumn), `Products` (not Items), item field `Identifier` (not Sku),
	// result `UnmatchedIdentifiers` (not UnmatchedSkus) — and extract the shared
	// host-side identifier-resolution helper. Method verbs stay separate.
	UpsertProductFeature(context RequestContext, request UpsertFeatureRequest) (UpsertFeatureResponse, error)
	ReplaceProductFeatures(context RequestContext, request ReplaceProductFeaturesRequest) (ReplaceProductFeaturesResult, error)
	PruneProductFeatures(context RequestContext, request PruneProductFeaturesRequest) (PruneProductFeaturesResult, error)
	// ImportProducts passes rows to the core product import engine (identifier
	// matching, modes, validation, customFields.<slug> columns).
	ImportProducts(context RequestContext, request ImportProductsRequest) (ImportProductsResponse, error)

	GetOrderById(context RequestContext, orderId uint) (Order, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
	GetOrderStatuses(context RequestContext, request OrderStatusesRequest) (OrderStatusesResponse, error)
	CreateOrder(context RequestContext, request OrderCreateRequest) (uint, error, OrderErrorResponse)
	PatchUpdateOrder(context RequestContext, orderId uint, request map[string]interface{}) error

	GetCountryByName(context RequestContext, name string) (Country, error)
	GetAllCountries(context RequestContext) ([]Country, error)

	CreateInvoice(context RequestContext, request InvoiceCreateUpdateRequest) (uint, error, InvoiceErrorResponse)
	GetInvoiceExistsByDocument(context RequestContext, document string) (bool, error)
	GetInvoicePayments(context RequestContext, request InvoicePaymentsRequest) (InvoicePaymentsResponse, error)
	GetJournalEntries(context RequestContext, request JournalEntriesRequest) (JournalEntriesResponse, error)
	GetJournalEntryById(context RequestContext, journalEntryId uint) (JournalEntry, error)
	PatchUpdateInvoice(context RequestContext, invoiceId uint, request map[string]interface{}) error
	CreateInvoiceReference(context RequestContext, request InvoiceReference) error

	CreateIntegrationSyncRecord(context RequestContext, sync IntegrationSyncRecord) (uint, error)
	DeleteIntegrationSyncRecordById(context RequestContext, syncRecordId uint) error
	GetIntegrationSyncRecords(context RequestContext, request IntegrationSyncRecordsRequest) ([]IntegrationSyncRecord, error)
	UpsertPluginSyncState(context RequestContext, state PluginSyncState) error
	ListPluginSyncStates(context RequestContext, request PluginSyncStateRequest) (PluginSyncStateResponse, error)
	CountPluginSyncStates(context RequestContext, request PluginSyncStateRequest) (int64, error)

	GetPosDiscountCards(context RequestContext, request PosDiscountCardsRequest) (PosDiscountCardsResponse, error)

	GetCustomFields(context RequestContext, request CustomFieldsRequest) (CustomFieldsResponse, error)
	CreateCustomField(context RequestContext, request CustomField) (uint, error)

	SendEmail(context RequestContext, request EmailRequest) error

	UploadFile(context RequestContext, request FileUploadRequest) (FileUploadResponse, error)
	DeleteFile(context RequestContext, uploadId uint) error

	// GetProductCategoriesChainByLastChildId(context RequestContext, categoryId uint) ([]ProductCategoriesShort, error)
}

// APIStopper lets long-running work ask the host whether it has begun shutting
// down, so it can stop at a chunk boundary with its progress saved instead of
// being severed when the host's grace period expires.
//
// Deliberately NOT part of API: plugins hand-roll full mocks of that interface
// in their tests, so every method added to it is a breaking change for every
// plugin on its next SDK bump. This is the same optional-interface upgrade used
// for CommonV2/CommonV3 — assert for it and degrade if it is absent:
//
//	if stopper, ok := api.(shared.APIStopper); ok {
//		if stop, err := stopper.ShouldStop(); err == nil && stop {
//			// save progress and return
//		}
//	}
//
// Polled rather than pushed because the host↔plugin link is net/rpc: a context
// cannot cross it, and a call already in flight has no deadline. One local
// round-trip per chunk is cheap next to the remote call the chunk itself makes.
//
// An error means no answer — typically a host older than this method — and must
// be read as "carry on". Never stop working because the question failed.
type APIStopper interface {
	ShouldStop() (bool, error)
}

// Client is a streamlined wrapper over the pikselis-business plugin API.
type Client struct {
	Api API
	Db  DB
}

// NewClient creates a new instance of Client.
//
// This client must only be created once per plugin to
// prevent reacquiring of resources.
func NewClient(api API, db DB) *Client {
	return &Client{
		Api: api,
		Db:  db,
	}
}
