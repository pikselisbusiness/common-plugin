package shared

import (
	"encoding/json"
	"time"

	common "github.com/pikselisbusiness/common-plugin/common"
	v2 "github.com/pikselisbusiness/common-plugin/v2"
)

type UserInfo = common.UserInfo
type Company struct {
	CompanyId      int32  `json:"companyId"`
	CompanyName    string `json:"companyName"`
	CompanyCode    string `json:"companyCode"`
	CompanyVatCode string `json:"companyVatCode"`
	IsPerson       bool   `json:"isPerson"`
	Address        string `json:"address"`
	PostCode       string `json:"postCode"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PhoneNumber    string `json:"phoneNumber"`

	WarnedDate        time.Time `json:"warnedDatetime"`
	DeferSaleDays     float64   `json:"deferSaleDays"`
	DeferPurchaseDays float64   `json:"deferPurchaseDays"`
	RequireDays       float64   `json:"requireDays"`
	ExecuteDays       float64   `json:"executeDays"`
	State             string    `json:"state"`
	Representative    string    `json:"representative"`
	Fax               string    `json:"fax"`
	URL               string    `json:"url"`
	Email             string    `json:"email"`
	// BankPav                   string    `json:"bank_pav"`
	// BankKod                   string    `json:"bank_kod"`
	// AtsSask                   string    `json:"ats_sask"`
	IsSupplier                bool            `json:"isSupplier"`
	IsBuyer                   bool            `json:"isBuyer"`
	IsBranch                  bool            `json:"isBranch"`
	IsManufacturer            bool            `json:"isManufacturer"`
	Contract                  string          `json:"contract"`
	ContractFromDatetime      time.Time       `json:"contractFromDatetime"`
	ContractToDatetime        time.Time       `json:"contractToDatetime"`
	MaxDebt                   float64         `json:"maxDebt"`
	MaxCurrency               string          `json:"maxCurrency"`
	PersonContactId           int32           `json:"personContactId"`
	DeferSaleMonths           int32           `json:"deferSaleMonths"`
	DeferSaleType             string          `json:"deferSaleType"`
	DeferPurchaseMonths       int32           `json:"deferPurchaseMonths"`
	DeferPurchaseType         string          `json:"deferPurchaseType"`
	WantsInvoiceByRegularMail int32           `json:"wantsInvoiceByRegularMail"`
	RivileCustomerCode        string          `json:"rivileCustomerCode"`
	PriceField                string          `json:"priceField"`
	CustomFields              json.RawMessage `json:"customFields"`

	CountVat         bool      `json:"countVat"`
	CustomerLanguage string    ` json:"customerLanguage"`
	InsertDatetime   time.Time `json:"insertDatetime"`
	InsertUserId     uint      `json:"insertUserId"`
	UpdateDatetime   time.Time `json:"updateDatetime"`
	UpdateUserId     uint      `json:"updateUserId"`

	DebtsMap       map[string]float64 `json:"debts"`
	MissedDebtsMap map[string]float64 `json:"missedDebts"`
	AdvanceSums    CompanyAdvanceSums `json:"advanceSums"`
}
type CompanyAdvanceSums struct {
	SupplierAdvances map[string]float64 `json:"supplierAdvances"`
	BuyerAdvances    map[string]float64 `json:"buyerAdvances"`
}

type CompaniesRequest struct {
	PerPage     int       `json:"perPage"`
	Page        int       `json:"page"`
	CompanyIds  []uint    `json:"companyIds"`
	SearchQuery string    `json:"searchQuery"`
	DateFrom    time.Time `json:"dateFrom"`
	DateTo      time.Time `json:"dateTo"`
	IsBuyer     bool      `json:"isBuyer"`
	IsSupplier  bool      `json:"isSupplier"`
	OrderBy     string    `json:"orderBy"`
	OrderWay    string    `json:"orderWay"`
}
type CompaniesResponse struct {
	Success        bool      `json:"success"`
	Companies      []Company `json:"companies"`
	CompaniesCount int64     `json:"companiesCount"`
}
type InvoicesRequest struct {
	UnlimitPerPage          bool
	PerPage                 int
	Page                    int
	DateFrom                time.Time
	DateTo                  time.Time
	SearchQuery             string
	SelectOnlySales         bool
	SelectOnlyPurchases     bool
	SelectOnlyWarehouseMove bool
	SelectOnlyDebtInvoices  bool
	FromWarehouses          []string
	ToWarehouses            []string
	Operations              []string
	InvoiceId               []uint
	CompanyIds              []uint
	ExcludeCompanyIds       []uint
	OrderWay                string
	Statuses                []string
}

type InvoiceListKind = common.InvoiceListKind

const (
	InvoiceListKindSale          = common.InvoiceListKindSale
	InvoiceListKindPurchase      = common.InvoiceListKindPurchase
	InvoiceListKindWarehouseMove = common.InvoiceListKindWarehouseMove
	InvoiceListKindAll           = common.InvoiceListKindAll
)

type InvoiceType = v2.InvoiceType

const (
	InvoiceTypeSale     = v2.InvoiceTypeSale
	InvoiceTypePurchase = v2.InvoiceTypePurchase
)

type InvoiceAddress = v2.InvoiceAddress
type InvoiceInstallment = v2.InvoiceInstallment
type InvoicesListResponse struct {
	Success                 bool               `json:"success"`
	Invoices                []InvoiceFull      `json:"invoices"`
	InvoicesCount           int64              `json:"invoicesCount"`
	TotalDebt               float64            `json:"totalDebt"`
	Currency                string             `json:"currency"`
	TotalDebtsByCurrency    map[string]float64 `json:"totalDebtsByCurrency"`
	Last30DaysInvoicesCount int64              `json:"last30DaysInvoicesCount"`
}

type InvoiceFull struct {
	InvoiceId            int32                `json:"invoiceId"`
	IsPurchase           bool                 `json:"isPurchase"`
	Operation            string               `json:"operation"`
	CompanyName          string               `json:"companyName"`
	CompanyId            int32                `json:"companyId"`
	DocSeries            string               `json:"docSeries"`
	Document             string               `json:"document"`
	Document2            string               `json:"document2"`
	InnerDocument        string               `json:"innerDocument"`
	InvoiceDate          time.Time            `json:"invoiceDate"`
	TotalVat             float64              `json:"totalVat"`
	TotalWoVat           float64              `json:"totalWoVat"`
	TotalWithVat         float64              `json:"totalWithVat"`
	TotalExclTax         float64              `json:"totalExclTax"`
	TotalTax             float64              `json:"totalTax"`
	TotalInclTax         float64              `json:"totalInclTax"`
	TotalCurrency        string               `json:"totalCurrency"`
	TotalDiscountExclTax float64              `json:"totalDiscountExclTax"`
	TotalDiscountInclTax float64              `json:"totalDiscountInclTax"`
	TotalCompensation    float64              `json:"totalCompensation"`
	InvoicePaid          bool                 `json:"invoicePaid"`
	DebtSum              float64              `json:"debtSum"`
	PaidSum              float64              `json:"paidSum"`
	PaidLastDate         time.Time            `json:"paidLastDate"`
	PayUntilDate         time.Time            `json:"payUntilDate"`
	ReceiveDate          time.Time            `json:"receiveDate"`
	TotalCostSum         float64              `json:"totalCostSum"`
	TotalGainSum         float64              `json:"totalGainSum"`
	FromWarehouse        string               `json:"fromWarehouse"`
	ToWarehouse          string               `json:"toWarehouse"`
	Comment              string               `json:"comment"`
	FreezeFinances       bool                 `json:"freezeFinances"`
	Costs                []InvoiceCost        `json:"costs"`
	ExtraSettings        ExtraInvoiceSettings `json:"extraSettings"`
	PosData              InvoicePosData       `json:"posData"`

	InsertDatetime time.Time `json:"insertDatetime"`
	ModifyDatetime time.Time `json:"updateDatetime"`
	InsertUserInfo UserInfo  `json:"insertUserInfo"`
	UpdateUserInfo UserInfo  `json:"updateUserInfo"`

	Company      Company         `json:"company,omitempty"`
	IsVerified   bool            `json:"isVerified"`
	Status       string          `json:"status"`
	CustomFields json.RawMessage `json:"customFields"`
}
type ExtraInvoiceSettings = common.ExtraInvoiceSettings
type InvoiceCost = common.InvoiceCost
type InvoiceProduct = common.InvoiceProduct
type InvoicePosData = common.InvoicePosData
type InvoiceReference = common.InvoiceReference
type InvoicePaymentsRequest = common.InvoicePaymentsRequest
type InvoicePayment = common.InvoicePayment
type InvoicePaymentsResponse = common.InvoicePaymentsResponse
type JournalEntriesRequest = common.JournalEntriesRequest
type JournalEntryLine = common.JournalEntryLine
type JournalEntry = common.JournalEntry
type JournalEntriesResponse = common.JournalEntriesResponse

type DocumentOperation = common.DocumentOperation
type DocumentOperationsRequest = common.DocumentOperationsRequest
type DocumentOperationsResponse = common.DocumentOperationsResponse

/**
---- DIVISIONS
*/

type DivisionsRequest = common.DivisionsRequest
type Division = common.Division
type DivisionWarehouse = common.DivisionWarehouse

type Product struct {
	ProductId     uint    `json:"productId"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"` //@TODO - on eshop - change this with priceWithVat
	PriceWithVat  float64 `json:"priceWithVat"`
	Currency      string  `json:"currency"`
	Price1        float64 `json:"price1"`
	Price1WithVat float64 `json:"price1WithVat"` // @TODO skip fields on create
	Currency1     string  `json:"currency1"`
	Price2        float64 `json:"price2"`
	Price2WithVat float64 `json:"price2WithVat"` // @TODO skip fields on create
	Currency2     string  `json:"currency2"`
	Price3        float64 `json:"price3"`
	Price3WithVat float64 `json:"price3WithVat"`
	Currency3     string  `json:"currency3"`

	Price4        float64 `json:"price4"`
	Price4WithVat float64 `json:"price4WithVat"`
	Currency4     string  `json:"currency4"`

	Price5        float64 `json:"price5"`
	Price5WithVat float64 `json:"price5WithVat"`
	Currency5     string  `json:"currency5"`

	Price6        float64 `json:"price6"`
	Price6WithVat float64 `json:"price6WithVat"`
	Currency6     string  `json:"currency6"`

	Price7        float64 `json:"price7"`
	Price7WithVat float64 `json:"price7WithVat"`
	Currency7     string  `json:"currency7"`

	Price8        float64 `json:"price8"`
	Price8WithVat float64 `json:"price8WithVat"`
	Currency8     string  `json:"currency8"`

	Price9        float64 `json:"price9"`
	Price9WithVat float64 `json:"price9WithVat"`
	Currency9     string  `json:"currency9"`

	Price10        float64 `json:"price10"`
	Price10WithVat float64 `json:"price10WithVat"`
	Currency10     string  `json:"currency10"`

	MeasurementUnit  string  `json:"measurementUnit"`
	Articule         string  `json:"articule"`
	NameEn           string  `json:"nameEn"`
	Description      string  `json:"description"`
	Countable        bool    `json:"countable"`
	OrderedQuantity  float64 `json:"orderedQuantity"`
	SupplierCode1    string  `json:"supplierCode1"`
	SupplierCode2    string  `json:"supplierCode2"`
	SupplierCode3    string  `json:"supplierCode3"`
	Brand            string  `json:"brand"`
	BrandEn          string  `json:"brandEn"`
	BrandRu          string  `json:"brandRu"`
	BrandPl          string  `json:"brandPl"`
	IsDiscounted     bool    `json:"isDiscounted"`
	IsInPriceList    bool    `json:"isInPriceList"`
	IsOld            bool    `json:"isOld"`
	Importer         string  `json:"imported"`
	IsNew            bool    `json:"isNew"`
	WarrantyInMonths int32   `json:"warrantyInMonths"`
	IsRecommended    bool    `json:"isRecommended"`
	IsRemoved        bool    `json:"isRemoved"`
	Comment          string  `json:"comment"`
	Link             string  `json:"link"`
	LinkEn           string  `json:"linkEn"`
	LinkRu           string  `json:"linkRu"`
	LinkPl           string  `json:"linkPl"`
	NameRu           string  `json:"nameRu"`
	DescriptionEn    string  `json:"descriptionEn"`
	DescriptionRu    string  `json:"descriptionRu"`

	NamePl        string `json:"namePl"`
	DescriptionPl string `json:"nescriptionPl"`

	PriceFor float64 `json:"priceFor"`
	Package1 int32   `json:"package1"`
	Package2 int32   `json:"package2"`
	Package3 int32   `json:"package3"`

	DiscountValue        float64 `json:"discountValue"`
	DiscountType         int32   `json:"discountType"`
	DiscountPercent      float64 `json:"discountPercent"`      // not in sql
	DiscountedPrice      float64 `json:"discountedPrice"`      // not in sql
	DiscountedPriceWoVat float64 `json:"discountedPriceWoVat"` // not in sql
	Orderby              uint    `json:"orderby"`
	MainCategoryId       int32   `json:"mainCategoryId"`

	Quantity     int64              `json:"quantity"`
	ImageMedium  string             `json:"image"`
	ImageOrigin  string             `json:"imageOrigin"`
	ImageSmall   string             `json:"imageSmall"`
	Images       []string           `json:"images"`
	ImagesOrigin []string           `json:"imagesOrigin"`
	ImagesSmall  []string           `json:"imagesSmall"`
	MediaItems   []ProductMediaFile `json:"mediaItems"`

	ProductWeight   float64 `json:"productWeight"`
	ProductSize     string  `json:"productSize"`
	PrimaryPrice    float64 `json:"primaryPrice"`
	CodeForGrouping string  `json:"codeForGrouping"`
	ProductSeason   string  `json:"productSeason"`
	Volume          float64 `json:"volume"`
	Barcode         string  `json:"barcode"`
	VatPercentage   float64 `json:"vatPercentage"`
	// Content   string  `json:"productSeason"`
	OriginCountry string `json:"originCountry"`

	SelectedCategoryIds []int32           `json:"selectedCategoryIds"`
	MetaValues          map[string]string `json:"metaValues"`

	AccountingNumbers ProductAccountingNumbers `json:"accountingNumbers"`

	SyncPrestashop           bool    `json:"syncPrestashop"`
	SyncWoocommerce          bool    `json:"syncWoocommerce"`
	ProductLocation          string  `json:"productLocation"`
	MainForEshop             bool    `json:"mainForEshop"`
	OnlyInPhysicalStore      bool    `json:"onlyInPhysicalStore"`
	NameInEshop              string  `json:"nameInEshop"`
	ProductLength            float32 `json:"productLength"`
	ProductWidth             float32 `json:"productWidth"`
	ProductHeight            float32 `json:"productHeight"`
	OrgEanCode               string  `json:"orgEanCode"`
	ShowInPiguXML            bool    `json:"showInPiguXml"`
	NameInPigu               string  `json:"nameInPigu"`
	ShowInSale               bool    `json:"showInSale"`
	IsPhysicalOnlyProduct    bool    `json:"isPhysicalOnlyProduct"`
	AvailableForOrderInEshop bool    `json:"availableForOrderInEshop"`
	IsElectronicLicense      bool    `json:"isElectronicLicense"`

	CategoryNames      []string                  `json:"categoryNames"`
	CategoryNamesList  [][]string                `json:"categoryNamesList"`
	Packages           []ProductPackage          `json:"packages"`
	AttributeVariation ProductAttributeVariation `json:"attributeVariation"`

	// InsertDatetime   time.Time `json:"insertDatetime"`
	// ModifyDatetime   time.Time `json:"modifyDatetime"`
	CreatedUserId   uint                `json:"-"` // inner only
	UpdatedUserId   uint                `json:"-"` // inner only
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
	CreatedUserInfo UserInfo            `json:"createdUserInfo"`
	UpdatedUserInfo UserInfo            `json:"updatedUserInfo"`
	External        ProductExternalInfo `json:"external"`

	PhotosUpdated bool
	CustomFields  json.RawMessage `json:"customFields"`
}

