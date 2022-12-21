package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

// ports represents our REST resource
type ports struct {
	//Create variables based on the ports.json file
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type datastore struct {
	m map[string]ports
	*sync.RWMutex
}

type portsHandler struct {
	store *datastore
}

var (
	// Regex to match the get endpoint
	getPortsRe = regexp.MustCompile(`^/ports/(\d+)$`)
)

func (h *portsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// all users request are going to be routed here
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && getPortsRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *portsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p ports
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	//Here I want to add the new port to the map

	h.store.Unlock()
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	// Initialize a ServerMux with a ports endpoint
	mux := http.NewServeMux()
	mux.Handle("/ports", &portsHandler{})
	http.ListenAndServe(":8080", mux)

}
