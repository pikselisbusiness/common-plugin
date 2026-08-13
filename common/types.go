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

// ---------------------------------------------------------------------------
// Product feature sync. Concrete/flat types (no interface{}). Org/user come from
// the RequestContext. Shared by the v1 and v2 plugin APIs.
// ---------------------------------------------------------------------------

// UpsertFeatureRequest idempotently creates or updates a feature (attribute
// definition) by name. Type is the raw external type; the host normalizes it.
// ExternalId is the source system's attribute id (reserved for a sync mapping).
type UpsertFeatureRequest struct {
	Name          string
	Type          string
	Unit          string
	UnitDimension string
	ExternalId    string
}

type UpsertFeatureResponse struct {
	FeatureId uint
}

// FeatureValueItem is one desired value for a product+feature. NumericValue is
// nullable; ColorCode applies to color features only.
type FeatureValueItem struct {
	Name         string
	NumericValue *float64
	Unit         string
	ColorCode    string
}

// ProductFeatureSet is the full desired value set for one feature on one product
// (replace semantics: the product's assignments for FeatureId become exactly Values).
type ProductFeatureSet struct {
	FeatureId uint
	Values    []FeatureValueItem
}

// ProductFeatureSyncItem carries one product (matched by SKU) and the features to
// replace on it.
type ProductFeatureSyncItem struct {
	Sku      string
	Features []ProductFeatureSet
}

// ReplaceProductFeaturesRequest is one page of products. MatchColumn is the
// whitelisted identifier column the SKU matches against (empty -> host default).
//
// ReplaceScope selects the delete scope of the replace:
//   - "" or "features": delete assignments only for the features present in Items
//     (per-feature replace; safe for partial/chunked writes).
//   - "source": first delete ALL assignments of features belonging to Source on the
//     matched products, then insert Items. This makes the request declarative for
//     the whole source ("the product's <Source> state is exactly this") and
//     self-heals stale assignments from renamed/removed source attributes. A chunked
//     caller must send it on the FIRST chunk only, else later chunks would wipe
//     earlier chunks' writes.
type ReplaceProductFeaturesRequest struct {
	Items        []ProductFeatureSyncItem
	DryRun       bool
	MatchColumn  string
	ReplaceScope string
	Source       string
}

type ReplaceProductFeaturesResult struct {
	MatchedProducts   int
	UnmatchedSkus     []string
	AssignedCount     int64
	DeassignedCount   int64
	CreatedValueCount int64
}

// PruneProductFeaturesRequest asks the host to delete features created by a source
// integration whose external key is no longer present. KeepExternalIds is the full
// current set of external ids to KEEP; features with the given Source whose
// external_id is not in it are deleted (with cache invalidation + reindex).
type PruneProductFeaturesRequest struct {
	Source          string
	KeepExternalIds []string
}

type PruneProductFeaturesResult struct {
	DeletedFeatures int
}

// ImportProductsRequest is a thin passthrough to the core product import engine
// (ProductImportService.ImportProducts) — the same machinery behind Excel/mass
// import: identifier matching, create/update modes, validation, custom-field
// merge (map a column to "customFields.<slug>"), and BulkUpdateJob progress.
// All values travel as strings (the engine stringifies custom-field values).
type ImportProductsRequest struct {
	// Products are the rows; keys are column names referenced by ColumnsMap and
	// ProductIdentifierColumn.
	Products []map[string]string
	// ColumnsMap maps row column names to system fields (e.g. "customFields.color").
	ColumnsMap map[string]string
	// ProductIdentifierColumn is the ROW column holding the match value.
	ProductIdentifierColumn string
	// ProductSystemIdentifierColumn is the SYSTEM column matched against
	// (articule | barcode | supplierCode1 | orgEanCode).
	ProductSystemIdentifierColumn string
	// ImportMode: "createUpdate" | "createOnly" | "updateOnly".
	ImportMode string
	DryRun     bool
}

type ImportProductsResponse struct {
	ProductsFound   int64
	ProductsCreated int64
	ProductsUpdated int64
	ProductsSkipped int64
}
