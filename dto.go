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
	AdvanceSumMap  map[string]float64 `json:"advanceSums"`
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

	Company    Company `json:"company,omitempty"`
	IsVerified bool    `json:"isVerified"`
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
}
type InvoicePosData struct {
	CashierId  uint   `json:"cashierId"`
	PosType    string `json:"posType"`
	ReceiptNo  string `json:"receiptNo"`
	ReceiptNo2 string `json:"receiptNo2"`
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
