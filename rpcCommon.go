package shared

import (
	"database/sql"
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/url"
	"reflect"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/pikselisbusiness/go-plugin"
)

func init() {
	gob.Register(ErrorString{})
	gob.Register(json.RawMessage{})
	gob.Register(HandleRouteRequest{})
	gob.Register(HandleRouteResponse{})
	gob.Register(RouteContext{})
	gob.Register(RequestContext{})
	gob.Register(&http.Request{})
	gob.Register(ModuleRightPermission{})
	gob.Register(&url.Values{})

}

func (p *CommonPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &CommonServerRPC{
		Impl:   p.Impl,
		broker: broker,
	}, nil
}

func (p *CommonPlugin) Client(broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &CommonClientRPC{
		client:  c,
		broker:  broker,
		apiImpl: p.apiImpl,
		dbImpl:  p.dbImpl,
		logger:  p.logger,
	}, nil
}

type CommonClientRPC struct {
	client  *rpc.Client
	apiImpl API
	dbImpl  DB
	broker  *plugin.MuxBroker
	doneWg  sync.WaitGroup
	logger  Logger
}

type CommonServerRPC struct {
	Impl         any
	broker       *plugin.MuxBroker
	apiRPCClient *apiRPCClient
	dbRPCClient  *dbRPCClient
}

type HandleRouteRequest struct {
	RouteType    string
	Url          string
	RouteContext RouteContext
}

type HandleRouteResponse struct {
	Code        int
	I           []byte
	ContentType string
}

type GetRoutesResponse struct {
	RouteUrls []RouteUrl
}
type OnActivateRequest struct {
	APIMuxId uint32
	DBMuxId  uint32
}
type OnActivateResponse struct {
	A error
}
type OnMigrateRightsRequest struct {
}
type OnMigrateRightsResponse struct {
	Rights []ModuleRight
}
type RunCronJobRequest struct {
}
type RunCronJobResponse struct {
	Error error
}

type RunCronJobWithTagRequest struct {
	Tag string
}
type RunCronJobWithTagResponse struct {
	Error error
}

type EmptyRequest struct{}
type EmptyResponse struct{}

func encodableError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*mysql.MySQLError); ok {
		return err
	}
	ret := &ErrorString{
		Err: err.Error(),
	}
	switch err {
	case io.EOF:
		ret.Code = 1
	case sql.ErrNoRows:
		ret.Code = 2
	case sql.ErrConnDone:
		ret.Code = 3
	case sql.ErrTxDone:
		ret.Code = 4
	case driver.ErrSkip:
		ret.Code = 5
	case driver.ErrBadConn:
		ret.Code = 6
	case driver.ErrRemoveArgument:
		ret.Code = 7
	}
	return ret
}
func decodableError(err error) error {
	if encErr, ok := err.(*ErrorString); ok {
		switch encErr.Code {
		case 1:
			return io.EOF
		case 2:
			return sql.ErrNoRows
		case 3:
			return sql.ErrConnDone
		case 4:
			return sql.ErrTxDone
		case 5:
			return driver.ErrSkip
		case 6:
			return driver.ErrBadConn
		case 7:
			return driver.ErrRemoveArgument
		}
	}
	return err
}

type ErrorString struct {
	Code int // Code to map to various error variables
	Err  string
}

func (e ErrorString) Error() string {
	return e.Err
}

func (m *CommonServerRPC) GetRoutes(req EmptyRequest, resp *GetRoutesResponse) error {

	// Check if implemented
	if hook, ok := m.Impl.(interface {
		GetRoutes() []RouteUrl
	}); ok {

		routes := hook.GetRoutes()

		fmt.Println("GET routes on server", routes)

		resp.RouteUrls = routes
	}
	return nil

}
func (m *CommonClientRPC) GetRoutes() []RouteUrl {

	var reply GetRoutesResponse

	err := m.client.Call("Plugin.GetRoutes", EmptyRequest{}, &reply)
	if err != nil {
		fmt.Println("Error getting routes", "error", err)
		return []RouteUrl{}
	}

	return reply.RouteUrls
}

func (m *CommonServerRPC) HandleRoute(req HandleRouteRequest, resp *HandleRouteResponse) error {
	// Check if implemented
	if hook, ok := m.Impl.(interface {
		HandleRoute(routeType, url string, rc RouteContext) RouteResponse
	}); ok {

		response := hook.HandleRoute(req.RouteType, req.Url, req.RouteContext)

		resp.Code = response.Code
		resp.I = response.I
		resp.ContentType = response.ContentType
	} else {
		fmt.Println("Unimplemented hook.HandleRoute On server side")
	}
	return nil
}
func (m *CommonClientRPC) HandleRoute(routeType, url string, rc RouteContext) RouteResponse {

	var reply RouteResponse
	err := m.client.Call("Plugin.HandleRoute", HandleRouteRequest{
		RouteType:    routeType,
		Url:          url,
		RouteContext: rc,
	}, &reply)
	if err != nil {
		m.logger.Error("Error calling hook.HandleRoute On client side", "error", err)
		return RouteResponse{}
	}

	return reply
}

