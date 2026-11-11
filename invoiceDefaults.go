package shared

import "errors"

// InvoiceDefaults are the organization's defaults for a new invoice, named as
// in the external API (invoicePrefix, operation, currency, taxPercentage).
type InvoiceDefaults struct {
	InvoicePrefix string  `json:"invoicePrefix"`
	Operation     string  `json:"operation"`
	Currency      string  `json:"currency"`
	TaxPercentage float64 `json:"taxPercentage"`
}

// Invoice statuses of InvoiceCreateUpdateRequest.Status.
const (
	InvoiceStatusDraft  = "draft"
	InvoiceStatusPosted = "posted"
)

// APIInvoiceDefaults is the host call behind GetInvoiceDefaults. Optional,
// like APIStopper: not part of API, so plugin test mocks of API keep
// compiling. Plain return values keep the host free of SDK types it may not
// have yet.
type APIInvoiceDefaults interface {
	GetInvoiceDefaults(context RequestContext, isPurchase bool) (invoicePrefix, operation, currency string, taxPercentage float64, err error)
}

// ErrInvoiceDefaultsUnsupported is returned for an API without the call.
var ErrInvoiceDefaultsUnsupported = errors.New("host does not provide invoice defaults")

// GetInvoiceDefaults asks the host for the organization's invoice defaults.
// An older host answers an error ("method not found" over RPC).
func GetInvoiceDefaults(api API, context RequestContext, isPurchase bool) (InvoiceDefaults, error) {
	host, ok := api.(APIInvoiceDefaults)
	if !ok {
		return InvoiceDefaults{}, ErrInvoiceDefaultsUnsupported
	}
	prefix, operation, currency, tax, err := host.GetInvoiceDefaults(context, isPurchase)
	if err != nil {
		return InvoiceDefaults{}, err
	}
	return InvoiceDefaults{InvoicePrefix: prefix, Operation: operation, Currency: currency, TaxPercentage: tax}, nil
}

// RPC

type GetInvoiceDefaultsRequest struct {
	Context    RequestContext
	IsPurchase bool
}
type GetInvoiceDefaultsResponse struct {
	Defaults InvoiceDefaults
	Error    error
}

// GetInvoiceDefaults is served through an assertion: APIInvoiceDefaults is
// optional.
func (m *apiRPCServer) GetInvoiceDefaults(req GetInvoiceDefaultsRequest, resp *GetInvoiceDefaultsResponse) error {
	defaults, err := GetInvoiceDefaults(m.impl, req.Context, req.IsPurchase)
	resp.Defaults = defaults
	resp.Error = encodableError(err)
	return nil
}

func (m *apiRPCClient) GetInvoiceDefaults(context RequestContext, isPurchase bool) (string, string, string, float64, error) {
	var reply GetInvoiceDefaultsResponse
	if err := m.client.Call("Plugin.GetInvoiceDefaults", GetInvoiceDefaultsRequest{Context: context, IsPurchase: isPurchase}, &reply); err != nil {
		return "", "", "", 0, err
	}
	d := reply.Defaults
	return d.InvoicePrefix, d.Operation, d.Currency, d.TaxPercentage, reply.Error
}
