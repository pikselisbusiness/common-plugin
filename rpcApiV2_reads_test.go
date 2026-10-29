package shared

import (
	"errors"
	v2 "github.com/pikselisbusiness/common-plugin/v2"
	"net"
	"net/rpc"
	"testing"
)

type featureReadProbe struct {
	API
	fail bool
}

func (s *featureReadProbe) V2() v2.API { return featureReadV2{probe: s} }

type featureReadV2 struct {
	v2.API
	probe *featureReadProbe
}

func (s featureReadV2) Products() v2.ProductsAPI { return featureReadProducts{probe: s.probe} }

type featureReadProducts struct {
	v2.ProductsAPI
	probe *featureReadProbe
}

func (s featureReadProducts) GetProductFeatureValues(context RequestContext, id uint) ([]v2.ProductFeatureValue, error) {
	if s.probe.fail {
		return nil, errors.New("read failed")
	}
	if context.OrganizationId != 7 || id != 5 {
		return nil, errors.New("incorrect RPC arguments")
	}
	return []v2.ProductFeatureValue{{FeatureId: 2, FeatureName: "Colour", ValueId: 4, ValueName: "Red"}}, nil
}
func featureReadClient(t *testing.T, impl API) *apiRPCClient {
	t.Helper()
	server := rpc.NewServer()
	if err := server.RegisterName("Plugin", &apiRPCServer{impl: impl}); err != nil {
		t.Fatal(err)
	}
	a, b := net.Pipe()
	client := rpc.NewClient(b)
	t.Cleanup(func() { client.Close() })
	go server.ServeConn(a)
	return &apiRPCClient{client: client}
}
func TestProductFeatureReadRPC(t *testing.T) {
	probe := &featureReadProbe{}
	client := featureReadClient(t, probe)
	values, err := client.V2().Products().GetProductFeatureValues(RequestContext{OrganizationId: 7}, 5)
	if err != nil || len(values) != 1 || values[0].ValueName != "Red" {
		t.Fatalf("feature roundtrip: %v %v", values, err)
	}
	probe.fail = true
	if _, err := client.V2().Products().GetProductFeatureValues(RequestContext{OrganizationId: 7}, 5); err == nil || err.Error() != "read failed" {
		t.Fatalf("error lost: %v", err)
	}
}

type unsupportedFeatureProbe struct{ API }

func TestProductFeatureReadUnsupported(t *testing.T) {
	client := featureReadClient(t, &unsupportedFeatureProbe{})
	if _, err := client.V2().Products().GetProductFeatureValues(RequestContext{}, 1); err == nil {
		t.Fatal("unsupported feature read succeeded")
	}
}
