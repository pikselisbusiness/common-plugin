package shared

import (
	"pikselis-business/utils/logger"

	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "COMMON_PLUGIN",
	MagicCookieValue: "common-plugin",
}

func GetPluginMap(api API, db DB) map[string]plugin.Plugin {

	// PluginMap is the map of plugins we can dispense.
	var PluginMap = map[string]plugin.Plugin{
		"common": &CommonPlugin{
			apiImpl: api,
			dbImpl:  db,
			logger:  logger.LoggerExternal,
		},
	}
	return PluginMap
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type CommonPlugin struct {
	// that are written in Go.
	Impl    Common
	apiImpl API
	dbImpl  DB
	Driver  Driver
	logger  logger.Logger
}
