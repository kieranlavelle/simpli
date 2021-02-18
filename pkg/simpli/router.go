package simpli

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	foundRoute, exists := r.matchRoute(path, route.Method)
	if exists {
		log.Fatalf("existing route %v: %v already exists so cant add route %v: %v",
			foundRoute.Method, foundRoute.Path, route.Method, route.Path)
	}

	if r.Routes[path] == nil {
		r.Routes[path] = make(map[string]Route)
	}
	r.Routes[path][route.Method] = route
}

func (r *router) matchRoute(path, method string) (Route, bool) {
	for existingPath, existingMethods := range r.Routes {
		for existingMethod, route := range existingMethods {
			if route.regexMatch && existingMethod == method {
				if route.regexPattern.MatchString(path) {
					names := route.regexPattern.SubexpNames()
					pathParts := strings.Split(path, "/")

					for index, group := range names {
						if group == "" {
							continue
						}

						if index == len(names)-1 {
							finalPart := strings.Join(pathParts[index:], "/")
							route.pathParams[group] = finalPart
						} else {
							route.pathParams[group] = pathParts[index]
						}
					}
					return route, true
				}
			} else if existingPath == path && existingMethod == method {
				return route, true
			}
		}
	}

	return Route{}, false
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

		route, exists := router.matchRoute(r.URL.Path, r.Method)
		if exists {
			s := newState(r, rw, route.pathParams)
			route.Handler(s)

			rw.WriteHeader(s.status)
			if s.json != nil {
				json.NewEncoder(rw).Encode(s.json)
			}

			return
		}

		rw.WriteHeader(404)
		return
	})

	http.ListenAndServe(address, nil)
}
