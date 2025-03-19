package shared

import (
	"encoding/json"
	"time"
)

type UserInfo struct {
	UserId       uint   `json:"userId"`
	UserImageUrl string `json:"userImageUrl"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}
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
	InvoiceId      int32                `json:"invoiceId"`
	IsPurchase     bool                 `json:"isPurchase"`
	Operation      string               `json:"operation"`
	CompanyName    string               `json:"companyName"`
	CompanyId      int32                `json:"companyId"`
	DocSeries      string               `json:"docSeries"`
	Document       string               `json:"document"`
	Document2      string               `json:"document2"`
	InnerDocument  string               `json:"innerDocument"`
	InvoiceDate    time.Time            `json:"invoiceDate"`
	TotalVat       float64              `json:"totalVat"`
	TotalWoVat     float64              `json:"totalWoVat"`
	TotalWithVat   float64              `json:"totalWithVat"`
	TotalCurrency  string               `json:"totalCurrency"`
	InvoicePaid    bool                 `json:"invoicePaid"`
	DebtSum        float64              `json:"debtSum"`
	PaidSum        float64              `json:"paidSum"`
	PaidLastDate   time.Time            `json:"paidLastDate"`
	PayUntilDate   time.Time            `json:"payUntilDate"`
	ReceiveDate    time.Time            `json:"receiveDate"`
	TotalCostSum   float64              `json:"totalCostSum"`
	TotalGainSum   float64              `json:"totalGainSum"`
	FromWarehouse  string               `json:"fromWarehouse"`
	ToWarehouse    string               `json:"toWarehouse"`
	Comment        string               `json:"comment"`
	FreezeFinances bool                 `json:"freezeFinances"`
	Costs          []InvoiceCost        `json:"costs"`
	ExtraSettings  ExtraInvoiceSettings `json:"extraSettings"`
	PosData        InvoicePosData       `json:"posData"`

	InsertDatetime time.Time `json:"insertDatetime"`
	ModifyDatetime time.Time `json:"updateDatetime"`
	InsertUserInfo UserInfo  `json:"insertUserInfo"`
	UpdateUserInfo UserInfo  `json:"updateUserInfo"`

	Company      Company         `json:"company,omitempty"`
	IsVerified   bool            `json:"isVerified"`
	Status       string          `json:"status"`
	CustomFields json.RawMessage `json:"customFields"`
}
type ExtraInvoiceSettings struct {
	IsafForm bool `json:"isafForm"`
}
type InvoiceCost struct {
	Count           bool    `json:"count"`
	Name            string  `json:"name"`
	Sum             float64 `json:"sum"`
	Currency        string  `json:"currency"`
	DebitAccountNo  string  `json:"debitAccountNo"`
	CreditAccountNo string  `json:"creditAccountNo"`
}
type InvoiceProduct struct {
	ProductName         string  `json:"name"` // for out
	ProductId           uint    `json:"productId"`
	Quantity            float64 `json:"quantity"`
	LeftQuantity        float64 `json:"leftQuantity"` // Used for purchase
	Price               float64 `json:"price"`
	PriceIsWithVat      bool    `json:"priceIsWithVat"`
	CostPrice           float64 `json:"costPrice"`
	CostCurrency        string  `json:"costCurrency"`
	VatPercentage       float64 `json:"vatPercentage"`
	VatClass            string  `json:"vatClass"`
	Articule            string  `json:"articule"`
	Description         string  `json:"description"`
	UseStockId          bool    `json:"useStockId"`
	StockId             int32   `json:"stockId"`
	LineId              int32   `json:"lineId"`    // for out
	Countable           bool    `json:"countable"` // for out
	Sum                 float64 `json:"sum"`
	UseCostPriceAsPrice bool
	DiscountValue       float64 `json:"discountValue"`
	DiscountType        string  `json:"discountType"`
	// For returns - this is lineId of original line that is returned
	ParentLineId  int32 `json:"parentLineId"`
	RelatedLineId int32 `json:"relatedLineId"`
}
type InvoicePosData struct {
	CashierId  uint   `json:"cashierId"`
	PosType    string `json:"posType"`
	ReceiptNo  string `json:"receiptNo"`
	ReceiptNo2 string `json:"receiptNo2"`
}

type InvoiceReference struct {
	InvoiceId         uint   `json:"invoiceId"`
	ReferredInvoiceId uint   `json:"referredInvoiceId"`
	ReferenceType     string `json:"referenceType"`
}

/**
---- DIVISIONS
*/

type DivisionsRequest struct {
	UsedForInvoiceReports bool
	UsedForRepairs        bool
	SearchQuery           string // unused
}

type Division struct {
	DivisionId      int32  `json:"divisionId"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
	HouseNumber     string `json:"houseNumber"`
	ApartmentNumber string `json:"apartmentNumber"`
	PostCode        string `json:"postCode"`
	// FullAddress           string              `json:"updateUserId"`
	UsedForRepairs        bool                `json:"usedForRepairs"`
	UsedForInvoiceReports bool                `json:"usedForInvoiceReports"`
	IsDefault             bool                `json:"isDefault"`
	InsertDatetime        time.Time           `json:"insertDatetime"`
	InsertUserId          uint                `json:"insertUserId"`
	UpdateUserId          uint                `json:"updateUserId"`
	UpdateDatetime        time.Time           `json:"updateDatetime"`
	Warehouses            []DivisionWarehouse `json:"warehouses"`
	WarehouseIds          []uint              `json:"warehouseIds"`
	PosType               string              `json:"posType"`
}
type DivisionWarehouse struct {
	Warehouse   string `json:"warehouse"`
	Name        string `json:"name"`
	WarehouseId int32  `json:"warehouseId"`
}

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
	Image1           string  `json:"image1"`
	Image2           string  `json:"image2"`
	Image3           string  `json:"image3"`
	Image4           string  `json:"image4"`
	Image5           string  `json:"image5"`
	Image6           string  `json:"image6"`
	Image7           string  `json:"image7"`
	Image8           string  `json:"image8"`
	Image9           string  `json:"image9"`
	Image10          string  `json:"image10"`
	Image11          string  `json:"image11"`
	Image12          string  `json:"image12"`
	Image13          string  `json:"image13"`
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

	Quantity     int64    `json:"quantity"`
	ImageMedium  string   `json:"image"`
	ImageOrigin  string   `json:"imageOrigin"`
	ImageSmall   string   `json:"imageSmall"`
	Images       []string `json:"images"`
	ImagesOrigin []string `json:"imagesOrigin"`
	ImagesSmall  []string `json:"imagesSmall"`

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
	InsertUserId   uint                `json:"-"` // inner only
	UpdateUserId   uint                `json:"-"` // inner only
	InsertDatetime time.Time           `json:"insertDatetime"`
	UpdateDatetime time.Time           `json:"updateDatetime"`
	InsertUserInfo UserInfo            `json:"insertUserInfo"`
	UpdateUserInfo UserInfo            `json:"updateUserInfo"`
	External       ProductExternalInfo `json:"external"`

	PhotosUpdated bool
	CustomFields  json.RawMessage `json:"customFields"`
}
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
type ProductAccountingNumbers struct {
	ReturnAccountNo              float64 `json:"returnAccountNo"`
	CostDebitAccountNo           string  `json:"costDebitAccountNo"`
	CostCreditAccountNo          string  `json:"costCreditAccountNo"`
	SaleAccountNo                string  `json:"saleAccountNo"`
	SaleCostAccountNo            string  `json:"saleCostAccountNo"`
	SupplierDebtAccountNo        string  `json:"supplierDebtAccountNo"`
	BuyerDebtAccountNo           string  `json:"buyerDebtAccountNo"`
	VatDebitAccountNo            string  `json:"vatDebitAccountNo"`
	VatCreditAccountNo           string  `json:"vatCreditAccountNo"`
	WriteOffAccountNo            string  `json:"writeOffAccountNo"`
	ManufacturingDebitAccountNo  string  `json:"manufacturingDebitAccountNo"`
	ManufacturingCreditAccountNo string  `json:"manufacturingCreditAccountNo"`
}
type ProductAttributeVariation struct {
	Attributes    []ProductAttribute `json:"attributes"`
	IsMainProduct bool               `json:"isMainProduct"`
}
type ProductAttribute struct {
	AttributeId        uint   `json:"attributeId" `
	AttributeName      string `json:"attributeName"`
	AttributeColorCode string `json:"attributeColorCode"`
	AttributeType      string `json:"attributeType"`
	GroupId            uint   `json:"groupId"`
	PrestaId           uint   `json:"prestaId"`
}

