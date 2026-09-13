package shared

import (
	"errors"
	"net"
	"net/rpc"
	"testing"
)

type lifecycleProbe struct{ activated, deactivated bool }

func (s *lifecycleProbe) OnActivate() error {
	s.activated = true
	return errors.New("activation failed")
}
func (s *lifecycleProbe) OnDeactivate() error {
	s.deactivated = true
	return errors.New("deactivation failed")
}

// Expose the post-handshake activation path over RPC to verify error encoding.
type LifecycleActivationRPC struct{ CommonServerRPC }

func (s *LifecycleActivationRPC) OnActivate(_ struct{}, response *OnActivateResponse) error {
	response.A = encodableError(s.activatePlugin())
	return nil
}

func TestLifecycleHooksReturnErrorsOverRPC(t *testing.T) {
	probe := &lifecycleProbe{}
	server := rpc.NewServer()
	if err := server.RegisterName("Plugin", &LifecycleActivationRPC{CommonServerRPC{Impl: probe}}); err != nil {
		t.Fatal(err)
	}
	serverConnection, clientConnection := net.Pipe()
	client := rpc.NewClient(clientConnection)
	defer client.Close()
	go server.ServeConn(serverConnection)
	var response OnActivateResponse
	if err := client.Call("Plugin.OnActivate", struct{}{}, &response); err != nil {
		t.Fatal(err)
	}
	if !probe.activated || response.A == nil || response.A.Error() != "activation failed" {
		t.Fatalf("activation error lost: %v", response.A)
	}
	pluginClient := &CommonClientRPC{client: client}
	if err := pluginClient.OnDeactivate(); !probe.deactivated || err == nil || err.Error() != "deactivation failed" {
		t.Fatalf("deactivation hook/error lost: %v", err)
	}
}
