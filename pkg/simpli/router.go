package simpli

import (
	"encoding/json"
	"log"
	"net/http"
)

type router struct {
	Routes map[string]map[string]Route
}

func New() *router {
	return &router{
		Routes: make(map[string]map[string]Route),
	}
}

func (r *router) addRoute(path string, route Route) {
	if methods, ok := r.Routes[path]; ok {
		if route, ok := methods[route.Method]; ok {
			log.Fatalf("existing route %v: %v already exists", route.Method, route.Path)
		}
	}
	if r.Routes[path] == nil {
		r.Routes[path] = make(map[string]Route)
	}
	r.Routes[path][route.Method] = route
}

func (r *router) GET(path string, handler func(*State)) {
	route := createRoute(path, "GET", handler)
	r.addRoute(path, route)
}

func (r *router) POST(path string, handler func(*State)) {
	route := createRoute(path, "POST", handler)
	r.addRoute(path, route)
}

func (r *router) PUT(path string, handler func(*State)) {
	route := createRoute(path, "PUT", handler)
	r.addRoute(path, route)
}

func (r *router) DELETE(path string, handler func(*State)) {
	route := createRoute(path, "DELETE", handler)
	r.addRoute(path, route)
}

func (r *router) OPTIONS(path string, handler func(*State)) {
	route := createRoute(path, "OPTIONS", handler)
	r.addRoute(path, route)
}

func (router *router) Run(address string) {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		methods, ok := router.Routes[r.URL.Path]
		if ok {
			route, ok := methods[r.Method]
			if ok {
				state := newState(r, rw)
				route.Handler(state)

				rw.WriteHeader(state.status)
				if state.json != nil {
					json.NewEncoder(rw).Encode(state.json)
				}

				return
			}
		}
		rw.WriteHeader(404)
		return
	})

	http.ListenAndServe(address, nil)
}
