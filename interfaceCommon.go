package shared

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var (
	HeaderContentType = "Content-Type"

	MIMEApplicationJSON = "application/json"
)

type ModuleRightDescription struct {
	LangShortCode string `json:"langShortCode"`
	Description   string `json:"description"`
}

type ModuleValue string

var (
	ModuleValueAll               ModuleValue = "all"
	ModuleValueRead              ModuleValue = "read"
	ModuleValueUpdateModuleValue ModuleValue = "update"
	ModuleValueCreate            ModuleValue = "create"
	ModuleValueDelete            ModuleValue = "delete"
	ModuleValueAdmin             ModuleValue = "admin"
)

type ModuleRight struct {
	Value        ModuleValue              `json:"value"`
	Descriptions []ModuleRightDescription `json:"descriptions"`
}

// For checking whether right is allowed for route context
// E.g. for user that calls route
type ModuleRightPermission struct {
	Value     ModuleValue `json:"value"`
	IsAllowed bool        `json:"isAllowed"`
}

type RouteUrl struct {
	Url  string
	Type string
}

type RouteResponse struct {
	Code int
	I    []byte
}

type RouteContext struct {
	Request           *http.Request
	RequestBody       []byte
	QueryParams       url.Values
	ParamNames        []string
	ParamValues       []string
	RequestContext    RequestContext
	ModulePermissions []ModuleRightPermission
}

func (c *RouteContext) JSON(statusCode int, i interface{}) (int, []byte) {
	jsonTest, _ := json.Marshal(i)
	return statusCode, jsonTest
}

func (c *RouteContext) Param(name string) string {

	for key, name := range c.ParamNames {
		if name == name {
			return c.ParamValues[key]
		}
	}

	return ""
}
func (c *RouteContext) QueryParam(name string) string {
	return c.QueryParams.Get(name)
}
func (c *RouteContext) Bind(i interface{}) (err error) {
	req := c.Request
	if req.ContentLength == 0 {
		return
	}

	ctype := req.Header.Get(HeaderContentType)
	switch {
	case strings.HasPrefix(ctype, MIMEApplicationJSON):
		err := json.Unmarshal(c.RequestBody, &i)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *RouteContext) GetModuleAllowed(value ModuleValue) bool {

	for _, permission := range c.ModulePermissions {
		if permission.Value == value {
			return permission.IsAllowed
		}
	}
	return false
}

type Common interface {
	// OnActivate is invoked when the plugin is activated. If an error is returned, the plugin
	// will be terminated. The plugin will not receive hooks until after OnActivate returns
	// without error. OnConfigurationChange will be called once before OnActivate.
	//
	// Minimum server version: 1.0
	OnActivate() error

	// Implemented returns a list of hooks that are implemented by the plugin.
	// Plugins do not need to provide an implementation. Any given will be ignored.
	//
	// Minimum server version: 1.0
	Implemented() ([]string, error)

	// OnDeactivate is invoked when the plugin is deactivated. This is the plugin's last chance to
	// use the API, and the plugin will be terminated shortly after this invocation. The plugin
	// will stop receiving hooks just prior to this method being called.
	//
	// Minimum server version: 1.0
	OnDeactivate() error

	// RunCronJob is invoked when the plugin cronjob schedule activates
	RunCronJob() error

	// GetRoutes returns all routes that plugin has for REST API
	GetRoutes() []RouteUrl
	// Execute route - plugin will handle this
	HandleRoute(routeType, url string, rc RouteContext) RouteResponse
}

// For additional hooks
// Will be merged with common in the future
type CommonV2 interface {
	// OnActivate is invoked when the plugin is activated. If an error is returned, the plugin
	// will be terminated. The plugin will not receive hooks until after OnActivate returns
	// without error. OnConfigurationChange will be called once before OnActivate.
	//
	// Minimum server version: 1.0
	OnActivate() error

	// Implemented returns a list of hooks that are implemented by the plugin.
	// Plugins do not need to provide an implementation. Any given will be ignored.
	//
	// Minimum server version: 1.0
	Implemented() ([]string, error)

	// OnDeactivate is invoked when the plugin is deactivated. This is the plugin's last chance to
	// use the API, and the plugin will be terminated shortly after this invocation. The plugin
	// will stop receiving hooks just prior to this method being called.
	//
	// Minimum server version: 1.0
	OnDeactivate() error

	// OnMigrateRights is invoked after plugin is activated. This is a hook for migrating module rights
	// plugins can have various rights for routes/services
	//
	OnMigrateRights() []ModuleRight

	// RunCronJob is invoked when the plugin cronjob schedule activates
	RunCronJob() error

	// GetRoutes returns all routes that plugin has for REST API
	GetRoutes() []RouteUrl
	// Execute route - plugin will handle this
	HandleRoute(routeType, url string, rc RouteContext) RouteResponse
}
