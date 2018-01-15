package config

import (
	"log"
)

// Shared is a shared AppConfig instance for reference
// throughout the application
var Shared *AppConfig

func init() {
	Shared = &AppConfig{
		DB: &DBConfig{
			Dialect:  "mysql",
			username: "root",
			password: "",
			name:     "go_graph_complex",
			charset:  "utf8",
		},
		Crypto: getCryptoConfig(),
	}
}

func fatalConfigError(errors ...error) {
	for _, err := range errors {
		if err != nil {
			log.Fatalf("Config error: %v", err)
		}
	}
}
