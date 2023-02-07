package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	mux *mux.Router
}

// NewMuxRouter creates a new mux router
func NewMuxRouter() Router {
	return &muxRouter{mux: mux.NewRouter()}
}

func (mr *muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.mux.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (mr *muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.mux.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (mr *muxRouter) Put(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.mux.HandleFunc(uri, f).Methods(http.MethodPut)
}

func (mr *muxRouter) Delete(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.mux.HandleFunc(uri, f).Methods(http.MethodDelete)
}

func (mr *muxRouter) Serve(port string) {
	fmt.Printf("Mux server running on port %s", port)
	http.ListenAndServe(port, mr.mux)
}