type ProductCategoriesShort struct {
	CategoryId      int32  `json:"categoryId"`
	Name            string `json:"name"`
	ParentName      string `json:"parentName"`
	ParentId        int32  `json:"parentId"`
	Level           int32  `json:"level"`
	Link            string `json:"link"`
	LinkAbsolute    string `json:"linkAbsolute"`
	MetaTitle       string `json:"metaTitle"`
	MetaDescription string `json:"metaDescription"`
	H1Title         string `json:"h1Title"`
	Image           string `json:"image"`
	ImageSmall      string `json:"imageSmall"`
	OrderBy         int32  `json:"orderBy"`
	ProductsCount   int32  `json:"productsCount"`
}

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
	LocationId       int32  `json:"locationId"`
	LocationName     string `json:"locationName"`
	LocationCity     string `json:"locationCity"`
	LocationPostCode string `json:"locationPostCode"`
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

	ProductId       uint    `json:"productId"`
	ProductName     string  `json:"productName"`
	ProductInfo     Product `json:"productInfo"`
	Quantity        float64 `json:"quantity"`
	OriginalPrice   float64 `json:"originalPrice"`
	CostPrice       float64 `json:"costPrice"`
	DealerPrice     float64 `json:"dealerPrice"`
	BuyerPrice      float64 `json:"price"`
	PriceWithVat    float64 `json:"priceWithVat"`
	Sum             float64 `json:"sum"`
	SumWithVat      float64 `json:"sumWithVat"`
	DiscountAmount  float64 `json:"discountAmount"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountSum     float64 `json:"discountSum"`
	Articule        string  `json:"articule"`
	CostCurrency    string  `json:"costCurrency"`
	DealerCurrency  string  `json:"dealerCurrency"`
	BuyerCurrency   string  `json:"currency"`
	MeasurementUnit string  `json:"measurementUnit"`
	VatPercentage   float64 `json:"vatPercentage"`
	VatClass        string  `json:"vatClass"`
	Package         string  `json:"package"`
	IsTransport     bool    `json:"isTransport"`
}
type OrderStatus struct {
	StatusId              int32  `json:"statusId"`
	Selector              int32  `json:"selector"`
	LabelLt               string `json:"labelLt"`
	LabelEn               string `json:"labelEn"`
	LabelRu               string `json:"labelRu"`
	IsInvoiceDownloadable bool   `json:"isInvoiceDownloadable"`
	SendEmail             int    `json:"sendEmail"`
	DoNotPromptEmailText  bool   `json:"doNotPromptEmailText"`
	IsOrderCompleted      bool   `json:"isOrderCompleted"`
	DocSeries             string `json:"docSeries"`
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
	// Possible types - contact, invoiceaddress, deliveryaddress, followupaddress, otheraddress
	ContactType string `json:"contactType"`
	Street      string `json:"street"`
	State       string `json:"state"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PostCode    string `json:"postCode"`
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
}
type ItemFilter struct {
	Operator string `json:"operator" query:"operator"`
	Value    string `json:"value" query:"value"`
	Field    string `json:"field" query:"field"`
}
type OrderCreateRequest struct {
	OrderId                int32                `json:"orderId"`
	Date                   time.Time            `json:"date"`
	ExpiryDate             time.Time            `json:"expiryDate"`
	CompanyId              int32                `json:"companyId"`
	IsOrder                bool                 `json:"isOrder"`
	IsProforma             bool                 `json:"isProforma"`
	Type                   string               `json:"type"`
	Comments               string               `json:"comments"`
	CurrencyId             int32                `json:"currencyId"`
	Currency               string               `json:"currency"`
	Period                 int32                `json:"period"`
	StatusLetter           string               `json:"statusLetter"`
	Title                  string               `json:"title"`
	ExternalUserId         uint                 `json:"externalUserId"`
	Email                  string               `json:"email"`
	ShippingAddress        OrderShippingAddress `json:"shippingAddress"`
	PickupLocation         OrderPickupLocation  `json:"pickupLocation"`
	PaymentId              uint                 `json:"paymentId"`
	PaymentType            string               `json:"paymentType"`
	PaymentIsCod           bool                 `json:"paymentIsCod"`
	PaymentInnerType       string               `json:"paymentInnerType"`
	ShippingId             uint                 `json:"shippingId"`
	ShippingMethodSelector string               `json:"shippingMethodSelector"`
	ShippingMethodName     string               `json:"shippingMethodName"`
	TrackingNumber         string               `json:"trackingNumber"`
	TrackingType           string               `json:"trackingType"`
	ReportType             string               `json:"reportType"` // PDF type
	PdfBlankId             int32                `json:"pdfBlankId"`
	Products               []OrderLine          `json:"products"`
	Token                  string               `json:"token"`
	PaymentReferenceNo     string               `json:"paymentReferenceNo"`
	CustomFields           json.RawMessage      `json:"customFields"`
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
	SelectOnlyUncompleted        bool   // in progress
	SelectIncludingDeletedOrders bool   // including deleted
	SelectOnlyDeletedOrders      bool   // deleted
	CompanyId                    uint   `json:"companyId"`
	ExternalUserId               uint   `json:"externalUserId"`
	StatusIds                    []uint `json:"statusIds"`
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
type PriceRange struct {
	PriceFrom        float64 `json:"priceFrom"`
	PriceTo          float64 `json:"priceTo"`
	PriceFromWithVat float64 `json:"priceFromWithVat"`
	PriceToWithVat   float64 `json:"priceToWithVat"`
}

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
type ProductImage struct {
	FileName string `json:"fileName"`
	FileBlob []byte `json:"fileBlob"`
}
type ProductCreateEditRequest struct {
	Product             Product           `json:"product"`
	MetaValues          map[string]string `json:"metaValues"`
	SelectedCategoryIds []int32           `json:"selectedCategoryIds"`
	// Only for new images
	Images []ProductImage `json:"images"`
	// For deleting existing images
	DeleteImagesKeys []int32 `json:"deleteImagesKeys"`
}

type ProductCategoriesRequest struct {
	SelectInEshop bool
}
type ProductCategoriesResponse struct {
	Categories []ProductCategoriesShort `json:"categories"`
}

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
type StockTotal struct {
	TotalSum      float64 `json:"totalSum"`
	TotalCurrency string  `json:"totalCurrency"`
	TotalQuantity float64 `json:"totalQuantity"`
	Warehouse     string  `json:"warehouse"`
}
type ProductStocksResponse struct {
	Stocks             []ProductStock
	TotalQuantity      float64
	TotalSum           float64
	TotalCurrency      string
	TotalsByWarehouses map[string]*StockTotal
	TotalCount         int64
}

type InvoiceSend struct {
	ToSend bool `json:"toSend"`
}
type InvoiceCreateUpdateRequest struct {
	InvoiceId            int32                `json:"invoiceId"` // FOR update only
	IsPurchase           bool                 `json:"isPurchase"`
	DocSeries            string               `json:"docSeries"`
	Document             string               `json:"document"`
	Document2            string               `json:"document2"`
	CompanyId            int32                `json:"companyId"`
	TotalCurrency        string               `json:"totalCurrency"`
	InvoiceDate          time.Time            `json:"invoiceDate"`
	PayUntilDate         time.Time            `json:"payUntilDate"`
	Comment              string               `json:"comment"`
	ReceiptNo            string               `json:"receiptNo"`
	PdfBlankId           int32                `json:"pdfBlankId"`
	IsafForm             bool                 `json:"isafForm"`
	ReceiveDate          time.Time            `json:"receiveDate"`
	WorkerId             int32                `json:"workerId"`
	FromWarehouse        string               `json:"fromWarehouse"`
	ToWarehouse          string               `json:"toWarehouse"`
	Products             []InvoiceProduct     `json:"products"`
	DeleteProductLineIds []int32              `json:"deleteProductLineIds"` // FOR update only
	Operation            string               `json:"operation"`
	Costs                []InvoiceCost        `json:"costs"`
	AutoSend             InvoiceSend          `json:"autoSend"`
	ExtraSettings        ExtraInvoiceSettings `json:"extraSettings"`
	PosData              InvoicePosData       `json:"posData"`
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

type Country struct {
	Country         string    `json:"country"`
	ShortCode       string    `json:"shortCode"`
	CountryId       int32     `json:"countryId"`
	CreatedAt       time.Time `json:"createdAt"`
	CratedUserInfo  UserInfo  `json:"createdUserInfo"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedUserInfo UserInfo  `json:"updatedUserInfo"`
}

type IntegrationSyncRecord struct {
	SyncId         uint      `json:"syncId"`
	Type           string    `json:"type"`
	Entity         string    `json:"entity"`
	EntityField    string    `json:"entityField"`
	LoginName      string    `json:"loginName"`
	Info           string    `json:"info"`
	DeviceId       string    `json:"deviceId"`
	OrganizationId uint      `json:"organizationId"`
	SyncTime       time.Time `json:"syncTime"`
}

type IntegrationSyncRecordsRequest struct {
	Page         int `json:"page"`
	PerPage      int `json:"perPage"`
	Type         string
	Entity       string `json:"entity"`
	EntityField  string `json:"syncField"`
	DeviceId     string `json:"deviceId"`
	SyncTimeFrom time.Time
	SyncTimeTo   time.Time
	LoginName    string
	Info         string
	OrderBy      string // by default by syncTime
	OrderWay     string // by default - DESC
}

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

type CustomFieldType string

func (s CustomFieldType) IsValid() bool {
	for _, fieldType := range ValidCustomFieldTypes {
		if fieldType == s {
			return true
		}
	}
	return false
}

var ValidCustomFieldTypes = []CustomFieldType{
	CustomFieldString, CustomFieldText, CustomFieldFloat, CustomFieldDate, CustomFieldDatetime, CustomFieldBoolean, CustomFieldSelect,
}

var (
	CustomFieldString   CustomFieldType = "string"
	CustomFieldText     CustomFieldType = "text"
	CustomFieldFloat    CustomFieldType = "float"
	CustomFieldDate     CustomFieldType = "date"
	CustomFieldDatetime CustomFieldType = "datetime"
	CustomFieldBoolean  CustomFieldType = "boolean"
	CustomFieldSelect   CustomFieldType = "select"
)

type CustomField struct {
	ID              uint            `json:"-"`
	OrganizationId  uint            `json:"-"`
	CustomFieldId   uint            `json:"customFieldId"`
	RelatedType     string          `json:"relatedType"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	Required        bool            `json:"required"`
	Type            CustomFieldType `json:"type"`
	Options         json.RawMessage `json:"options"`
	DisplayInline   bool            `json:"displayInline"`
	FieldOrder      uint            `json:"fieldOrder"`
	Active          bool            `json:"active"`
	ShowInList      bool            `json:"showInList"`
	ShowOnlyAdmin   bool            `json:"showOnlyAdmin"`
	IsIndexedField  bool            `json:"isIndexedField"`
	CreatedAt       time.Time       `json:"createdAt"`
	CreatedUserInfo UserInfo        `json:"createdUserInfo"`
	UpdatedUserInfo UserInfo        `json:"updatedUserInfo"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

type CustomFieldsRequest struct {
	CustomFieldIds []uint
	Names          []string
	Slugs          []string
	RelatedType    string
	SearchQuery    string // unused
}
type CustomFieldsResponse struct {
	Fields []CustomField
}
type EmailSender struct {
	Host     string
	Name     string
	Email    string
	Username string
	Password string
	Bcc      string
}
type EmailRequest struct {
	ToName          string
	ToEmail         string
	Cc              []string
	Subject         string
	Html            string
	UseVariables    bool
	Attachments     map[string][]byte
	UseCustomSender bool
	Sender          EmailSender
}
