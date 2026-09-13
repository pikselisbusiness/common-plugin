package shared

import "testing"

func TestOrganizationContextRequiresHostIdentity(t *testing.T) {
	for _, value := range []string{"", "0", "-1", "invalid", "99999999999999999999999999999999"} {
		t.Setenv("PB_PLUGIN_ORGANIZATION_ID", value)
		if _, err := OrganizationContext(); err == nil {
			t.Fatalf("accepted invalid identity %q", value)
		}
	}
	t.Setenv("PB_PLUGIN_ORGANIZATION_ID", "42")
	context, err := OrganizationContext()
	if err != nil || context.OrganizationId != 42 || !context.IsSystemUser || context.UserId != 0 {
		t.Fatalf("unexpected background context: %+v %v", context, err)
	}
}
