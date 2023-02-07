package router

import "net/http"

// Router is an interface for a router
type Router interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Put(uri string, f func(w http.ResponseWriter, r *http.Request))
	Delete(uri string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}
