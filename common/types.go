package common

type RequestContext struct {
	UserId         uint
	UserGroupId    uint
	OrganizationId uint
	HostName       string
	UserName       string
	UserFirstName  string
	UserLastName   string
	UserEmail      string
	UserIp         string
	AccessedPath   string
	AccessedMethod string
	UserLanguage   string
	IsSystemUser   bool
	IdempotencyKey *string
	RoleIDs        []uint
}

type UserInfo struct {
	UserId       uint   `json:"userId"`
	UserImageUrl string `json:"userImageUrl"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type InvoiceListKind string

const (
	InvoiceListKindSale          InvoiceListKind = "sale"
	InvoiceListKindPurchase      InvoiceListKind = "purchase"
	InvoiceListKindWarehouseMove InvoiceListKind = "warehouseMove"
	InvoiceListKindAll           InvoiceListKind = "all"
)

type ItemFilter struct {
	Operator string `json:"operator" query:"operator"`
	Value    string `json:"value" query:"value"`
	Field    string `json:"field" query:"field"`
}

type PriceRange struct {
	PriceFrom        float64 `json:"priceFrom"`
	PriceTo          float64 `json:"priceTo"`
	PriceFromWithVat float64 `json:"priceFromWithVat"`
	PriceToWithVat   float64 `json:"priceToWithVat"`
}

type InvoiceErrorResponse struct {
	ErrorAtProduct  bool
	ErrorProductKey int
	ErrorField      string
	ErrorMessage    string
}

type OrderErrorResponse struct {
	ErrorAtProduct  bool
	ErrorProductKey int
	ErrorField      string
	ErrorMessage    string
}
