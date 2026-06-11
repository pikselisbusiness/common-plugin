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
	ModuleValueAll    ModuleValue = "all"
	ModuleValueRead   ModuleValue = "read"
	ModuleValueUpdate ModuleValue = "update"
	ModuleValueCreate ModuleValue = "create"
	ModuleValueDelete ModuleValue = "delete"
	ModuleValueAdmin  ModuleValue = "admin"
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
	// Possible to skip auth for route - to implement your own auth
	// Use with caution
	SkipAuth bool
	// RequiredRight is an optional module right value the host can enforce
	// for this route. Empty keeps legacy plugin-level permission behavior.
	RequiredRight ModuleValue
}

type RouteResponse struct {
	Code int
	I    []byte
	// If no content type is provided - application/json is used
	ContentType string
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

func (c *RouteContext) JSON(statusCode int, i interface{}) RouteResponseRouter {
	jsonBytes, _ := json.Marshal(i)
	return RouteResponseRouter{statusCode, jsonBytes, "application/json"}
}

func (c *RouteContext) Param(nameSearch string) string {

	for key, name := range c.ParamNames {
		if name == nameSearch {
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

	// RunCronJobWithTag is invoked when the plugin cronjob schedule activates with a tag
	// mulitple cronjobs can be registered with different schedules
	RunCronJobWithTag(tag string) error

	// GetRoutes returns all routes that plugin has for REST API
	GetRoutes() []RouteUrl
	// Execute route - plugin will handle this
	HandleRoute(routeType, url string, rc RouteContext) RouteResponse

	// Plugins can handle events such as cart, order, etc.
	HandleEvent(context RequestContext, event EventEnvelope) (HandleResult, error)
}

// PluginCapability describes one host capability a plugin wants to use.
// Hosts can use this for install-time review, permission enforcement, and
// admin UI. A capability name should be stable, for example "orders.read".
type PluginCapability struct {
	Name        string         `json:"name"`
	Access      []ModuleValue  `json:"access,omitempty"`
	Description string         `json:"description,omitempty"`
	Optional    bool           `json:"optional,omitempty"`
	Props       map[string]any `json:"props,omitempty"`
}

// PluginRouteCapability is declarative route metadata. It can mirror routes
// returned by GetRoutes while adding information that old RouteUrl fields did
// not carry, such as permission requirements.
type PluginRouteCapability struct {
	Type          string      `json:"type"`
	Url           string      `json:"url"`
	SkipAuth      bool        `json:"skipAuth,omitempty"`
	RequiredRight ModuleValue `json:"requiredRight,omitempty"`
}

// PluginEventSubscription tells the host which event topics should be delivered
// to the plugin. This keeps plugins from receiving broad default topics when
// they only care about a specific business event stream.
type PluginEventSubscription struct {
	Topic   string `json:"topic"`
	Version string `json:"version,omitempty"`
}

// PluginCronCapability documents cron jobs a plugin expects to register.
// Cron execution still happens through RegisterCronJob/RegisterCronJobWithTag.
type PluginCronCapability struct {
	Tag      string `json:"tag,omitempty"`
	Schedule string `json:"schedule"`
}

// PluginUICapability describes where a plugin wants frontend surface area.
// The backend can expose this metadata before every frontend slot is supported.
type PluginUICapability struct {
	Slot      string         `json:"slot"`
	Component string         `json:"component,omitempty"`
	Props     map[string]any `json:"props,omitempty"`
}

// PluginCapabilities is the optional CommonV3 contract. It explains what a
// plugin needs and where it extends the product without changing legacy route,
// cron, or event method signatures.
type PluginCapabilities struct {
	PluginAPIVersion    string                    `json:"pluginApiVersion,omitempty"`
	HostAPIVersionRange string                    `json:"hostApiVersionRange,omitempty"`
	HostVersionRange    string                    `json:"hostVersionRange,omitempty"`
	APIVersion          string                    `json:"apiVersion,omitempty"` // Deprecated: use PluginAPIVersion.
	RequiredHostScopes  []PluginCapability        `json:"requiredHostScopes,omitempty"`
	OptionalHostScopes  []PluginCapability        `json:"optionalHostScopes,omitempty"`
	Routes              []PluginRouteCapability   `json:"routes,omitempty"`
	EventSubscriptions  []PluginEventSubscription `json:"eventSubscriptions,omitempty"`
	CronJobs            []PluginCronCapability    `json:"cronJobs,omitempty"`
	UIExtensions        []PluginUICapability      `json:"uiExtensions,omitempty"`
	Props               map[string]any            `json:"props,omitempty"`
}

// PluginConfigValidationResult is returned before saving or applying plugin
// settings. Field keys should match config field IDs where possible.
type PluginConfigValidationResult struct {
	Valid  bool              `json:"valid"`
	Errors map[string]string `json:"errors,omitempty"`
}

// PluginConfigChange is passed after settings change so a plugin can refresh
// clients, caches, schedules, or local state without a full restart.
type PluginConfigChange struct {
	OldConfig map[string]any `json:"oldConfig,omitempty"`
	NewConfig map[string]any `json:"newConfig,omitempty"`
}

// CommonV3 is optional. Hosts must continue to support Common and CommonV2
// plugins and should use this richer contract only when a plugin implements it.
type CommonV3 interface {
	CommonV2
	PluginCapabilities() (PluginCapabilities, error)
	ValidateConfig(context RequestContext, config map[string]any) (PluginConfigValidationResult, error)
	OnConfigChanged(context RequestContext, change PluginConfigChange) error
}
