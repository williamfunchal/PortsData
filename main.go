package main

import (
	"net/http"
	"regexp"
)

var (
	// Regex to match the get endpoint
	getPortsRe = regexp.MustCompile(`^/ports/(\d+)$`)
)

type portsHandler struct{}

func (h *portsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// all users request are going to be routed here
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getPortsRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func main() {
	// Initialize a ServerMux with a ports endpoint
	mux := http.NewServeMux()
	mux.Handle("/ports", &portsHandler{})
	http.ListenAndServe(":8080", mux)

}
