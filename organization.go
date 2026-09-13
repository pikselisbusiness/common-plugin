package shared

import (
	"fmt"
	"os"
	"strconv"
)

// OrganizationContext returns the immutable host-assigned identity for startup,
// cron and plugin-owned storage. Missing identity is an error, never organization 1.
func OrganizationContext() (RequestContext, error) {
	organizationId, err := strconv.ParseUint(os.Getenv("PB_PLUGIN_ORGANIZATION_ID"), 10, strconv.IntSize)
	if err != nil || organizationId == 0 {
		return RequestContext{}, fmt.Errorf("host-assigned plugin organization id is required")
	}
	return RequestContext{OrganizationId: uint(organizationId), IsSystemUser: true}, nil
}