// Get rpc implemented methods
func (m *CommonServerRPC) Implemented(args struct{}, reply *[]string) error {
	ifaceType := reflect.TypeOf((*Common)(nil)).Elem()
	implType := reflect.TypeOf(m.Impl)
	selfType := reflect.TypeOf(m)
	var methods []string
	for i := 0; i < ifaceType.NumMethod(); i++ {
		method := ifaceType.Method(i)
		if m, ok := implType.MethodByName(method.Name); !ok {
			continue
		} else if m.Type.NumIn() != method.Type.NumIn()+1 {
			continue
		} else if m.Type.NumOut() != method.Type.NumOut() {
			continue
		} else {
			match := true
			for j := 0; j < method.Type.NumIn(); j++ {
				if m.Type.In(j+1) != method.Type.In(j) {
					match = false
					break
				}
			}
			for j := 0; j < method.Type.NumOut(); j++ {
				if m.Type.Out(j) != method.Type.Out(j) {
					match = false
					break
				}
			}
			if !match {
				continue
			}
		}
		if _, ok := selfType.MethodByName(method.Name); !ok {
			continue
		}
		methods = append(methods, method.Name)
	}
	*reply = methods
	return nil
}
func (m *CommonClientRPC) Implemented() ([]string, error) {

	type ImplementedReply struct {
		Implemented []string
		Err         error
	}
	var reply ImplementedReply

	err := m.client.Call("Plugin.Implemented", EmptyRequest{}, &reply)
	if err != nil {
		return []string{}, err
	}
	return reply.Implemented, reply.Err
}
func (m *CommonServerRPC) OnActivate(args OnActivateRequest, resp *OnActivateResponse) error {
	connection, err := m.broker.Dial(args.APIMuxId)
	if err != nil {
		return err
	}

	connectionDb, err := m.broker.Dial(args.DBMuxId)
	if err != nil {
		return err
	}

	m.apiRPCClient = &apiRPCClient{
		client: rpc.NewClient(connection),
		broker: m.broker,
	}

	m.dbRPCClient = &dbRPCClient{
		client: rpc.NewClient(connectionDb),
		broker: m.broker,
	}

	if psplugin, ok := m.Impl.(interface {
		SetAPI(api API)
		SetDB(db DB)
	}); ok {
		psplugin.SetAPI(m.apiRPCClient)
		psplugin.SetDB(m.dbRPCClient)
	}

	if common, ok := m.Impl.(interface {
		OnActivate() error
	}); ok {
		resp.A = (common.OnActivate())
	}

	resp.A = nil
	return nil
}
func (m *CommonClientRPC) OnActivate() error {

	muxId := m.broker.NextId()
	m.doneWg.Add(1)
	go func() {
		defer m.doneWg.Done()
		m.broker.AcceptAndServe(muxId, &apiRPCServer{
			impl:   m.apiImpl,
			broker: m.broker,
		})
	}()

	dbMuxId := m.broker.NextId()
	m.doneWg.Add(1)
	go func() {
		defer m.doneWg.Done()
		m.broker.AcceptAndServe(dbMuxId, &dbRPCServer{
			impl: m.dbImpl,
			// broker: m.broker,
		})
	}()

	var reply OnActivateResponse

	err := m.client.Call("Plugin.OnActivate", OnActivateRequest{
		APIMuxId: muxId,
		DBMuxId:  dbMuxId,
	}, &reply)
	if err != nil {
		m.logger.Error("Error calling Plugin.OnActivate", "error", err)
		return err
	}
	return reply.A
}
func (m *CommonServerRPC) OnDeactivate(args struct{}, resp *OnActivateResponse) error {
	resp.A = nil
	return nil
}

func (m *CommonClientRPC) OnDeactivate() error {

	var reply error

	err := m.client.Call("Plugin.OnDeactive", EmptyRequest{}, &reply)
	if err != nil {
		return err
	}
	return reply
}
func (m *CommonServerRPC) OnMigrateRights(args struct{}, resp *OnMigrateRightsResponse) error {

	// Check if implemented
	if hook, ok := m.Impl.(interface {
		OnMigrateRights() []ModuleRight
	}); ok {
		rights := hook.OnMigrateRights()
		resp.Rights = rights
	}
	return nil

}

func (m *CommonClientRPC) OnMigrateRights() []ModuleRight {

	var reply OnMigrateRightsResponse
	err := m.client.Call("Plugin.OnMigrateRights", OnMigrateRightsRequest{}, &reply)
	if err != nil {
		return []ModuleRight{}
	}

	return reply.Rights
}
func (m *CommonServerRPC) RunCronJob(args *RunCronJobRequest, resp *RunCronJobResponse) error {
	// Check if implemented
	if hook, ok := m.Impl.(interface {
		RunCronJob() error
	}); ok {
		err := hook.RunCronJob()
		resp.Error = err
	}
	return nil
}

func (m *CommonClientRPC) RunCronJob() error {

	var reply RunCronJobResponse

	err := m.client.Call("Plugin.RunCronJob", EmptyRequest{}, &reply)
	if err != nil {
		return err
	}
	return reply.Error
}

func (m *CommonServerRPC) RunCronJobWithTag(args *RunCronJobWithTagRequest, resp *RunCronJobWithTagResponse) error {
	// Check if implemented
	if hook, ok := m.Impl.(interface {
		RunCronJobWithTag(tag string) error
	}); ok {
		err := hook.RunCronJobWithTag(args.Tag)
		resp.Error = err
	}
	return nil
}

func (m *CommonClientRPC) RunCronJobWithTag(tag string) error {

	var reply RunCronJobWithTagResponse

	err := m.client.Call("Plugin.RunCronJobWithTag", RunCronJobWithTagRequest{
		Tag: tag,
	}, &reply)
	if err != nil {
		return err
	}
	return reply.Error
}
