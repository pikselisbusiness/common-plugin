package v2

import (
	"encoding/json"
	"time"
)

// Company mirrors the newer external company shape used by plugin API v2.
type Company struct {
	CompanyId          int32  `json:"companyId"`
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	TaxNumber          string `json:"taxNumber"`

	IsPerson     bool   `json:"isPerson"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	AddressLine3 string `json:"addressLine3"`
	PostCode     string `json:"postCode"`
	City         string `json:"city"`
	CountryCode  string `json:"countryCode"`
	CountryName  string `json:"countryName"`
	PhoneNumber  string `json:"phoneNumber"`

	WarnedDate        time.Time `json:"warnedDate"`
	DeferSaleDays     int32     `json:"deferSaleDays"`
	DeferPurchaseDays int32     `json:"deferPurchaseDays"`
	RequireDays       int32     `json:"requireDays"`
	ExecuteDays       int32     `json:"executeDays"`
	Region            string    `json:"region"`
	Representative    string    `json:"representative"`
	Website           string    `json:"website"`
	Email             string    `json:"email"`
	IsSupplier        bool      `json:"isSupplier"`
	IsBuyer           bool      `json:"isBuyer"`
	IsBranch          bool      `json:"isBranch"`
	IsManufacturer    bool      `json:"isManufacturer"`

	ContractNumber    string    `json:"contractNumber"`
	ContractStartDate time.Time `json:"contractStartDate"`
	ContractEndDate   time.Time `json:"contractEndDate"`

	MaxDebt                   float64         `json:"maxDebt"`
	MaxCurrency               string          `json:"maxDebtCurrency"`
	DeferSaleMonths           int32           `json:"deferSaleMonths"`
	DeferSaleType             string          `json:"deferSaleType"`
	DeferPurchaseMonths       int32           `json:"deferPurchaseMonths"`
	DeferPurchaseType         string          `json:"deferPurchaseType"`
	SalePaymentTermId         uint            `json:"salePaymentTermId"`
	PurchasePaymentTermId     uint            `json:"purchasePaymentTermId"`
	WantsInvoiceByRegularMail bool            `json:"wantsInvoiceByRegularMail"`
	PriceField                string          `json:"priceField"`
	CustomFields              json.RawMessage `json:"customFields"`
	TaxExempt                 bool            `json:"taxExempt"`
	TaxProfile                json.RawMessage `json:"taxProfile"`
	CustomerLanguage          string          `json:"customerLanguage"`
	Source                    string          `json:"source"`
	IsVerified                bool            `json:"isVerified"`
	CreatedAt                 time.Time       `json:"createdAt"`
	CreatedUserId             uint            `json:"createdUserId"`
	UpdatedAt                 time.Time       `json:"updatedAt"`
	UpdatedUserId             uint            `json:"updatedUserId"`
}

type TaxProfileRequest struct {
	RequestType       *string `json:"requestType,omitempty"`
	Country           *string `json:"country,omitempty"`
	Region            *string `json:"region,omitempty"`
	TaxId             *string `json:"taxId,omitempty"`
	CertificateNumber *string `json:"certificateNumber,omitempty"`
	ValidFrom         *string `json:"validFrom,omitempty"`
	ValidUntil        *string `json:"validUntil,omitempty"`
}

type CompanyCreateUpdateRequest struct {
	Name                  string             `json:"name"`
	RegistrationNumber    string             `json:"registrationNumber"`
	TaxNumber             string             `json:"taxNumber"`
	IsPerson              bool               `json:"isPerson"`
	AddressLine1          string             `json:"addressLine1"`
	AddressLine2          string             `json:"addressLine2"`
	AddressLine3          string             `json:"addressLine3"`
	PostCode              string             `json:"postCode"`
	City                  string             `json:"city"`
	CountryCode           string             `json:"countryCode"`
	PhoneNumber           string             `json:"phoneNumber"`
	WarnedDate            time.Time          `json:"warnedDate"`
	DeferSaleDays         int32              `json:"deferSaleDays"`
	DeferPurchaseDays     int32              `json:"deferPurchaseDays"`
	Region                string             `json:"region"`
	Representative        string             `json:"representative"`
	Fax                   string             `json:"fax"`
	Website               string             `json:"website"`
	Email                 string             `json:"email"`
	IsBranch              bool               `json:"isBranch"`
	IsManufacturer        bool               `json:"isManufacturer"`
	ContractNumber        string             `json:"contractNumber"`
	ContractStartDate     time.Time          `json:"contractStartDate"`
	ContractEndDate       time.Time          `json:"contractEndDate"`
	MaxDebt               float64            `json:"maxDebt"`
	MaxCurrency           string             `json:"maxCurrency"`
	DeferSaleMonths       int32              `json:"deferSaleMonths"`
	DeferSaleType         string             `json:"deferSaleType"`
	DeferPurchaseMonths   int32              `json:"deferPurchaseMonths"`
	DeferPurchaseType     string             `json:"deferPurchaseType"`
	SalePaymentTermId     uint               `json:"salePaymentTermId"`
	PurchasePaymentTermId uint               `json:"purchasePaymentTermId"`
	PriceField            string             `json:"priceField"`
	CustomFields          json.RawMessage    `json:"customFields"`
	TaxExempt             *bool              `json:"taxExempt"`
	TaxProfileRequest     *TaxProfileRequest `json:"taxProfileRequest,omitempty"`
	CustomerLanguage      string             `json:"customerLanguage"`
	Source                string             `json:"source"`
	IsVerified            *bool              `json:"isVerified"`
}

type CompanyPatchRequest struct {
	Name                  *string          `json:"name,omitempty"`
	RegistrationNumber    *string          `json:"registrationNumber,omitempty"`
	TaxNumber             *string          `json:"taxNumber,omitempty"`
	IsPerson              *bool            `json:"isPerson,omitempty"`
	AddressLine1          *string          `json:"addressLine1,omitempty"`
	AddressLine2          *string          `json:"addressLine2,omitempty"`
	AddressLine3          *string          `json:"addressLine3,omitempty"`
	PostCode              *string          `json:"postCode,omitempty"`
	City                  *string          `json:"city,omitempty"`
	CountryCode           *string          `json:"countryCode,omitempty"`
	PhoneNumber           *string          `json:"phoneNumber,omitempty"`
	WarnedDate            *time.Time       `json:"warnedDate,omitempty"`
	DeferSaleDays         *int32           `json:"deferSaleDays,omitempty"`
	DeferPurchaseDays     *int32           `json:"deferPurchaseDays,omitempty"`
	Region                *string          `json:"region,omitempty"`
	Representative        *string          `json:"representative,omitempty"`
	Fax                   *string          `json:"fax,omitempty"`
	Website               *string          `json:"website,omitempty"`
	Email                 *string          `json:"email,omitempty"`
	IsSupplier            *bool            `json:"isSupplier,omitempty"`
	IsBuyer               *bool            `json:"isBuyer,omitempty"`
	IsBranch              *bool            `json:"isBranch,omitempty"`
	IsManufacturer        *bool            `json:"isManufacturer,omitempty"`
	ContractNumber        *string          `json:"contractNumber,omitempty"`
	ContractStartDate     *time.Time       `json:"contractStartDate,omitempty"`
	ContractEndDate       *time.Time       `json:"contractEndDate,omitempty"`
	MaxDebt               *float64         `json:"maxDebt,omitempty"`
	MaxCurrency           *string          `json:"maxCurrency,omitempty"`
	DeferSaleMonths       *int32           `json:"deferSaleMonths,omitempty"`
	DeferSaleType         *string          `json:"deferSaleType,omitempty"`
	DeferPurchaseMonths   *int32           `json:"deferPurchaseMonths,omitempty"`
	DeferPurchaseType     *string          `json:"deferPurchaseType,omitempty"`
	SalePaymentTermId     *uint            `json:"salePaymentTermId,omitempty"`
	PurchasePaymentTermId *uint            `json:"purchasePaymentTermId,omitempty"`
	PriceField            *string          `json:"priceField,omitempty"`
	CustomFields          *json.RawMessage `json:"customFields,omitempty"`
	TaxExempt             *bool            `json:"taxExempt,omitempty"`
	CustomerLanguage      *string          `json:"customerLanguage,omitempty"`
	Source                *string          `json:"source,omitempty"`
	IsVerified            *bool            `json:"isVerified,omitempty"`
}

type CompaniesRequest struct {
	PerPage       int          `json:"perPage,omitempty"`
	Page          int          `json:"page,omitempty"`
	SearchQuery   string       `json:"searchQuery,omitempty"`
	CompanyIds    []uint       `json:"companyIds,omitempty"`
	CreatedAtFrom time.Time    `json:"createdAtFrom,omitempty"`
	CreatedAtTo   time.Time    `json:"createdAtTo,omitempty"`
	UpdatedAtFrom time.Time    `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   time.Time    `json:"updatedAtTo,omitempty"`
	Filters       []ItemFilter `json:"filters,omitempty"`
}

