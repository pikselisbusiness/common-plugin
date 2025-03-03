package shared

// directly.
type API interface {
	RegisterCronJob(schedule string)
	GetConfigVariable(name string) (string, error)
	GetUserInfoForUserId(userId uint) UserInfo
	GetInvoices(request InvoicesRequest) (InvoicesListResponse, error)
	GetInvoiceProducts(context RequestContext, invoiceId uint) ([]InvoiceProduct, error)
	GetInvoiceReferences(context RequestContext, invoiceId uint) ([]InvoiceReference, error)
	GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error)

	GetCompanyById(context RequestContext, companyId uint) (Company, error)
	GetCompaniesMap(context RequestContext, companyIds []uint) (map[uint]Company, error)
	GetCompanyByCode(context RequestContext, companyCode string) (Company, error)
	GetCompanyByVatCode(context RequestContext, companyVatCode string) (Company, error)
	CreateCompany(context RequestContext, company Company) (uint, error)

	GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error)
	GetProductsMap(context RequestContext, productsIds []uint) (map[uint]Product, error)
	GetProductById(context RequestContext, productId uint) (Product, error)
	GetProductByAnyField(context RequestContext, fieldName string, fieldValue any) (Product, error)
	GetProductCategories(context RequestContext, request ProductCategoriesRequest) (ProductCategoriesResponse, error)
	GetProductStocks(context RequestContext, request ProductStocksRequest) (ProductStocksResponse, error)
	GetAvailableProductQuantityByWarehouses(context RequestContext, productId uint, warehouses []string) (float64, error)
	CreateProduct(context RequestContext, request ProductCreateEditRequest) (uint, error)

	GetOrderById(context RequestContext, orderId uint) (Order, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
	CreateOrder(context RequestContext, request OrderCreateRequest) (uint, error, OrderErrorResponse)

	GetCountryByName(context RequestContext, name string) (Country, error)
	GetAllCountries(context RequestContext) ([]Country, error)

	CreateInvoice(context RequestContext, request InvoiceCreateUpdateRequest) (uint, error, InvoiceErrorResponse)
	GetInvoiceExistsByDocument(context RequestContext, document string) (bool, error)
	PatchUpdateInvoice(context RequestContext, invoiceId uint, request map[string]interface{}) error
	CreateInvoiceReference(context RequestContext, request InvoiceReference) error

	CreateIntegrationSyncRecord(context RequestContext, sync IntegrationSyncRecord) (uint, error)
	DeleteIntegrationSyncRecordById(context RequestContext, syncRecordId uint) error
	GetIntegrationSyncRecords(context RequestContext, request IntegrationSyncRecordsRequest) ([]IntegrationSyncRecord, error)

	// GetProductCategoriesChainByLastChildId(context RequestContext, categoryId uint) ([]ProductCategoriesShort, error)
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
