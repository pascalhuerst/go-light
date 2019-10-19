package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pascalhuerst/go-light/data"
)

func jsonRespose(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		fmt.Printf("Cannot marshal JSON for %s: %v", r.RequestURI, err)
		return
	}

	w.Write(b)
}

// StartHTTPServer start the backend server
func StartHTTPServer(port uint16, fd *data.FixtureDefinition) {
	http.HandleFunc("/fixture", func(w http.ResponseWriter, r *http.Request) {

		jsonRespose(w, r, fd)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