type CompaniesResponse struct {
	Companies      []Company `json:"companies"`
	CompaniesCount int64     `json:"companiesCount"`
}

type CompanyMutateResponse struct {
	CompanyId uint `json:"companyId"`
}

// InvoicesRequest is the preferred plugin invoice list request. It follows
// the newer external API filters more closely than the legacy InvoicesRequest.
type InvoicesRequest struct {
	UnlimitPerPage          bool            `json:"unlimitPerPage,omitempty"`
	PerPage                 int             `json:"perPage,omitempty"`
	Page                    int             `json:"page,omitempty"`
	DateFrom                time.Time       `json:"dateFrom,omitempty"`
	DateTo                  time.Time       `json:"dateTo,omitempty"`
	CreatedAtFrom           time.Time       `json:"createdAtFrom,omitempty"`
	CreatedAtTo             time.Time       `json:"createdAtTo,omitempty"`
	UpdatedAtFrom           time.Time       `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo             time.Time       `json:"updatedAtTo,omitempty"`
	SearchQuery             string          `json:"searchQuery,omitempty"`
	Token                   string          `json:"token,omitempty"`
	ListKind                InvoiceListKind `json:"listKind,omitempty"`
	SelectOnlySales         bool            `json:"selectOnlySales,omitempty"`
	SelectOnlyPurchases     bool            `json:"selectOnlyPurchases,omitempty"`
	SelectOnlyWarehouseMove bool            `json:"selectOnlyWarehouseMove,omitempty"`
	SelectOnlyDebtInvoices  bool            `json:"selectOnlyDebtInvoices,omitempty"`
	SelectSalesAndPurchases bool            `json:"selectSalesAndPurchases,omitempty"`
	DivisionIds             []uint          `json:"divisionIds,omitempty"`
	FromWarehouses          []string        `json:"fromWarehouses,omitempty"`
	ToWarehouses            []string        `json:"toWarehouses,omitempty"`
	Warehouses              []string        `json:"warehouses,omitempty"`
	Operations              []string        `json:"operations,omitempty"`
	OperationIds            []uint          `json:"operationIds,omitempty"`
	DocSeries               []string        `json:"docSeries,omitempty"`
	InvoiceIds              []uint          `json:"invoiceIds,omitempty"`
	CompanyIds              []uint          `json:"companyIds,omitempty"`
	ExcludeCompanyIds       []uint          `json:"excludeCompanyIds,omitempty"`
	OrderWay                string          `json:"orderWay,omitempty"`
	OrderBy                 string          `json:"orderBy,omitempty"`
	Statuses                []string        `json:"statuses,omitempty"`
	CreatedUserIDs          []uint          `json:"createdUserIds,omitempty"`
}

type InvoicesListResponse struct {
	Invoices   []Invoice `json:"invoices"`
	TotalCount int64     `json:"totalCount"`
}

type InvoiceType string

const (
	InvoiceTypeSale     InvoiceType = "sale"
	InvoiceTypePurchase InvoiceType = "purchase"
)

type InvoiceAddress struct {
	AddressId    uint   `json:"addressId,omitempty"`
	Name         string `json:"name,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	Region       string `json:"region,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	PostCode     string `json:"postCode,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

type InvoiceInstallment struct {
	InstallmentNo int       `json:"installmentNo"`
	DueDate       time.Time `json:"dueDate"`
	Amount        float64   `json:"amount"`
	PaidAmount    float64   `json:"paidAmount"`
	DebtAmount    float64   `json:"debtAmount"`
	Percentage    float64   `json:"percentage"`
}

// Invoice mirrors the newer external invoice list shape. Prefer this
// over InvoiceFull for new plugins.
type Invoice struct {
	InvoiceId         int32                `json:"invoiceId"`
	Token             string               `json:"token"`
	Type              InvoiceType          `json:"type"`
	Operation         string               `json:"operation"`
	OperationId       int32                `json:"operationId"`
	CompanyName       string               `json:"companyName"`
	CompanyId         int32                `json:"companyId"`
	InvoicePrefix     string               `json:"invoicePrefix"`
	InvoiceNumber     string               `json:"invoiceNumber"`
	InvoiceNumber2    string               `json:"invoiceNumber2"`
	InternalNumber    string               `json:"internalNumber"`
	Date              time.Time            `json:"date"`
	TotalTax          float64              `json:"totalTax"`
	TotalExclTax      float64              `json:"totalExclTax"`
	TotalInclTax      float64              `json:"totalInclTax"`
	Currency          string               `json:"currency"`
	IsInvoicePaid     bool                 `json:"isInvoicePaid"`
	TotalDebt         float64              `json:"totalDebt"`
	TotalPaid         float64              `json:"totalPaid"`
	PaidLastDate      time.Time            `json:"paidLastDate"`
	DueDate           time.Time            `json:"dueDate"`
	Installments      []InvoiceInstallment `json:"installments,omitempty"`
	ReceiveDate       time.Time            `json:"receiveDate"`
	TotalCost         float64              `json:"totalCost"`
	TotalMargin       float64              `json:"totalMargin"`
	FromWarehouse     string               `json:"fromWarehouse"`
	ToWarehouse       string               `json:"toWarehouse"`
	Comment           string               `json:"comment"`
	TotalReverseVat   float64              `json:"totalReverseVat"`
	TotalCompensation float64              `json:"totalCompensation"`
	CreatedAt         time.Time            `json:"createdAt"`
	UpdatedAt         time.Time            `json:"updatedAt"`
	CreatedUserInfo   UserInfo             `json:"createdUserInfo"`
	UpdatedUserInfo   UserInfo             `json:"updatedUserInfo"`
	IsVerified        bool                 `json:"isVerified"`
	PdfDocumentId     uint                 `json:"pdfDocumentId"`
	Status            string               `json:"status"`
	CustomFields      json.RawMessage      `json:"customFields"`
	CustomerReference string               `json:"customerReference"`
	DeliveryContactId uint                 `json:"deliveryContactId"`
	OrdererContactId  uint                 `json:"ordererContactId"`
	BillingAddress    InvoiceAddress       `json:"billingAddress"`
	ShippingAddress   InvoiceAddress       `json:"shippingAddress"`
}

type ProductMediaFile struct {
	MediaFileId  uint   `json:"mediaFileId"`
	FileUrl      string `json:"fileUrl"`      // CDN/public URL
	MediaType    string `json:"mediaType"`    // image, video, pdf, etc.
	MimeType     string `json:"mimeType"`     // image/jpeg, video/mp4, etc.
	OriginalName string `json:"originalName"` // original file name
	Hash         string `json:"hash"`         // (optional) for deduplication

	// extended from assignments
	OwnerType string `json:"ownerType"` // e.g., "product_merge", "product_variant", "brand", "category"
	OwnerId   uint   `json:"ownerId"`   // e.g., product_id, merge_id, brand_id
	SortOrder int    `json:"sortOrder"` // sort order for display
	IsPrimary bool   `json:"isPrimary"` // primary image indicator

	Content ProductMediaContent `json:"content"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type ProductMediaContent struct {
	AltText       string `json:"altText"`
	LangShortCode string `json:"langShortCode"`
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

type OrderAddress struct {
	AddressId    uint   `json:"addressId,omitempty"`
	Name         string `json:"name,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	CompanyName  string `json:"companyName,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	PostCode     string `json:"postCode,omitempty"`
	Phone        string `json:"phone,omitempty"`
	MobilePhone  string `json:"mobilePhone,omitempty"` // Deprecated: use phone.
}

type OrderPickupLocation struct {
	LocationId          int32  `json:"locationId,omitempty"`
	LocationName        string `json:"locationName,omitempty"`
	LocationCity        string `json:"locationCity,omitempty"`
	LocationPostCode    string `json:"locationPostCode,omitempty"`
	LocationCompanyCode string `json:"locationCompanyCode,omitempty"`
}

type DocumentChargeCreateRequest struct {
	ChargeTypeID    uint     `json:"chargeTypeId" binding:"required"`
	CalculationType string   `json:"calculationType" binding:"required,oneof=percentage fixedAmount perUnit"`
	Percentage      *float64 `json:"percentage"`
	Amount          *float64 `json:"amount"`
	Quantity        *float64 `json:"quantity"`
	IsTaxable       *bool    `json:"isTaxable"`
	TaxPercentage   *float64 `json:"taxPercentage"`
	TaxAmount       *float64 `json:"taxAmount"`
	Remarks         string   `json:"remarks"`
}

type LineChargeCreateRequest struct {
	ChargeTypeID    uint     `json:"chargeTypeId" binding:"required"`
	CalculationType string   `json:"calculationType" binding:"required,oneof=percentage fixedAmount perUnit"`
	Percentage      *float64 `json:"percentage"`
	Amount          *float64 `json:"amount"`
	IsTaxable       *bool    `json:"isTaxable"`
	TaxPercentage   *float64 `json:"taxPercentage"`
	TaxAmount       *float64 `json:"taxAmount"`
	Remarks         string   `json:"remarks"`
}

type OrderLineRequest struct {
	LineType              uint                      `json:"lineType,omitempty"`
	Title                 string                    `json:"title,omitempty"`
	Description           string                    `json:"description,omitempty"`
	ProductId             uint                      `json:"productId,omitempty"`
	Quantity              float64                   `json:"quantity"`
	MeasurementUnit       string                    `json:"measurementUnit,omitempty"`
	PriceExclTax          float64                   `json:"priceExclTax"`
	TaxPercentage         float64                   `json:"taxPercentage,omitempty"`
	TaxCode               string                    `json:"taxCode,omitempty"`
	TaxAmount             *float64                  `json:"taxAmount,omitempty"`
	DiscountPerUnit       float64                   `json:"discountPerUnit,omitempty"`
	DiscountPercent       float64                   `json:"discountPercent,omitempty"`
	PriceIncludesDiscount *bool                     `json:"priceIncludesDiscount,omitempty"`
	ShipByDate            *time.Time                `json:"shipByDate,omitempty"`
	Charges               []LineChargeCreateRequest `json:"charges,omitempty"`
}

type OrderCreateRequest struct {
	Date                 time.Time                     `json:"date"`
	ExpiryDate           time.Time                     `json:"expiryDate,omitempty"`
	CompanyId            int32                         `json:"companyId"`
	Comments             string                        `json:"comments,omitempty"`
	CustomerNote         string                        `json:"customerNote,omitempty"`
	DeliveryInstructions string                        `json:"deliveryInstructions,omitempty"`
	CurrencyId           int32                         `json:"currencyId,omitempty"`
	Currency             string                        `json:"currency,omitempty"`
	PaymentTermId        uint                          `json:"paymentTermId,omitempty"`
	StatusId             uint                          `json:"statusId,omitempty"`
	Title                string                        `json:"title,omitempty"`
	ExternalUserId       uint                          `json:"externalUserId,omitempty"`
	Email                string                        `json:"email,omitempty"`
	ShippingAddress      OrderAddress                  `json:"shippingAddress,omitempty"`
	BillingAddress       OrderAddress                  `json:"billingAddress,omitempty"`
	PickupLocation       OrderPickupLocation           `json:"pickupLocation,omitempty"`
	AddressId            uint                          `json:"addressId,omitempty"` // Deprecated: use shippingAddress.addressId.
	DeliveryContactId    uint                          `json:"deliveryContactId,omitempty"`
	OrdererContactId     uint                          `json:"ordererContactId,omitempty"`
	PaymentMethod        string                        `json:"paymentMethod,omitempty"`
	PaymentProvider      string                        `json:"paymentProvider,omitempty"`
	PaymentIsCod         bool                          `json:"paymentIsCod,omitempty"`
	PaymentReference     string                        `json:"paymentReference,omitempty"`
	ShippingMethod       string                        `json:"shippingMethod,omitempty"`
	ShippingMethodCode   string                        `json:"shippingMethodCode,omitempty"`
	TrackingNumber       string                        `json:"trackingNumber,omitempty"`
	TrackingCarrier      string                        `json:"trackingCarrier,omitempty"`
	CarrierID            *uint                         `json:"carrierId,omitempty"`
	CarrierServiceID     *uint                         `json:"carrierServiceId,omitempty"`
	CarrierCode          string                        `json:"carrierCode,omitempty"`
	CarrierServiceCode   string                        `json:"carrierServiceCode,omitempty"`
	LineItems            []OrderLineRequest            `json:"lineItems"`
	DocumentCharges      []DocumentChargeCreateRequest `json:"documentCharges,omitempty"`
	ChannelID            uint                          `json:"channelId,omitempty"`
	ChannelUUID          string                        `json:"channelUuid,omitempty"`
	ExternalOrderID      string                        `json:"externalOrderId,omitempty"`
	PONumber             string                        `json:"poNumber,omitempty"`
	CustomerReference    string                        `json:"customerReference,omitempty"`
	CustomFields         json.RawMessage               `json:"customFields,omitempty"`
}

type OrderUpdateRequest struct {
	Date                 *time.Time                    `json:"date,omitempty"`
	ExpiryDate           *time.Time                    `json:"expiryDate,omitempty"`
	CompanyId            *int32                        `json:"companyId,omitempty"`
	Comments             *string                       `json:"comments,omitempty"`
	CustomerNote         *string                       `json:"customerNote,omitempty"`
	DeliveryInstructions *string                       `json:"deliveryInstructions,omitempty"`
	CurrencyId           *int32                        `json:"currencyId,omitempty"`
	Currency             *string                       `json:"currency,omitempty"`
	PaymentTermId        *uint                         `json:"paymentTermId,omitempty"`
	StatusId             *uint                         `json:"statusId,omitempty"`
	Title                *string                       `json:"title,omitempty"`
	ExternalUserId       *uint                         `json:"externalUserId,omitempty"`
	Email                *string                       `json:"email,omitempty"`
	ShippingAddress      *OrderAddress                 `json:"shippingAddress,omitempty"`
	BillingAddress       *OrderAddress                 `json:"billingAddress,omitempty"`
	PickupLocation       *OrderPickupLocation          `json:"pickupLocation,omitempty"`
	AddressId            *uint                         `json:"addressId,omitempty"`
	DeliveryContactId    *uint                         `json:"deliveryContactId,omitempty"`
	OrdererContactId     *uint                         `json:"ordererContactId,omitempty"`
	PaymentMethod        *string                       `json:"paymentMethod,omitempty"`
	PaymentProvider      *string                       `json:"paymentProvider,omitempty"`
	PaymentIsCod         *bool                         `json:"paymentIsCod,omitempty"`
	PaymentReference     *string                       `json:"paymentReference,omitempty"`
	ShippingMethod       *string                       `json:"shippingMethod,omitempty"`
	ShippingMethodCode   *string                       `json:"shippingMethodCode,omitempty"`
	TrackingNumber       *string                       `json:"trackingNumber,omitempty"`
	TrackingCarrier      *string                       `json:"trackingCarrier,omitempty"`
	CarrierID            *uint                         `json:"carrierId,omitempty"`
	CarrierServiceID     *uint                         `json:"carrierServiceId,omitempty"`
	CarrierCode          *string                       `json:"carrierCode,omitempty"`
	CarrierServiceCode   *string                       `json:"carrierServiceCode,omitempty"`
	LineItems            []OrderLineRequest            `json:"lineItems,omitempty"`
	DocumentCharges      []DocumentChargeCreateRequest `json:"documentCharges,omitempty"`
	ChannelID            *uint                         `json:"channelId,omitempty"`
	ChannelUUID          *string                       `json:"channelUuid,omitempty"`
	ExternalOrderID      *string                       `json:"externalOrderId,omitempty"`
	PONumber             *string                       `json:"poNumber,omitempty"`
	CustomerReference    *string                       `json:"customerReference,omitempty"`
	CustomFields         json.RawMessage               `json:"customFields,omitempty"`
}

type OrderPatchRequest struct {
	AddressId            *uint                      `json:"addressId,omitempty"`
	ShippingAddress      *OrderPatchShippingAddress `json:"shippingAddress,omitempty"`
	BillingAddress       *OrderPatchBillingAddress  `json:"billingAddress,omitempty"`
	DeliveryContactId    *uint                      `json:"deliveryContactId,omitempty"`
	OrdererContactId     *uint                      `json:"ordererContactId,omitempty"`
	PaymentMethod        *string                    `json:"paymentMethod,omitempty"`
	PaymentProvider      *string                    `json:"paymentProvider,omitempty"`
	PaymentIsCod         *bool                      `json:"paymentIsCod,omitempty"`
	PaymentId            *int32                     `json:"paymentId,omitempty"`
	PaymentReference     *string                    `json:"paymentReference,omitempty"`
	PaymentTermId        *uint                      `json:"paymentTermId,omitempty"`
	PONumber             *string                    `json:"poNumber,omitempty"`
	CustomerReference    *string                    `json:"customerReference,omitempty"`
	ChannelUUID          *string                    `json:"channelUuid,omitempty"`
	CustomerNote         *string                    `json:"customerNote,omitempty"`
	DeliveryInstructions *string                    `json:"deliveryInstructions,omitempty"`
	CustomFields         map[string]any             `json:"customFields,omitempty"`
}

type OrderPatchShippingAddress struct {
	AddressId   *uint   `json:"addressId,omitempty"`
	Name        *string `json:"name,omitempty"`
	CompanyName *string `json:"companyName,omitempty"` // Deprecated: use name.
}

type OrderPatchBillingAddress struct {
	AddressId *uint `json:"addressId,omitempty"`
}

type OrderStatus struct {
	StatusId                 uint   `json:"statusId"`
	Code                     string `json:"code"`
	Name                     string `json:"name"`
	FormInvoiceOnChange      bool   `json:"formInvoiceOnChange"`
	IsInvoiceDownloadable    bool   `json:"isInvoiceDownloadable"`
	ShowAsPaid               bool   `json:"showAsPaid"`
	IsWaitingForPayment      bool   `json:"isWaitingForPayment"`
	IsCancelled              bool   `json:"isCancelled"`
	IsOrderCompleted         bool   `json:"isOrderCompleted"`
	IsActive                 bool   `json:"isActive"`
	IncludeInOrderedQuantity bool   `json:"includeInOrderedQuantity"`
}

type Order struct {
	OrderId              uint            `json:"orderId,omitempty"`
	QuoteId              uint            `json:"quoteId,omitempty"`
	DocumentType         string          `json:"documentType"`
	Date                 time.Time       `json:"date"`
	CreatedAt            time.Time       `json:"createdAt"`
	CompanyId            int32           `json:"companyId"`
	CompanyName          string          `json:"companyName"`
	Status               OrderStatus     `json:"status"`
	Title                string          `json:"title"`
	Document             string          `json:"document"`
	IsDeleted            bool            `json:"isDeleted"`
	Comments             string          `json:"comments"`
	CustomerNote         string          `json:"customerNote"`
	DeliveryInstructions string          `json:"deliveryInstructions"`
	ShippingAddress      OrderAddress    `json:"shippingAddress"`
	BillingAddress       OrderAddress    `json:"billingAddress"`
	AddressId            uint            `json:"addressId,omitempty"` // Deprecated: use shippingAddress.addressId.
	DeliveryContactId    uint            `json:"deliveryContactId,omitempty"`
	OrdererContactId     uint            `json:"ordererContactId,omitempty"`
	PaymentMethod        string          `json:"paymentMethod"`
	PaymentProvider      string          `json:"paymentProvider,omitempty"`
	PaymentIsCod         bool            `json:"paymentIsCod"`
	PaymentReference     string          `json:"paymentReference,omitempty"`
	PaymentTermId        uint            `json:"paymentTermId,omitempty"`
	ShippingMethodCode   string          `json:"shippingMethodCode,omitempty"`
	Email                string          `json:"email"`
	ExternalUserId       uint            `json:"externalUserId"`
	CreatedUserId        uint            `json:"createdUserId"`
	CreatedUserInfo      UserInfo        `json:"createdUserInfo"`
	UpdatedUserInfo      UserInfo        `json:"updatedUserInfo"`
	CustomFields         json.RawMessage `json:"customFields,omitempty"`
	Currency             string          `json:"currency"`
	TotalExclTax         float64         `json:"totalExclTax"`
	TotalTax             float64         `json:"totalTax"`
	TotalInclTax         float64         `json:"totalInclTax"`
	ChannelID            uint            `json:"channelId,omitempty"`
	ChannelUUID          string          `json:"channelUuid,omitempty"`
	ExternalOrderID      string          `json:"externalOrderId,omitempty"`
	PONumber             string          `json:"poNumber,omitempty"`
	CustomerReference    string          `json:"customerReference,omitempty"`
}

type OrderLine struct {
	ID                    int32      `json:"id"`
	LineType              uint       `json:"lineType"`
	Title                 string     `json:"title"`
	Description           string     `json:"description,omitempty"`
	ProductId             uint       `json:"productId,omitempty"`
	ProductName           string     `json:"productName,omitempty"`
	Article               string     `json:"article,omitempty"`
	Quantity              float64    `json:"quantity"`
	MeasurementUnit       string     `json:"measurementUnit,omitempty"`
	Currency              string     `json:"currency"`
	PriceExclTax          float64    `json:"priceExclTax"`
	PriceInclTax          float64    `json:"priceInclTax"`
	OriginalPrice         float64    `json:"originalPrice,omitempty"`
	TaxPercentage         float64    `json:"taxPercentage"`
	TaxCode               string     `json:"taxCode,omitempty"`
	TaxAmount             *float64   `json:"taxAmount,omitempty"`
	DiscountPerUnit       float64    `json:"discountPerUnit"`
	DiscountPercent       float64    `json:"discountPercent"`
	PriceIncludesDiscount bool       `json:"priceIncludesDiscount"`
	ShipByDate            *time.Time `json:"shipByDate,omitempty"`
	TotalExclTax          float64    `json:"totalExclTax"`
	TotalInclTax          float64    `json:"totalInclTax"`
}

type OrderDetail struct {
	Order
	LineItems       []OrderLine         `json:"lineItems"`
	ExpiryDate      time.Time           `json:"expiryDate"`
	CurrencyId      int32               `json:"currencyId"`
	PickupLocation  OrderPickupLocation `json:"pickupLocation"`
	ShippingMethod  string              `json:"shippingMethod"`
	TrackingNumber  string              `json:"trackingNumber,omitempty"`
	TrackingCarrier string              `json:"trackingCarrier,omitempty"`
	Token           string              `json:"token"`
	DocumentCharges []map[string]any    `json:"documentCharges,omitempty"`
}

type OrdersRequest struct {
	PerPage                      int       `json:"perPage"`
	Page                         int       `json:"page"`
	DateFrom                     time.Time `json:"dateFrom"`
	DateTo                       time.Time `json:"dateTo"`
	SearchQuery                  string    `json:"searchQuery"`
	SelectOnlyQuotes             bool      `json:"selectOnlyQuotes"`
	SelectOnlyOrders             bool      `json:"selectOnlyOrders"`
	SelectOnlyProformas          bool      `json:"selectOnlyProformas"`
	SelectOnlyCompleted          bool      `json:"selectOnlyCompleted"`
	SelectOnlyUncompleted        bool      `json:"selectOnlyUncompleted"`
	SelectIncludingDeletedOrders bool      `json:"selectIncludingDeletedOrders"`
	SelectOnlyDeletedOrders      bool      `json:"selectOnlyDeletedOrders"`
	CompanyId                    uint      `json:"companyId"`
	ExternalUserId               uint      `json:"externalUserId"`
	StatusIds                    []uint    `json:"statusIds"`
	StatusCodes                  []string  `json:"statusCodes"`
	StatusLetters                []string  `json:"statusLetters"`
	Token                        string    `json:"token"`
	QueryType                    string    `json:"queryType"`
	OrderIds                     []uint    `json:"orderIds"`
}

type OrdersResponse struct {
	Orders     []Order `json:"orders"`
	TotalCount int64   `json:"totalCount"`
}

type OrderMutateResponse struct {
	OrderID      uint   `json:"orderId"`
	DocumentType string `json:"documentType"`
}

type OrderInvoiceCreateRequest struct {
	IsPurchase        bool                      `json:"isPurchase,omitempty"`
	FromWarehouse     string                    `json:"fromWarehouse,omitempty"`
	PdfDocumentId     uint                      `json:"pdfDocumentId,omitempty"`
	LineItems         []OrderInvoiceLineRequest `json:"lineItems,omitempty"`
	ChargePolicy      string                    `json:"chargePolicy,omitempty"`
	DocumentChargeIds []uint                    `json:"documentChargeIds,omitempty"`
}

type OrderInvoiceLineRequest struct {
	OrderLineId int32   `json:"orderLineId"`
	Quantity    float64 `json:"quantity"`
}

type OrderInvoiceCreateResponse struct {
	InvoiceId uint `json:"invoiceId"`
}

type OrderPaymentStatus string

const (
	OrderPaymentStatusPending                OrderPaymentStatus = "pending"
	OrderPaymentStatusAuthorized             OrderPaymentStatus = "authorized"
	OrderPaymentStatusSubmittedForSettlement OrderPaymentStatus = "submittedForSettlement"
	OrderPaymentStatusSettled                OrderPaymentStatus = "settled"
	OrderPaymentStatusVoided                 OrderPaymentStatus = "voided"
	OrderPaymentStatusFailed                 OrderPaymentStatus = "failed"
	OrderPaymentStatusRefunded               OrderPaymentStatus = "refunded"
)

type OrderPaymentCreateRequest struct {
	Provider              string             `json:"provider"`
	ProviderTransactionID *string            `json:"providerTransactionId,omitempty"`
	TransactionDate       *time.Time         `json:"transactionDate,omitempty"`
	Amount                float64            `json:"amount"`
	Currency              string             `json:"currency"`
	Status                OrderPaymentStatus `json:"status,omitempty"`
}

type OrderPaymentUpdateRequest struct {
	Provider              *string             `json:"provider,omitempty"`
	ProviderTransactionID *string             `json:"providerTransactionId,omitempty"`
	TransactionDate       *time.Time          `json:"transactionDate,omitempty"`
	Amount                *float64            `json:"amount,omitempty"`
	Currency              *string             `json:"currency,omitempty"`
	Status                *OrderPaymentStatus `json:"status,omitempty"`
}

type OrderPayment struct {
	PaymentID             uint               `json:"paymentId"`
	OrderID               uint               `json:"orderId"`
	Provider              string             `json:"provider"`
	ProviderTransactionID *string            `json:"providerTransactionId,omitempty"`
	TransactionDate       *time.Time         `json:"transactionDate,omitempty"`
	Amount                float64            `json:"amount"`
	Currency              string             `json:"currency"`
	Status                OrderPaymentStatus `json:"status"`
	CreatedAt             time.Time          `json:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt"`
}

type OrderPaymentsRequest struct {
	OrderId             uint      `json:"orderId"`
	Provider            string    `json:"provider"`
	Status              string    `json:"status"`
	Page                int       `json:"page"`
	PerPage             int       `json:"perPage"`
	TransactionDateFrom time.Time `json:"transactionDateFrom"`
	TransactionDateTo   time.Time `json:"transactionDateTo"`
}

type OrderPaymentsResponse struct {
	Payments   []OrderPayment `json:"payments"`
	TotalCount int64          `json:"totalCount"`
}

type OrderPaymentMutateResponse struct {
	PaymentID uint `json:"paymentId"`
}

type Product struct {
	ProductId                uint                     `json:"productId"`
	Name                     string                   `json:"name"`
	Price                    float64                  `json:"price"`
	Currency                 string                   `json:"currency"`
	Price1                   float64                  `json:"price1"`
	Currency1                string                   `json:"currency1"`
	Price2                   float64                  `json:"price2"`
	Currency2                string                   `json:"currency2"`
	Price3                   float64                  `json:"price3"`
	Currency3                string                   `json:"currency3"`
	Price4                   float64                  `json:"price4"`
	Currency4                string                   `json:"currency4"`
	Price5                   float64                  `json:"price5"`
	Currency5                string                   `json:"currency5"`
	Price6                   float64                  `json:"price6"`
	Currency6                string                   `json:"currency6"`
	Price7                   float64                  `json:"price7"`
	Currency7                string                   `json:"currency7"`
	Price8                   float64                  `json:"price8"`
	Currency8                string                   `json:"currency8"`
	Price9                   float64                  `json:"price9"`
	Currency9                string                   `json:"currency9"`
	Price10                  float64                  `json:"price10"`
	Currency10               string                   `json:"currency10"`
	MeasurementUnit          string                   `json:"measurementUnit"`
	Article                  string                   `json:"article"`
	Description              string                   `json:"description"`
	Countable                bool                     `json:"countable"`
	OrderedQuantity          float64                  `json:"orderedQuantity"`
	SupplierCode1            string                   `json:"supplierCode1"`
	SupplierCode2            string                   `json:"supplierCode2"`
	SupplierCode3            string                   `json:"supplierCode3"`
	Brand                    string                   `json:"brand"`
	BrandId                  uint                     `json:"brandId"`
	IsDiscounted             bool                     `json:"isDiscounted"`
	IsInPriceList            bool                     `json:"isInPriceList"`
	Importer                 string                   `json:"importer"`
	WarrantyInMonths         int32                    `json:"warrantyInMonths"`
	IsRecommended            bool                     `json:"isRecommended"`
	IsRemoved                bool                     `json:"isRemoved"`
	Comment                  string                   `json:"comment"`
	Slug                     string                   `json:"slug"`
	MetaTitle                string                   `json:"metaTitle"`
	MetaDescription          string                   `json:"metaDescription"`
	MetaKeywords             string                   `json:"metaKeywords"`
	SearchKeywords           string                   `json:"searchKeywords"`
	CanonicalSlug            string                   `json:"canonicalSlug"`
	DiscountValue            float64                  `json:"discountValue"`
	DiscountType             int32                    `json:"discountType"`
	DiscountPercent          float64                  `json:"discountPercent"`
	DiscountedPrice          float64                  `json:"discountedPrice"`
	DiscountedPriceWoVat     float64                  `json:"discountedPriceWoVat"`
	SortOrder                uint                     `json:"sortOrder"`
	MainCategoryId           int32                    `json:"mainCategoryId"`
	Quantity                 int64                    `json:"quantity"`
	Weight                   float64                  `json:"weight"`
	Volume                   float64                  `json:"volume"`
	Barcode                  string                   `json:"barcode"`
	VatPercentage            float64                  `json:"vatPercentage"`
	OriginCountry            string                   `json:"originCountry"`
	CategoryIds              []int32                  `json:"categoryIds"`
	MetaValues               map[string]string        `json:"metaValues"`
	AccountingNumbers        ProductAccountingNumbers `json:"accountingNumbers"`
	Length                   float32                  `json:"length"`
	Width                    float32                  `json:"width"`
	Height                   float32                  `json:"height"`
	IsPhysicalOnlyProduct    bool                     `json:"isPhysicalOnlyProduct"`
	AvailableForOrderInEshop bool                     `json:"availableForOrderInEshop"`
	ShowInEshop              bool                     `json:"showInEshop"`
	EshopOutOfStockBehaviour string                   `json:"eshopOutOfStockBehaviour"`
	SalesMinQuantity         float64                  `json:"salesMinQuantity"`
	SalesMultiple            float64                  `json:"salesMultiple"`
	IsDigital                bool                     `json:"isDigital"`
	CategoryNames            []string                 `json:"categoryNames"`
	MediaItems               []ProductMediaFile       `json:"mediaItems"`
	CreatedAt                time.Time                `json:"createdAt"`
	UpdatedAt                time.Time                `json:"updatedAt"`
	CreatedUserInfo          UserInfo                 `json:"createdUserInfo"`
	UpdatedUserInfo          UserInfo                 `json:"updatedUserInfo"`
	CustomFields             json.RawMessage          `json:"customFields"`
}

type ProductCreateEdit struct {
	Product
	MediaItems         []ProductImage    `json:"mediaItems"`
	DeleteMediaFileIds []uint            `json:"deleteMediaFileIds"`
	MediaReorder       *[]MediaItemOrder `json:"mediaReorder,omitempty"`
	AssignMediaFileIds []uint            `json:"assignMediaFileIds"`
}

type MediaItemOrder struct {
	MediaFileId uint `json:"mediaFileId"`
	SortOrder   int  `json:"sortOrder"`
}

type ProductCreateEditRequest struct {
	ProductCreateEdit
}

type ProductPatchRequest struct {
	Product    ProductCreateEdit `json:"product"`
	UpdateMask []string          `json:"updateMask"`
}

type ProductsRequest struct {
	SelectInEshop          bool       `json:"selectInEshop"`
	SelectInWarehouse      bool       `json:"selectInWarehouse"`
	Warehouses             []string   `json:"warehouses"`
	SelectUncountable      bool       `json:"selectUncountable"`
	SelectArchived         bool       `json:"selectArchived"`
	CategoryUri            string     `json:"categoryUri"`
	CategoryIds            []uint     `json:"categoryIds"`
	PerPage                int        `json:"perPage"`
	Page                   int        `json:"page"`
	Brands                 []string   `json:"brands"`
	PriceRange             PriceRange `json:"priceRange"`
	OrderWay               string     `json:"orderWay"`
	OrderBy                string     `json:"orderBy"`
	SearchQuery            string     `json:"searchQuery"`
	ProductIds             []uint     `json:"productIds"`
	SkipQuantityColumn     bool       `json:"skipQuantityColumn"`
	UnlimitedProducts      bool       `json:"unlimitedProducts"`
	UseIndexedSearch       bool       `json:"useIndexedSearch"`
	SelectMerged           bool       `json:"selectMerged"`
	ExpandVariations       bool       `json:"expandVariations"`
	CreatedAtFrom          time.Time  `json:"createdAtFrom"`
	CreatedAtTo            time.Time  `json:"createdAtTo"`
	UpdatedAtFrom          time.Time  `json:"updatedAtFrom"`
	UpdatedAtTo            time.Time  `json:"updatedAtTo"`
	CreatedOrUpdatedAtFrom time.Time  `json:"createdOrUpdatedAtFrom"`
	CreatedOrUpdatedAtTo   time.Time  `json:"createdOrUpdatedAtTo"`
}

type ProductsResponse struct {
	Products            []Product  `json:"products"`
	ProductsCount       int64      `json:"productsCount"`
	AvailablePriceRange PriceRange `json:"availablePriceRange"`
	Brands              []string   `json:"brands"`
}

type ProductMutateResponse struct {
	ProductId uint `json:"productId"`
}

type ProductImage struct {
	FileName string `json:"fileName"`
	FileBlob []byte `json:"fileBlob"`
}

type StockTotal struct {
	TotalSum      float64 `json:"totalSum"`
	TotalCurrency string  `json:"totalCurrency"`
	TotalQuantity float64 `json:"totalQuantity"`
	Warehouse     string  `json:"warehouse"`
}

// ProductStock mirrors the external product-stock response. It uses
// Product for productInfo so stock-aware plugins can work with the same
// product names used by the external API.
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

type ProductStocksRequest struct {
	Warehouses       []string     `json:"warehouses,omitempty"`
	ProductIds       []uint       `json:"productIds,omitempty"`
	PurchaseDateFrom time.Time    `json:"dateFrom,omitempty"`
	PurchaseDateTo   time.Time    `json:"dateTo,omitempty"`
	SearchQuery      string       `json:"searchQuery,omitempty"`
	Filters          []ItemFilter `json:"filters,omitempty"`
	GroupBy          []string     `json:"groupBy,omitempty"`
	PerPage          int          `json:"perPage,omitempty"`
	Page             int          `json:"page,omitempty"`
}

type ProductStocksResponse struct {
	Stocks             []ProductStock         `json:"stocks"`
	TotalQuantity      float64                `json:"totalQuantity"`
	TotalSum           float64                `json:"totalSum"`
	TotalCurrency      string                 `json:"totalCurrency"`
	TotalsByWarehouses map[string]*StockTotal `json:"totalsByWarehouses"`
	TotalCount         int64                  `json:"totalCount"`
}

type WarehouseProductQuantity struct {
	Warehouse         string  `json:"warehouse"`
	WarehouseFullname string  `json:"warehouseFullname"`
	WarehouseId       uint    `json:"warehouseId"`
	Quantity          float64 `json:"quantity"`
}

type ProductWarehousesQuantitiesResponse struct {
	Quantities []WarehouseProductQuantity `json:"quantities"`
}

type StockAdjustmentReason string

const (
	StockAdjustmentReasonInitialBalance StockAdjustmentReason = "initialBalance"
	StockAdjustmentReasonCycleCount     StockAdjustmentReason = "cycleCount"
	StockAdjustmentReasonDamage         StockAdjustmentReason = "damage"
	StockAdjustmentReasonShrinkage      StockAdjustmentReason = "shrinkage"
	StockAdjustmentReasonCorrection     StockAdjustmentReason = "correction"
	StockAdjustmentReasonReceived       StockAdjustmentReason = "received"
	StockAdjustmentReasonReturned       StockAdjustmentReason = "returned"
)

type ProductStockSetRequest struct {
	Stocks []ProductStockSetEntry `json:"stocks"`
	Date   *time.Time             `json:"date,omitempty"`
	Memo   string                 `json:"memo,omitempty"`
	Reason StockAdjustmentReason  `json:"reason,omitempty"`
}

type ProductStockSetEntry struct {
	ProductId   uint    `json:"productId"`
	WarehouseId int32   `json:"warehouseId"`
	Quantity    float64 `json:"quantity"`
	CostPrice   float64 `json:"costPrice"`
}

type ProductStockAdjustRequest struct {
	Adjustments []ProductStockAdjustEntry `json:"adjustments"`
	Date        *time.Time                `json:"date,omitempty"`
	Memo        string                    `json:"memo,omitempty"`
	Reason      StockAdjustmentReason     `json:"reason,omitempty"`
}

type ProductStockAdjustEntry struct {
	ProductId        uint    `json:"productId"`
	WarehouseId      int32   `json:"warehouseId"`
	AdjustQuantityBy float64 `json:"adjustQuantityBy"`
	CostPrice        float64 `json:"costPrice,omitempty"`
}

type ProductStockAdjustmentResponse struct {
	CreatedInvoices []StockAdjustmentInvoiceInfo `json:"createdInvoices"`
	TotalAdjusted   int                          `json:"totalAdjusted"`
	Errors          []StockAdjustmentError       `json:"errors,omitempty"`
}

type StockAdjustmentInvoiceInfo struct {
	InvoiceId     int32  `json:"invoiceId"`
	WarehouseId   int32  `json:"warehouseId"`
	WarehouseCode string `json:"warehouseCode"`
	ProductCount  int    `json:"productCount"`
}

type StockAdjustmentError struct {
	ProductId   uint   `json:"productId"`
	WarehouseId int32  `json:"warehouseId"`
	Error       string `json:"error"`
}