type ProductMediaFile = v2.ProductMediaFile
type ProductMediaContent = v2.ProductMediaContent

type ProductPackage struct {
	PackageId               uint    `json:"packageId"`
	Quantity                string  `json:"quantity"`
	MeasurementUnit         string  `json:"measurementUnit"`
	Price1                  float64 `json:"price1"`
	Price1WithVat           float64 `json:"price1WithVat"`
	DiscountedPrice1        float64 `json:"discountedPrice1"`
	DiscountedPrice1WithVat float64 `json:"discountedPrice1WithVat"`
	Currency1               string  `json:"currency1"`
	Price2                  float64 `json:"price2"`
	Price2WithVat           float64 `json:"price2WithVat"`
	Currency2               string  `json:"currency2"`
	Discount                float64 `json:"discount"`
}

type ProductExternalInfo struct {
	PrestaId            uint `json:"prestaId"`
	PrestaCombinationId uint `json:"prestaCombinationId"`
}
type ProductAccountingNumbers = v2.ProductAccountingNumbers

type ProductAttributeVariation struct {
	Attributes    []ProductAttribute `json:"attributes"`
	IsMainProduct bool               `json:"isMainProduct"`
}
type ProductAttribute struct {
	AttributeValueId   uint   `json:"attributeValueId" `
	AttributeName      string `json:"attributeName"`
	AttributeColorCode string `json:"attributeColorCode"`
	AttributeType      string `json:"attributeType"`
	GroupId            uint   `json:"groupId"`
	PrestaId           uint   `json:"prestaId"`
}

