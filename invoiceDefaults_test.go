package shared

import (
	"bytes"
	"encoding/gob"
	"errors"
	"testing"
)

type invoiceDefaultsProbe struct {
	API
	purchase bool
}

func (p *invoiceDefaultsProbe) GetInvoiceDefaults(context RequestContext, isPurchase bool) (string, string, string, float64, error) {
	p.purchase = isPurchase
	if context.OrganizationId == 0 {
		return "", "", "", 0, errors.New("no organization")
	}
	return "INV", "Purchase", "EUR", 21, nil
}

func TestInvoiceDefaultsOverRPC(t *testing.T) {
	probe := &invoiceDefaultsProbe{}
	client := featureReadClient(t, probe)
	defaults, err := GetInvoiceDefaults(client, RequestContext{OrganizationId: 7}, true)
	if err != nil || defaults != (InvoiceDefaults{InvoicePrefix: "INV", Operation: "Purchase", Currency: "EUR", TaxPercentage: 21}) || !probe.purchase {
		t.Fatalf("defaults: %+v %v", defaults, err)
	}
	if _, err := GetInvoiceDefaults(client, RequestContext{}, false); err == nil || err.Error() != "no organization" {
		t.Fatalf("error lost: %v", err)
	}
}

type noInvoiceDefaultsHost struct{ API }

func TestInvoiceDefaultsOnOlderHost(t *testing.T) {
	client := featureReadClient(t, &noInvoiceDefaultsHost{})
	if _, err := GetInvoiceDefaults(client, RequestContext{OrganizationId: 7}, false); err == nil {
		t.Fatal("a host without the call must answer an error")
	}
	if _, err := GetInvoiceDefaults(&noInvoiceDefaultsHost{}, RequestContext{}, false); !errors.Is(err, ErrInvoiceDefaultsUnsupported) {
		t.Fatalf("an API without the call: %v", err)
	}
}

// The invoice status travels to the host with the create request.
func TestInvoiceStatusCrossesRPC(t *testing.T) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(InvoiceCreateUpdateRequest{IsPurchase: true, Status: InvoiceStatusDraft}); err != nil {
		t.Fatal(err)
	}
	var decoded InvoiceCreateUpdateRequest
	if err := gob.NewDecoder(&buffer).Decode(&decoded); err != nil || decoded.Status != InvoiceStatusDraft {
		t.Fatalf("status: %q %v", decoded.Status, err)
	}
}
