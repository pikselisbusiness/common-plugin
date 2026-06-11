package common

import (
	"encoding/json"
	"time"
)

type DocumentOperation struct {
	OperationId uint   `json:"operationId"`
	Operation   string `json:"operation"`
	// Whether operation is purchase
	OperPurchase bool `json:"operPurchase"`
	// Whether operation is sale
	OperSale bool `json:"operSale"`
	// Whether operation is with VAT
	OperVat      string `json:"operVat"`
	OperSupplier bool   `json:"operSupplier"`
	OperBuyer    bool   `json:"operBuyer"`
	// Whether operation needs payment
	OperPay                   bool      `json:"operPay"`
	OperCost                  bool      `json:"operCost"`
	OperMark                  bool      `json:"operMark"`
	OperReturn                bool      `json:"operReturn"`
	OperWarehouse             bool      `json:"operWarehouse"`
	OperStockTaking           bool      `json:"operStockTaking"`
	OperWriteOff              bool      `json:"operWriteOff"`
	CostDebitAccountNo        string    `json:"costDebitAccountNo"`
	CostCreditAccountNo       string    `json:"costCreditAccountNo"`
	SaleDebitAccountNo        string    `json:"saleDebitAccountNo"`
	SaleCreditAccountNo       string    `json:"saleCreditAccountNo"`
	VatDebitAccountNo         string    `json:"vatDebitAccountNo"`
	VatCreditAccountNo        string    `json:"vatCreditAccountNo"`
	ReverseVatDebitAccountNo  string    `json:"reverseVatDebitAccountNo"`
	ReverseVatCreditAccountNo string    `json:"reverseVatCreditAccountNo"`
	DebtDebitAccountNo        string    `json:"debtDebitAccountNo"`
	DebtCreditAccountNo       string    `json:"debtCreditAccountNo"`
	CreatedAt                 time.Time `json:"createdAt"`
	CreatedUserId             uint      `json:"createdUserId"`
	UpdatedAt                 time.Time `json:"updatedAt"`
	UpdatedUserId             uint      `json:"updatedUserId"`
}

type DocumentOperationsRequest struct {
	OperationType string `json:"operationType"`
}