type ProductCategoriesShort = common.ProductCategoriesShort

type OrderShippingAddress struct {
	ShippingFirstName   string `json:"shippingFirstName"`
	ShippingLastName    string `json:"shippingLastName"`
	ShippingCompanyName string `json:"shippingCompanyName"`
	ShippingAddress     string `json:"shippingAddress"`
	ShippingCity        string `json:"shippingCity"`
	ShippingCountry     string `json:"shippingCountry"`
	ShippingPostCode    string `json:"shippingPostCode"`
	ShippingHomePhone   string `json:"shippingHomePhone"`
	ShippingMobilePhone string `json:"shippingMobilePhone"`
}
type OrderPickupLocation struct {
	LocationId          int32  `json:"locationId"`
	LocationName        string `json:"locationName"`
	LocationCity        string `json:"locationCity"`
	LocationPostCode    string `json:"locationPostCode"`
	LocationCompanyCode string `json:"locationCompanyCode"`
}

type OrderHeadingsWithLines struct {
	Title        string      `json:"title"`
	TotalWoVat   float64     `json:"totalWoVat"`
	TotalWithVat float64     `json:"totalWithVat"`
	TotalVat     float64     `json:"totalVat"`
	Currency     string      `json:"currency"`
	Lines        []OrderLine `json:"lines"`
}
type OrderLine struct {
	ID             int32  `json:"id"`
	LineType       uint   `json:"lineType"`
	LineTypeString string `json:"lineTypeString"`
	Title          string `json:"title"`
	Description    string `json:"description"`

	ProductId       uint                      `json:"productId"`
	ProductName     string                    `json:"productName"`
	ProductInfo     Product                   `json:"productInfo"`
	Quantity        float64                   `json:"quantity"`
	OriginalPrice   float64                   `json:"originalPrice"`
	CostPrice       float64                   `json:"costPrice"`
	DealerPrice     float64                   `json:"dealerPrice"`
	BuyerPrice      float64                   `json:"price"`
	PriceWithVat    float64                   `json:"priceWithVat"`
	Sum             float64                   `json:"sum"`
	SumWithVat      float64                   `json:"sumWithVat"`
	DiscountAmount  float64                   `json:"discountAmount"`
	DiscountPercent float64                   `json:"discountPercent"`
	DiscountSum     float64                   `json:"discountSum"`
	Articule        string                    `json:"articule"`
	CostCurrency    string                    `json:"costCurrency"`
	DealerCurrency  string                    `json:"dealerCurrency"`
	BuyerCurrency   string                    `json:"currency"`
	MeasurementUnit string                    `json:"measurementUnit"`
	VatPercentage   float64                   `json:"vatPercentage"`
	VatClass        string                    `json:"vatClass"`
	Package         string                    `json:"package"`
	IsTransport     bool                      `json:"isTransport"`
	Charges         []LineChargeCreateRequest `json:"charges,omitempty"`
}
type OrderStatus struct {
	StatusId              int32  `json:"statusId"`
	Selector              int32  `json:"selector"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	FormInvoiceOnChange   bool   `json:"formInvoiceOnChange"`
	IsInvoiceDownloadable bool   `json:"isInvoiceDownloadable"`
	SendEmail             int    `json:"sendEmail"`
	ShowAsPaid            bool   `json:"showAsPaid"`
	IsWaitingForPayment   bool   `json:"isWaitingForPayment"`
	DoNotPromptEmailText  bool   `json:"doNotPromptEmailText"`
	IsOrderCompleted      bool   `json:"isOrderCompleted"`
	DocSeries             string `json:"docSeries"`
	IsCancelled           bool   `json:"isCancelled"`
	IsActive              bool   `json:"isActive"`
}
type OrderCompanyInfo struct {
	CompanyName    string `json:"companyName"`
	CompanyId      uint   `json:"companyId"`
	CompanyCode    string `json:"companyCode"`
	CompanyVatCode string `json:"companyVatCode"`
	IsPerson       bool   `json:"isPerson"`
	Address        string `json:"address"`
	PostCode       string `json:"postCode"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PhoneNumber    string `json:"phoneNumber"`
}
type OrderCompanyAddressInfo struct {
	ContactId    uint   `json:"contactId"`
	Title        string `json:"title"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	MobilePhone  string `json:"mobilePhone"`
	Email        string `json:"email"`
	SendContract bool   `json:"sendContract"`
	SendInvoice  bool   `json:"sendInvoice"`
	SendEmail    bool   `json:"sendEmail"`
	ContactType  string `json:"contactType"`
	Address      string `json:"address"`
	AddressLine2 string `json:"addressLine2"`
	State        string `json:"state"`
	Region       string `json:"region"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PostCode     string `json:"postCode"`
}
type OrderInvoiceInfo struct {
	InvoiceIsFormed bool      `json:"invoiceIsFormed"`
	InvoiceId       uint      `json:"invoiceId"`
	DocSeries       string    `json:"docSeries"`
	Document        string    `json:"document"`
	InvoiceDate     time.Time `json:"invoiceDate"`
	InsertUserInfo  UserInfo  `json:"insertUserInfo"`
	InsertDatetime  time.Time `json:"insertDatetime"`
}
type Order struct {
	OrderId          int32                `json:"orderId"`
	Date             time.Time            `json:"date"`
	InsertDatetime   time.Time            `json:"insertDatetime"`
	Type             string               `json:"type"`
	CompanyId        int32                `json:"companyId"`
	CompanyName      string               `json:"companyName"`
	OrderStatus      string               `json:"orderStatus"`
	OrderStatusColor string               `json:"orderStatusColor"`
	Title            string               `json:"title"`
	Period           int32                `json:"period"`
	InsertUserid     int32                `json:"insertUserId"`
	Document         string               `json:"document"`
	TotalVat         float64              `json:"totalVat"`
	TotalWoVat       float64              `json:"totalWoVat"`
	TotalWithVat     float64              `json:"totalWithVat"`
	TotalCurrency    string               `json:"totalCurrency"`
	IsDeleted        bool                 `json:"isDeleted"`
	StatusLetter     string               `json:"statusLetter"`
	ReportType       string               `json:"reportType"` // PDF type
	Comments         string               `json:"comments"`
	PdfBlankId       int32                `json:"pdfBlankId"`
	ProductQuantity  float64              `json:"productQuantity"`
	ShippingAddress  OrderShippingAddress `json:"shippingAddress"`

	// Additional extended fields
	OtherDoc               string                   `json:"otherDoc"`
	Email                  string                   `json:"email"`
	ExternalUserId         uint                     `json:"externalUserId"`
	TotalDiscount          float64                  `json:"totalDiscount"`
	TotalDiscountWithVat   float64                  `json:"totalDiscountWithVat"`
	DiscountPercent        float64                  `json:"discountPercent"`
	TotalShipping          float64                  `json:"totalShipping"`
	TotalShippingWithVat   float64                  `json:"totalShippingWithVat"`
	CompanyInfo            OrderCompanyInfo         `json:"companyInfo"`
	OrderStatusInfo        OrderStatus              `json:"orderStatusInfo"` // @TODO rename
	OrderStatuses          []OrderStatus            `json:"orderStatuses"`   // @TODO rename
	Products               []OrderLine              `json:"products"`
	LinesByHeadings        []OrderHeadingsWithLines `json:"linesByHeadings"`
	ExpiryDate             time.Time                `json:"expiryDate"`
	CurrencyId             int32                    `json:"currencyId"`
	PickupLocation         OrderPickupLocation      `json:"pickupLocation"`
	PaymentType            string                   `json:"paymentType"`
	PaymentInnerType       string                   `json:"paymentInnerType"`
	PaymentIsCod           bool                     `json:"paymentIsCod"`
	PaymentReferenceNo     string                   `json:"paymentReferenceNo"`
	ShippingMethodName     string                   `json:"shippingMethodName"`
	ShippingMethodSelector string                   `json:"shippingMethodSelector"`
	TrackingNumber         string                   `json:"trackingNumber"`
	TrackingType           string                   `json:"trackingType"`
	// Below from contact_id_address
	CompanyAddressInfo OrderCompanyAddressInfo `json:"companyAddressInfo"`
	// Additional extended fields
	InvoiceInfo    OrderInvoiceInfo `json:"invoiceInfo"` // if invoice is formed from order - returns invoice info
	InsertUserInfo UserInfo         `json:"insertUserInfo"`
	UpdateUserInfo UserInfo         `json:"updateUserInfo"`
	CustomFields   json.RawMessage  `json:"customFields"`
	Token          string           `json:"token"`
}
type ItemFilter = common.ItemFilter
type OrderCreateRequest struct {
	OrderId                int32                         `json:"orderId"`
	Date                   time.Time                     `json:"date"`
	ExpiryDate             time.Time                     `json:"expiryDate"`
	CompanyId              int32                         `json:"companyId"`
	IsOrder                bool                          `json:"isOrder"`
	IsProforma             bool                          `json:"isProforma"`
	Type                   string                        `json:"type"`
	Comments               string                        `json:"comments"`
	CurrencyId             int32                         `json:"currencyId"`
	Currency               string                        `json:"currency"`
	Period                 int32                         `json:"period"`
	StatusLetter           string                        `json:"statusLetter"`
	Title                  string                        `json:"title"`
	ExternalUserId         uint                          `json:"externalUserId"`
	Email                  string                        `json:"email"`
	ShippingAddress        OrderShippingAddress          `json:"shippingAddress"`
	PickupLocation         OrderPickupLocation           `json:"pickupLocation"`
	PaymentId              uint                          `json:"paymentId"`
	PaymentType            string                        `json:"paymentType"`
	PaymentIsCod           bool                          `json:"paymentIsCod"`
	PaymentInnerType       string                        `json:"paymentInnerType"`
	ShippingId             uint                          `json:"shippingId"`
	ShippingMethodSelector string                        `json:"shippingMethodSelector"`
	ShippingMethodName     string                        `json:"shippingMethodName"`
	TrackingNumber         string                        `json:"trackingNumber"`
	TrackingType           string                        `json:"trackingType"`
	ReportType             string                        `json:"reportType"` // PDF type
	PdfBlankId             int32                         `json:"pdfBlankId"`
	Products               []OrderLine                   `json:"products"`
	Token                  string                        `json:"token"`
	PaymentReferenceNo     string                        `json:"paymentReferenceNo"`
	CustomFields           json.RawMessage               `json:"customFields"`
	StatusId               uint                          `json:"statusId"`
	ChannelID              uint                          `json:"channelId"`
	ExternalOrderID        string                        `json:"externalOrderId"`
	DocumentCharges        []DocumentChargeCreateRequest `json:"documentCharges,omitempty"`
}
type OrdersRequest struct {
	PerPage                      int       `json:"perPage"`
	Page                         int       `json:"page"`
	DateFrom                     time.Time `json:"dateFrom"`
	DateTo                       time.Time `json:"dateTo"`
	SearchQuery                  string    `json:"searchQuery"`
	SelectOnlyOffers             bool      `json:"selectOnlyOffers"`
	SelectOnlyOrders             bool      `json:"selectOnlyOrders"`
	SelectOnlyProformas          bool      `json:"selectOnlyProformas"`
	SelectOnlyPeriodic           bool      `json:"selectOnlyPeriodic"`
	SelectOnlyCompleted          bool
	SelectOnlyUncompleted        bool     // in progress
	SelectIncludingDeletedOrders bool     // including deleted
	SelectOnlyDeletedOrders      bool     // deleted
	CompanyId                    uint     `json:"companyId"`
	ExternalUserId               uint     `json:"externalUserId"`
	StatusIds                    []uint   `json:"statusIds"`
	StatusCodes                  []string `json:"statusCodes"`
	StatusLetters                []string `json:"statusLetters"`
	Filters                      []ItemFilter
	Token                        string `json:"token"`
	QueryType                    string `json:"queryType"`
	OrderIds                     []uint `json:"orderIds"`
}
type OrdersResponse struct {
	Success          bool               `json:"success"`
	Orders           []Order            `json:"orders"`
	OrdersCount      int64              `json:"ordersCount"`
	Currency         string             `json:"currency"`
	TotalsByCurrency map[string]float64 `json:"totalsByCurrency"`
}

