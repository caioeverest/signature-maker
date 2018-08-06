package main

import (
	"fmt"
	"net/http"

	"github.com/couchbase/goutils/logging"

	"github.com/caioever/signature-maker/config"
	"github.com/caioever/signature-maker/routes"
)

func main() {
	cfg := config.CreateConfig()

	api := routes.API(cfg)
	logging.Infof(fmt.Sprintf("Starting server at port %s", cfg.HttpPort))
	http.ListenAndServe(":"+cfg.HttpPort, api)
}
