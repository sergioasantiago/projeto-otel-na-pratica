// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/telemetry"
	"google.golang.org/grpc"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	c, _ := config.LoadConfig(*configFlag)

	// Initialize telemetry
	ctx := context.Background()
	shutdown, err := telemetry.InitTelemetry(ctx, &c.Telemetry, "plans")
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown()

	// starts the gRPC server
	lis, _ := net.Listen("tcp", c.Server.Endpoint.GRPC)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	a := app.NewPlan(&c.Plans)
	a.RegisterRoutes(http.DefaultServeMux, grpcServer)

	go func() {
		_ = grpcServer.Serve(lis)
	}()

	_ = http.ListenAndServe(c.Server.Endpoint.HTTP, http.DefaultServeMux)
}