type OrderStatusesRequest struct {
	// to be filled
}
type OrderStatusesResponse struct {
	Statuses []OrderStatus `json:"statuses"`
	Count    int64         `json:"count"`
}

type PriceRange = common.PriceRange

type ProductsRequest struct {
	SelectInEshop      bool
	SelectInWarehouse  bool
	Warehouses         []string
	SelectUncountable  bool
	SelectArchived     bool
	CategoryUri        string
	CategoryIds        []uint
	PerPage            int
	Page               int
	Brands             []string
	PriceRange         PriceRange
	OrderWay           string
	OrderBy            string
	SearchQuery        string
	Filters            []ItemFilter
	ProductIds         []uint
	SkipQuantityColumn bool
	UnlimitedProducts  bool
	UseIndexedSearch   bool
	SelectMerged       bool // merged products by attributes
	// When selecting merged products (main products of variants) - products could also be expanded and selected with variants
	ExpandVariations bool
	// Created from/to, updated from/to
	CreatedAtFrom          time.Time
	CreatedAtTo            time.Time
	UpdatedAtFrom          time.Time
	UpdatedAtTo            time.Time
	CreatedOrUpdatedAtFrom time.Time
	CreatedOrUpdatedAtTo   time.Time
}
type ProductsResponse struct {
	Success             bool       `json:"success"`
	Products            []Product  `json:"products"`
	ProductsCount       int64      `json:"productsCount"`
	AvailablePriceRange PriceRange `json:"availablePriceRange"`
	Brands              []string   `json:"brands"`
}

