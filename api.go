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
	GetProducts(context RequestContext, request ProductsRequest) (ProductsResponse, error)
	GetProductsMap(context RequestContext, productsIds []uint) (map[uint]Product, error)
	GetProductById(context RequestContext, productId uint) (Product, error)
	GetOrderById(context RequestContext, orderId uint) (Order, error)
	GetOrders(context RequestContext, request OrdersRequest) (OrdersResponse, error)
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
