package shared

type RequestContext struct {
	UserId         uint
	UserGroupId    uint
	OrganizationId uint
	HostName       string
	UserName       string
	UserFirstName  string
	UserLastName   string
	UserEmail      string
	UserIp         string
	AccessedPath   string
	AccessedMethod string
	UserLanguage   string
	IsSystemUser   bool
	IdempotencyKey *string
}
