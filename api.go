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
	CreateProduct(context RequestContext, request ProductCreateEditRequest) (uint, error)

	GetOrderById(context RequestContext, orderId uint) (Order, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
	CreateOrder(context RequestContext, request OrderCreateRequest) (uint, error, OrderErrorResponse)

	CreateInvoice(context RequestContext, request InvoiceCreateUpdateRequest) (uint, error, InvoiceErrorResponse)
	GetInvoiceExistsByDocument(context RequestContext, document string) (bool, error)

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
