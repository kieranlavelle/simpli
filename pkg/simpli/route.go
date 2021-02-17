package simpli

type Route struct {
	Path    string
	Method  string
	Handler func(*State)
}

func createRoute(path, method string, handler func(*State)) Route {
	return Route{path, method, handler}
}