type MediaItemOrder = v2.MediaItemOrder
type ProductImage = v2.ProductImage

type ProductCreateEditRequest struct {
	Product             Product           `json:"product"`
	MetaValues          map[string]string `json:"metaValues"`
	SelectedCategoryIds []int32           `json:"selectedCategoryIds"`
	// Only for new images
	Images []ProductImage `json:"images"`
	// For deleting existing images
	DeleteImagesKeys []int32 `json:"deleteImagesKeys"`
}

type ProductCategoriesRequest = common.ProductCategoriesRequest
type ProductCategoriesResponse = common.ProductCategoriesResponse

type ProductStocksRequest struct {
	Warehouses       []string
	ProductIds       []uint
	PurchaseDateFrom time.Time
	PurchaseDateTo   time.Time
	SearchQuery      string
	Filters          []ItemFilter
	GroupBy          []string
	PerPage          int
}

type ProductStock struct {
	ProductId         uint      `json:"productId"`
	Warehouse         string    `json:"warehouse"`
	Quantity          float64   `json:"quantity"`
	CostPrice         float64   `json:"costPrice"`
	CostCurrency      string    `json:"costCurrency"`
	CostSum           float64   `json:"costSum"`
	LineId            int32     `json:"lineId"`
	StockId           int32     `json:"stockId"`
	InvoiceId         int32     `json:"invoiceId"`
	Date              time.Time `json:"date"`
	AccountNo         string    `json:"accountNo"`
	MeasurementUnit   string    `json:"measurementUnit"`
	CompanyName       string    `json:"companyName"`
	CompanyId         uint      `json:"companyId"`
	RetailPrice       float64   `json:"retailPrice"`
	RetailCurrency    string    `json:"retailCurrency"`
	WholesalePrice    float64   `json:"wholesalePrice"`
	WholesaleCurrency string    `json:"wholesaleCurrency"`
	ProductInfo       Product   `json:"productInfo"`
	CategoryNames     []string  `json:"categories"`
	AverageCostPrice  float64   `json:"averageCostPrice"`
}

