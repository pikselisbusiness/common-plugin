package shared

import "encoding/json"

type RouteProps struct {
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

type RouteResponseRouter struct {
	Status      int
	Body        []byte
	ContentType string
}

// response code, bytes and application type (e.g. application/json)
type HandlerFunc func(rc RouteContext) RouteResponseRouter

type Route struct {
	Type        string
	Url         string
	HandlerFunc HandlerFunc
}

type InnerRouter struct {
	Routes []*Route
}

func (s *InnerRouter) GET(url string, handlerFunc HandlerFunc) {

	s.Routes = append(s.Routes, &Route{
		Type:        "GET",
		Url:         url,
		HandlerFunc: handlerFunc,
	})
}
func (s *InnerRouter) POST(url string, handlerFunc HandlerFunc) {

	s.Routes = append(s.Routes, &Route{
		Type:        "POST",
		Url:         url,
		HandlerFunc: handlerFunc,
	})
}
func (s *InnerRouter) PUT(url string, handlerFunc HandlerFunc) {

	s.Routes = append(s.Routes, &Route{
		Type:        "PUT",
		Url:         url,
		HandlerFunc: handlerFunc,
	})
}
func (s *InnerRouter) DELETE(url string, handlerFunc HandlerFunc) {

	s.Routes = append(s.Routes, &Route{
		Type:        "DELETE",
		Url:         url,
		HandlerFunc: handlerFunc,
	})
}
func (s *InnerRouter) HandleRoute(routeType, url string, rc RouteContext) RouteResponseRouter {
	for _, route := range s.Routes {
		if route.Type == routeType && route.Url == url {
			return route.HandlerFunc(rc)
		}
	}
	return RouteResponseRouter{500, nil, "application/json"}
}

func ToJSON(statusCode int, i interface{}) RouteResponseRouter {

	jsonBytes, _ := json.Marshal(i)

	return RouteResponseRouter{statusCode, jsonBytes, "application/json"}

}
