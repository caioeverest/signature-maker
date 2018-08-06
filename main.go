package main

import (
	"fmt"
	"net/http"

	"./config"
	"./routes"
	"github.com/couchbase/goutils/logging"
)

func main() {
	cfg := config.CreateConfig()

	api := routes.API(cfg)
	logging.Infof(fmt.Sprintf("Starting server at port %s", cfg.HttpPort))
	http.ListenAndServe(":"+cfg.HttpPort, api)
}