type ProductStocksResponse struct {
	Stocks             []ProductStock
	TotalQuantity      float64
	TotalSum           float64
	TotalCurrency      string
	TotalsByWarehouses map[string]*StockTotal
	TotalCount         int64
}

type StockTotal = v2.StockTotal
type WarehouseProductQuantity = v2.WarehouseProductQuantity
type ProductWarehousesQuantitiesResponse = v2.ProductWarehousesQuantitiesResponse
type StockAdjustmentReason = v2.StockAdjustmentReason

const (
	StockAdjustmentReasonInitialBalance = v2.StockAdjustmentReasonInitialBalance
	StockAdjustmentReasonCycleCount     = v2.StockAdjustmentReasonCycleCount
	StockAdjustmentReasonDamage         = v2.StockAdjustmentReasonDamage
	StockAdjustmentReasonShrinkage      = v2.StockAdjustmentReasonShrinkage
	StockAdjustmentReasonCorrection     = v2.StockAdjustmentReasonCorrection
	StockAdjustmentReasonReceived       = v2.StockAdjustmentReasonReceived
	StockAdjustmentReasonReturned       = v2.StockAdjustmentReasonReturned
)

type ProductStockSetRequest = v2.ProductStockSetRequest
type ProductStockSetEntry = v2.ProductStockSetEntry
type ProductStockAdjustRequest = v2.ProductStockAdjustRequest
type ProductStockAdjustEntry = v2.ProductStockAdjustEntry
type ProductStockAdjustmentResponse = v2.ProductStockAdjustmentResponse
type StockAdjustmentInvoiceInfo = v2.StockAdjustmentInvoiceInfo
type StockAdjustmentError = v2.StockAdjustmentError

