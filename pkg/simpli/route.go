package simpli

import "strings"

type Route struct {
	Path       string
	Method     string
	Handler    func(*State)
	regexMatch bool
}

func (r *Route) getRegexPath() string {
	parts := strings.Split(r.Path, "/")
	newPath := []string{}
	for _, part := range parts {
		if strings.Contains(part, ":") {
			newPath = append(newPath, ".+")
		} else {
			newPath = append(newPath, part)
		}
	}

	return strings.Join(newPath, "/")
}

func createRoute(path, method string, handler func(*State)) Route {
	regexMatch := strings.Contains(path, ":")
	return Route{path, method, handler, regexMatch}
}
