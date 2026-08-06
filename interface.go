// Portions of this file are derived from Mattermost
// (https://github.com/mattermost/mattermost), licensed under the
// Apache License, Version 2.0.
// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE-APACHE-2.0.txt and NOTICE in this module.
// Modified by Pikselis, UAB.

package shared

import (
	"github.com/pikselisbusiness/go-plugin"
)

const CurrentPluginAPIVersion = "2.0.0"

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
			logger:  LoggerExternal,
		},
	}
	return PluginMap
}

// Common plugin implementation for RPC mostly
type CommonPlugin struct {
	// that are written in Go.
	Impl    any
	apiImpl API
	dbImpl  DB
	logger  Logger
}