type InvoiceSend = common.InvoiceSend
type InvoiceCreateUpdateRequest = common.InvoiceCreateUpdateRequest

type InvoiceErrorResponse = common.InvoiceErrorResponse
type OrderErrorResponse = common.OrderErrorResponse

type Country = common.Country
type CommerceChannel = common.CommerceChannel
type IntegrationSyncRecord = common.IntegrationSyncRecord
type IntegrationSyncRecordsRequest = common.IntegrationSyncRecordsRequest
type PluginSyncState = common.PluginSyncState
type PluginSyncStateRequest = common.PluginSyncStateRequest
type PluginSyncStateResponse = common.PluginSyncStateResponse

type PosDiscountCard struct {
	ID              int       `json:"cardId"`
	CardNumber      string    `json:"cardNumber"`
	DiscountPercent float64   `json:"discountPercent"`
	Email           string    `json:"email"`
	HasAgreedEmail  bool      `json:"hasAgreedEmail"`
	FirsName        string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Branch          string    `json:"branch"`
	DivisionId      uint      `json:"divisionId"`
	ValidUntil      time.Time `json:"validUntil"`
	PhoneNumber     string    `json:"phoneNumber"`
	BirthDate       time.Time `json:"birthDate"`
	Address         string    `json:"address"`
	Sex             int       `json:"sex"`
	HasAgreedSms    bool      `json:"hasAgreedSms"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedUserId   int       `json:"createdUserId"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedUserId   uint      `json:"updatedUserId"`
}
type PosDiscountCardsResponse struct {
	Cards      []PosDiscountCard
	CardsCount int64
}
type PosDiscountCardsRequest struct {
	CreatedAtFrom          time.Time
	CreatedAtTo            time.Time
	UpdatedAtFrom          time.Time
	UpdatedAtTo            time.Time
	CreatedOrUpdatedAtFrom time.Time
	CreatedOrUpdatedAtTo   time.Time
	Branch                 string
	DivisionId             uint
	PerPage                int
	Page                   int
	SearchQuery            string
}

type CustomFieldType = common.CustomFieldType

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

type CustomField = common.CustomField
type CustomFieldsRequest = common.CustomFieldsRequest
type CustomFieldsResponse = common.CustomFieldsResponse
type FileUploadRequest = common.FileUploadRequest
type FileUploadResponse = common.FileUploadResponse
type EmailSender = common.EmailSender
type EmailRequest = common.EmailRequest

type DocumentChargeCreateRequest = v2.DocumentChargeCreateRequest
type LineChargeCreateRequest = v2.LineChargeCreateRequest

// Product feature sync DTOs live in the common package (shared by the v1 and v2
// plugin APIs). Concrete/flat types — no gob.Register needed. Org/user come from
// the RequestContext, not these structs.
type UpsertFeatureRequest = common.UpsertFeatureRequest
type UpsertFeatureResponse = common.UpsertFeatureResponse
type FeatureValueItem = common.FeatureValueItem
type ProductFeatureSet = common.ProductFeatureSet
type ProductFeatureSyncItem = common.ProductFeatureSyncItem
type ReplaceProductFeaturesRequest = common.ReplaceProductFeaturesRequest
type ReplaceProductFeaturesResult = common.ReplaceProductFeaturesResult
