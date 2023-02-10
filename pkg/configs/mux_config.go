package configs

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MuxConfig returns a http server with the mux config
func MuxConfig(router *mux.Router, port string) *http.Server {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	return &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}
