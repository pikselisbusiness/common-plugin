package shared

import (
	"github.com/hashicorp/go-plugin"
)

func ClientMain(pluginImplementation any) {

	if impl, ok := pluginImplementation.(interface {
		SetAPI(api API)
		SetDB(db DB)
	}); !ok {
		panic("Plugin implementation given must embed plugin.PBPlugin")
	} else {
		impl.SetAPI(nil)
		impl.SetDB(nil)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			"common": &CommonPlugin{Impl: pluginImplementation},
		},
	})
}

type PBPlugin struct {
	// API exposes the plugin api, and becomes available just prior to the OnActive hook.
	API API
	DB  DB
}

// SetAPI persists the given API interface to the plugin. It is invoked just prior to the
// OnActivate hook, exposing the API for use by the plugin.
func (p *PBPlugin) SetAPI(api API) {
	p.API = api
}

// SetAPI persists the given API interface to the plugin. It is invoked just prior to the
// OnActivate hook, exposing the API for use by the plugin.
func (p *PBPlugin) SetDB(db DB) {
	p.DB = db
}
