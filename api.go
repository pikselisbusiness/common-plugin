package shared

// directly.
type API interface {
	GetConfigVariable(name string) (string, error)
	GetUserInfoForUserId(userId uint) (UserInfo)
	GetInvoices(request InvoicesRequest) (InvoicesListResponse, error) 
	GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error)
}

// Client is a streamlined wrapper over the pikselis-business plugin API.
type Client struct {
	Api API
	Db  DB
}

// NewClient creates a new instance of Client.
//
// This client must only be created once per plugin to
// prevent reacquiring of resources.
func NewClient(api API, db DB) *Client {
	return &Client{
		Api: api,
		Db:  db,
	}
}
