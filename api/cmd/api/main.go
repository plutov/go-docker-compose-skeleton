// API
//
//     Schemes: http
//     Host: 127.0.0.1:8080
//     BasePath: /
//     Version: 0.1.0
//     Contact: a.pliutau@gmail.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
//go:generate swagger generate spec -o ../../swagger.json
package main

import (
	"github.com/plutov/go-docker-compose-skeleton/api/pkg/api"
	"github.com/plutov/go-docker-compose-skeleton/api/pkg/env"
	"github.com/plutov/go-docker-compose-skeleton/api/pkg/shared"
	"os"
)

func main() {
	e, err := env.New()
	shared.LogErr(err)
	if err != nil {
		os.Exit(1)
	}

	api.ListenAndServe(e)
}
