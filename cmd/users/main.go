// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/telemetry"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	c, _ := config.LoadConfig(*configFlag)

	// Initialize telemetry
	ctx := context.Background()
	shutdown, err := telemetry.InitTelemetry(ctx, &c.Telemetry, "users")
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown()

	a := app.NewUser(&c.Users)
	a.RegisterRoutes(http.DefaultServeMux)
	_ = http.ListenAndServe(c.Server.Endpoint.HTTP, http.DefaultServeMux)
}
