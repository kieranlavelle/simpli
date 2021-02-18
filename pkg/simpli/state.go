package simpli

import "net/http"

type State struct {
	Request        *http.Request
	Params         map[string]string
	responseWriter http.ResponseWriter
	status         int
	json           interface{}
}

type J map[string]interface{}

func newState(r *http.Request, rw http.ResponseWriter, params map[string]string) *State {
	return &State{
		Request:        r,
		Params:         params,
		responseWriter: rw,
		status:         200,
		json:           nil,
	}
}

func (s *State) Param(key string) string {
	return s.Params[key]
}

func (s *State) SetHeader(key, value string) {
	s.responseWriter.Header().Add(key, value)
}

func (s *State) GetHeader(key string) string {
	return s.Request.Header.Get(key)
}

func (s *State) JSONResponse(status int, json interface{}) {
	s.status = status
	s.json = json
	s.SetHeader("Content-Type", "application/json")
}
