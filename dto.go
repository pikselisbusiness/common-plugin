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
