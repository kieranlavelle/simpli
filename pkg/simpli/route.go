package simpli

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Route struct {
	Path         string
	Method       string
	Handler      func(*State)
	regexMatch   bool
	regexPattern *regexp.Regexp
	pathParams   map[string]string
}

func getRegexPath(path string) *regexp.Regexp {
	parts := strings.Split(path, "/")
	newPath := []string{}
	for _, part := range parts {

		if strings.Contains(part, "*") {
			groupName := strings.Replace(part, "*", "", -1)
			regexPart := fmt.Sprintf("(?P<%v>.+)", groupName)
			newPath = append(newPath, regexPart)
			break
		} else if strings.Contains(part, ":") {
			groupName := strings.Replace(part, ":", "", -1)
			regexPart := fmt.Sprintf("(?P<%v>.+)", groupName)
			newPath = append(newPath, regexPart)
		} else {
			newPath = append(newPath, part)
		}
	}

	pattern, err := regexp.Compile(strings.Join(newPath, "/"))
	if err != nil {
		log.Fatalf("error creting path regex: %v", err)
	}
	return pattern
}

func createRoute(path, method string, handler func(*State)) Route {
	regexMatch := strings.Contains(path, ":")
	pattern := &regexp.Regexp{}
	if regexMatch {
		pattern = getRegexPath(path)
	}

	return Route{
		Path:         path,
		Method:       method,
		Handler:      handler,
		regexMatch:   regexMatch,
		regexPattern: pattern,
		pathParams:   make(map[string]string),
	}
}
