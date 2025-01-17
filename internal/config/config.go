// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Payments      Payments        `yaml:"payments"`
	Subscriptions Subscriptions   `yaml:"subscriptions"`
	Plans         Plans           `yaml:"plans"`
	Users         Users           `yaml:"users"`
	Server        Server          `yaml:"server"`
	Telemetry     TelemetryConfig `yaml:"telemetry"`
}

type Server struct {
	Endpoint Endpoint `yaml:"endpoint"`
}

type Endpoint struct {
	GRPC string `yaml:"grpc"`
	HTTP string `yaml:"http"`
}

type Payments struct {
	SubscriptionsEndpoint string  `yaml:"subscriptions_endpoint"`
	SQLLite               SQLLite `yaml:"sqlite"`
	NATS                  NATS    `yaml:"nats"`
}

type NATS struct {
	Endpoint     string `yaml:"endpoint"`
	Subject      string `yaml:"subject"`
	Stream       string `yaml:"stream"`
	ConsumerName string `yaml:"consumer_name"`
}

type SQLLite struct {
	DSN string `yaml:"dsn"`
}

type Subscriptions struct {
	UsersEndpoint string `yaml:"users_endpoint"`
	PlansEndpoint string `yaml:"plans_endpoint"`
}

type Plans struct {
}

type Users struct {
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	cfg := getDefaultConfig()

	if filename == "" {
		return cfg, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Replace environment variable placeholders
	data = []byte(os.ExpandEnv(string(data)))

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}

func getDefaultConfig() *Config {
	return &Config{
		Payments: Payments{
			SubscriptionsEndpoint: "http://localhost:8080/subscriptions",
			SQLLite: SQLLite{
				DSN: "file::memory:?cache=shared",
			},
			NATS: NATS{
				Endpoint:     "nats://localhost:4222",
				Subject:      "payment.process",
				Stream:       "payments",
				ConsumerName: "payments",
			},
		},
		Subscriptions: Subscriptions{
			UsersEndpoint: "http://localhost:8080/users",
			PlansEndpoint: "http://localhost:8080/plans",
		},
		Plans: Plans{},
		Users: Users{},
		Server: Server{
			Endpoint: Endpoint{
				GRPC: ":8081",
				HTTP: ":8080",
			},
		},
	}
}
