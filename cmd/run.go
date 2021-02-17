package main

import (
	"log"
	"net/http"

	"github.com/kieranlavelle/simpli/pkg/simpli"
)

func test(s *simpli.State) {
	log.Printf("Hello from: %v", s.Request.URL.Path)

	s.JSONResponse(http.StatusOK, simpli.J{
		"detail": "working",
	})
}

func main() {

	r := simpli.New()

	r.GET("/test", test)
	r.Run("0.0.0.0:10000")
}