type DocumentOperationsResponse struct {
	Operations []DocumentOperation `json:"operations"`
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
	ProductName         string   `json:"name"`
	ProductId           uint     `json:"productId"`
	Quantity            float64  `json:"quantity"`
	MeasurementUnit     string   `json:"measurementUnit"`
	LeftQuantity        float64  `json:"leftQuantity"`
	Price               float64  `json:"price"`
	PriceExclTax        float64  `json:"priceExclTax"`
	PriceInclTax        float64  `json:"priceInclTax"`
	PriceIsWithVat      bool     `json:"priceIsWithVat"`
	CostPrice           float64  `json:"costPrice"`
	CostCurrency        string   `json:"costCurrency"`
	VatPercentage       float64  `json:"vatPercentage"`
	VatClass            string   `json:"vatClass"`
	TaxCode             string   `json:"taxCode,omitempty"`
	TaxAmount           *float64 `json:"taxAmount,omitempty"`
	Articule            string   `json:"articule"`
	Description         string   `json:"description"`
	UseStockId          bool     `json:"useStockId"`
	StockId             int32    `json:"stockId"`
	LineId              int32    `json:"lineId"`
	Countable           bool     `json:"countable"`
	Sum                 float64  `json:"sum"`
	TotalExclTax        float64  `json:"totalExclTax"`
	TotalInclTax        float64  `json:"totalInclTax"`
	UseCostPriceAsPrice bool
	DiscountValue       float64 `json:"discountValue"`
	DiscountAmount      float64 `json:"discountAmount"`
	DiscountPerUnit     float64 `json:"discountPerUnit"`
	DiscountPercent     float64 `json:"discountPercent"`
	DiscountType        string  `json:"discountType"`
	ParentLineId        int32   `json:"parentLineId"`
	RelatedLineId       int32   `json:"relatedLineId"`
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

type InvoicePaymentsRequest struct {
	InvoiceId  uint      `json:"invoiceId"`
	InvoiceIds []uint    `json:"invoiceIds"`
	DateFrom   time.Time `json:"dateFrom"`
	DateTo     time.Time `json:"dateTo"`
	Page       int       `json:"page"`
	PerPage    int       `json:"perPage"`
}

type InvoicePayment struct {
	InvoiceId   uint      `json:"invoiceId"`
	Account     string    `json:"account"`
	Sum         float64   `json:"sum"`
	Currency    string    `json:"currency"`
	PaymentDate time.Time `json:"paymentDate"`
}

type InvoicePaymentsResponse struct {
	Success  bool             `json:"success"`
	Payments []InvoicePayment `json:"payments"`
	Count    int64            `json:"count"`
}

type JournalEntriesRequest struct {
	PerPage         int       `json:"perPage"`
	Page            int       `json:"page"`
	JournalEntryIds []uint    `json:"journalEntryIds"`
	EntryTypes      []string  `json:"entryTypes"`
	DateFrom        time.Time `json:"dateFrom"`
	DateTo          time.Time `json:"dateTo"`
	CustomerIds     []uint    `json:"customerIds"`
	InvoiceIds      []uint    `json:"invoiceIds"`
	AdvanceIds      []uint    `json:"advanceIds"`
	SearchQuery     string    `json:"searchQuery"`
}

type JournalEntryLine struct {
	ID             uint      `json:"id"`
	JournalEntryId uint      `json:"journalEntryId"`
	Date           time.Time `json:"date"`
	Debit          float64   `json:"debit"`
	Credit         float64   `json:"credit"`
	InvoiceId      uint      `json:"invoiceId"`
	CustomerId     uint      `json:"customerId"`
	CompanyName    string    `json:"companyName"`
	Sum            float64   `json:"sum"`
	Currency       string    `json:"currency"`
	AdvanceId      uint      `json:"advanceId"`
	AccountNo      string    `json:"accountNo"`
}

type JournalEntry struct {
	ID          uint               `json:"id"`
	EntryType   string             `json:"entryType"`
	EntryDate   time.Time          `json:"entryDate"`
	BankDate    time.Time          `json:"bankDate"`
	DocNo       string             `json:"docNo"`
	AccountNo   string             `json:"accountNo"`
	CustomerId  uint               `json:"customerId"`
	CompanyName string             `json:"companyName"`
	InvoiceId   uint               `json:"invoiceId"`
	AdvanceId   uint               `json:"advanceId"`
	TotalSum    float64            `json:"totalSum"`
	Currency    string             `json:"currency"`
	Lines       []JournalEntryLine `json:"lines"`
}

type JournalEntriesResponse struct {
	Success        bool           `json:"success"`
	JournalEntries []JournalEntry `json:"journalEntries"`
	Count          int64          `json:"count"`
	Currency       string         `json:"currency"`
}

type InvoiceSend struct {
	ToSend bool     `json:"toSend"`
	Emails []string `json:"emails"`
}

type InvoiceCreateUpdateRequest struct {
	InvoiceId            int32                `json:"invoiceId"`
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
	DeleteProductLineIds []int32              `json:"deleteProductLineIds"`
	Operation            string               `json:"operation"`
	Costs                []InvoiceCost        `json:"costs"`
	AutoSend             InvoiceSend          `json:"autoSend"`
	ExtraSettings        ExtraInvoiceSettings `json:"extraSettings"`
	PosData              InvoicePosData       `json:"posData"`
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

type ProductCategoriesRequest struct {
	SelectInEshop     bool
	SelectInPriceList bool
	SelectInWarehouse bool
}

type ProductCategoriesResponse struct {
	Categories []ProductCategoriesShort `json:"categories"`
}

type DivisionsRequest struct {
	UsedForInvoiceReports bool
	UsedForRepairs        bool
	SearchQuery           string
}

type Division struct {
	DivisionId            int32               `json:"divisionId"`
	Name                  string              `json:"name"`
	Country               string              `json:"country"`
	City                  string              `json:"city"`
	Street                string              `json:"street"`
	HouseNumber           string              `json:"houseNumber"`
	ApartmentNumber       string              `json:"apartmentNumber"`
	PostCode              string              `json:"postCode"`
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

type Country struct {
	Country         string    `json:"country"`
	ShortCode       string    `json:"shortCode"`
	CountryId       int32     `json:"countryId"`
	CreatedAt       time.Time `json:"createdAt"`
	CratedUserInfo  UserInfo  `json:"createdUserInfo"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedUserInfo UserInfo  `json:"updatedUserInfo"`
}

type CommerceChannel struct {
	ChannelID     uint      `json:"channelId"`
	ChannelUUID   string    `json:"channelUuid"`
	Name          string    `json:"name"`
	SourceSystem  string    `json:"sourceSystem"`
	ChannelType   string    `json:"channelType"`
	IsActive      bool      `json:"isActive"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedUserID uint      `json:"createdUserId"`
	UpdatedUserID uint      `json:"updatedUserId"`
}

type IntegrationSyncRecord struct {
	SyncId         uint            `json:"syncId"`
	Type           string          `json:"type"`
	Entity         string          `json:"entity"`
	EntityField    string          `json:"entityField"`
	LoginName      string          `json:"loginName"`
	Info           string          `json:"info"`
	DeviceId       string          `json:"deviceId"`
	OrganizationId uint            `json:"organizationId"`
	SyncTime       time.Time       `json:"syncTime"`
	Extra          json.RawMessage `json:"extra"`
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
	OrderBy      string
	OrderWay     string
}

type PluginSyncState struct {
	SyncId    uint      `json:"syncId"`
	Plugin    string    `json:"plugin"`
	Entity    string    `json:"entity"`
	EntityId  string    `json:"entityId"`
	ToSync    bool      `json:"toSync"`
	Done      bool      `json:"done"`
	Error     bool      `json:"error"`
	ErrorText string    `json:"errorText"`
	SyncTime  time.Time `json:"syncTime"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PluginSyncStateRequest struct {
	Page     int    `json:"page"`
	PerPage  int    `json:"perPage"`
	Plugin   string `json:"plugin"`
	Entity   string `json:"entity"`
	EntityId string `json:"entityId"`
	ToSync   bool   `json:"toSync"`
	Done     bool   `json:"done"`
	Error    bool   `json:"error"`
	OrderBy  string `json:"orderBy"`
	OrderWay string `json:"orderWay"`
}

type PluginSyncStateResponse struct {
	Success bool              `json:"success"`
	States  []PluginSyncState `json:"states"`
	Count   int64             `json:"count"`
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
	ShowInEcommerce bool            `json:"showInEcommerce"`
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
	SearchQuery    string
}

type CustomFieldsResponse struct {
	Fields []CustomField
}

type FileUploadRequest struct {
	ParentId    uint           `json:"parentId"`
	ParentType  string         `json:"parentType"`
	FileName    string         `json:"fileName"`
	MimeType    string         `json:"mimeType"`
	Content     []byte         `json:"content,omitempty"`
	IsPublic    bool           `json:"isPublic"`
	Options     map[string]any `json:"options,omitempty"`
	Type        string         `json:"type,omitempty"`
	ExternalUrl string         `json:"externalUrl,omitempty"`
}

type FileUploadResponse struct {
	UploadId    uint   `json:"uploadId"`
	PreviewPath string `json:"previewPath"`
	Type        string `json:"type"`
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
