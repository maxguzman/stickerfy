package configs

import (
	"net/http"
	"os"
	"stickerfy/pkg/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MuxConfig returns a http server with the mux config
func MuxConfig(router *mux.Router) *http.Server {
	writeTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	addr, _ := utils.URLBuilder("server")

	return &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * time.Duration(writeTimeoutSecondsCount),
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}
