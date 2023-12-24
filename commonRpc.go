package shared

import (
	"database/sql"
	"database/sql/driver"
	"encoding/gob"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"pikselis-business/utils/logger"
	"reflect"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-plugin"
)

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
	logger  logger.Logger
}

type CommonServerRPC struct {
	Impl         Common
	broker       *plugin.MuxBroker
	apiRPCClient *apiRPCClient
	dbRPCClient  *dbRPCClient
}

type apiRPCClient struct {
	client *rpc.Client
	broker *plugin.MuxBroker
}

type apiRPCServer struct {
	impl   API
	broker *plugin.MuxBroker
}

type dbRPCClient struct {
	client *rpc.Client
	broker *plugin.MuxBroker
}

type dbRPCServer struct {
	impl   DB
	broker *plugin.MuxBroker
}

type HandleRouteRequest struct {
	RouteType    string
	Url          string
	RouteContext RouteContext
}

type HandleRouteResponse struct {
	Code int
	I    []byte
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
type ServeHTTPRequest struct {
	ResponseWriterStream uint32
	Request              *http.Request
	Context              *RequestContext
	RequestBodyStream    uint32
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
func init() {
	gob.Register(ErrorString{})
}

type ErrorString struct {
	Code int // Code to map to various error variables
	Err  string
}

func (e ErrorString) Error() string {
	return e.Err
}

func (m *CommonServerRPC) GetRoutes(req EmptyRequest, resp *GetRoutesResponse) error {

	routes := m.Impl.GetRoutes()

	fmt.Println("GET routes on server", routes)

	resp.RouteUrls = routes

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
	response := m.Impl.HandleRoute(req.RouteType, req.Url, req.RouteContext)

	resp.Code = response.Code
	resp.I = response.I
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

func (m *CommonServerRPC) ServeHTTP(req ServeHTTPRequest, rep *struct{}) error {

	return nil
}

func (m *CommonClientRPC) ServeHTTP(rc *RequestContext, w http.ResponseWriter, r *http.Request) {

}

//GetConfigVariable

type GetConfigVariableRequest struct {
	VariableName string
}
type GetConfigVariableResponse struct {
	VariableData string
	Error        error
}

//GetUserInfoForUserId

type GetUserInfoForUserIdRequest struct {
	UserId uint
}
type GetUserInfoForUserIdResponse struct {
	UserInfo UserInfo
}

// GetInvoices
type GetInvoicesRequest struct {
	Request InvoicesRequest
}
type GetInvoicesResponse struct {
	Response InvoicesListResponse
	Error    error
}

// GetDivisions
type GetDivisionsRequest struct {
	Context RequestContext
	Request DivisionsRequest
}
type GetDivisionsResponse struct {
	Divisions []Division
	Error     error
}

func (m *apiRPCServer) GetConfigVariable(req GetConfigVariableRequest, resp *GetConfigVariableResponse) error {
	data, err := m.impl.GetConfigVariable(req.VariableName)

	resp.VariableData = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetConfigVariable(variableName string) (string, error) {

	var reply GetConfigVariableResponse
	err := m.client.Call("Plugin.GetConfigVariable", GetConfigVariableRequest{
		VariableName: variableName,
	}, &reply)
	if err != nil {
		return "", err
	}

	return reply.VariableData, reply.Error
}

func (m *apiRPCServer) GetUserInfoForUserId(req GetUserInfoForUserIdRequest, resp *GetUserInfoForUserIdResponse) error {
	data := m.impl.GetUserInfoForUserId(req.UserId)

	resp.UserInfo = data

	return nil
}
func (m *apiRPCClient) GetUserInfoForUserId(userId uint) UserInfo {

	var reply GetUserInfoForUserIdResponse
	err := m.client.Call("Plugin.GetUserInfoForUserId", GetUserInfoForUserIdRequest{
		UserId: userId,
	}, &reply)
	if err != nil {
		return UserInfo{}
	}

	return reply.UserInfo
}

func (m *apiRPCServer) GetInvoices(req GetInvoicesRequest, resp *GetInvoicesResponse) error {
	data, err := m.impl.GetInvoices(req.Request)

	resp.Response = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetInvoices(request InvoicesRequest) (InvoicesListResponse, error) {

	var reply GetInvoicesResponse
	err := m.client.Call("Plugin.GetInvoices", GetInvoicesRequest{
		Request: request,
	}, &reply)
	if err != nil {
		return InvoicesListResponse{}, err
	}

	return reply.Response, reply.Error
}

func (m *apiRPCServer) GetDivisions(req GetDivisionsRequest, resp *GetDivisionsResponse) error {
	data, err := m.impl.GetDivisions(req.Context, req.Request)

	resp.Divisions = data
	resp.Error = err

	return nil
}
func (m *apiRPCClient) GetDivisions(context RequestContext, request DivisionsRequest) ([]Division, error) {

	var reply GetDivisionsResponse
	err := m.client.Call("Plugin.GetDivisions", GetDivisionsRequest{
		Context: context,
		Request: request,
	}, &reply)
	if err != nil {
		return []Division{}, err
	}

	return reply.Divisions, reply.Error
}

// api
type DbMigrateRequest struct {
	Model interface{}
}

type DbMigrateResponse struct {
	Error error
}
type DbStatementRequest struct {
	Sql    string
	Values []interface{}
}

type DbStatementResponse struct {
	Error error
}

type DbRawResponse struct {
	Items []map[string]interface{}
	Error error
}

func (m *dbRPCServer) Migrate(req DbMigrateRequest, resp *DbMigrateResponse) error {
	err := m.impl.Migrate(req.Model)

	if err != nil {
		fmt.Println("Error calling Plugin.Migrate", "error", err)
	}
	resp.Error = err

	return nil
}
func (m *dbRPCClient) Migrate(model interface{}) error {

	var reply DbMigrateResponse
	err := m.client.Call("Plugin.Migrate", DbMigrateRequest{
		Model: model,
	}, &reply)
	if err != nil {
		fmt.Println("Error calling Plugin.Migrate", "error", err)
		return err
	}

	return reply.Error
}

func (m *dbRPCServer) Exec(req DbStatementRequest, resp *DbStatementResponse) error {
	err := m.impl.Exec(req.Sql, req.Values...)

	if err != nil {
		fmt.Println("Error calling Plugin.Exec", "error", err)
	}
	resp.Error = encodableError(err)

	return nil
}
func (m *dbRPCClient) Exec(sql string, values ...interface{}) error {

	var reply DbStatementResponse
	err := m.client.Call("Plugin.Exec", DbStatementRequest{
		Sql:    sql,
		Values: values,
	}, &reply)
	if err != nil {
		fmt.Println("Error calling Plugin.Exec", "error", err)
		return err
	}

	return reply.Error
}

func (m *dbRPCServer) Raw(req DbStatementRequest, resp *DbRawResponse) error {
	items, err := m.impl.Raw(req.Sql, req.Values...)

	if err != nil {
		fmt.Println("Error calling Plugin.Raw", "error", err)
	}
	resp.Items = items
	resp.Error = encodableError(err)

	return nil
}
func (m *dbRPCClient) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {

	var reply DbRawResponse
	err := m.client.Call("Plugin.Raw", DbStatementRequest{
		Sql:    sql,
		Values: values,
	}, &reply)
	if err != nil {
		fmt.Println("Error calling Plugin.Raw", "error", err)
		return []map[string]interface{}{}, err
	}

	return reply.Items, reply.Error
}
