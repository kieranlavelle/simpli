package simpli

import "net/http"

type State struct {
	Request        *http.Request
	responseWriter http.ResponseWriter
	status         int
	json           interface{}
}

func newState(r *http.Request, rw http.ResponseWriter) *State {
	return &State{
		Request:        r,
		responseWriter: rw,
		status:         200,
		json:           nil,
	}
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
