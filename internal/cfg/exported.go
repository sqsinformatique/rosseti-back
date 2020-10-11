package cfg

import (
	// stdlib
	"fmt"

	// other
	"github.com/vrischmann/envconfig"
)

type AppCfg struct {
	Database struct {
		DSN                   string `envconfig:"default=postgres://postgres:secret@postgres:5432/gifts"`
		Params                string `envconfig:"default=connect_timeout=10&sslmode=disable"`
		MaxIdleConnections    int    `envconfig:"default=5"`
		MaxOpenedConnections  int    `envconfig:"default=10"`
		MaxConnectionLifetime string `envconfig:"default=10s"`
	}

	Pg struct {
		Migrate struct {
			Dir    string `envconfig:"default=/internal/db/migrations/pg"`
			Action string `envconfig:"default=up"`
		}
	}

	PublicHTTP struct {
		WaitForSeconds int    `envconfig:"default=10"`
		Listen         string `envconfig:"default=0.0.0.0:9000"`
	}

	PrivateHTTP struct {
		WaitForSeconds int    `envconfig:"default=10"`
		Listen         string `envconfig:"default=0.0.0.0:9100"`
	}

	Logger struct {
		Level           string `envconfig:"default=INFO"`
		SuperVerbosive  bool   `envconfig:"default=false"`
		NoColoredOutput bool   `envconfig:"default=true"`
	}
}

// Initialize reads configuration values from environment.
func (c *AppCfg) initialize() {
	// Error here will never rise as we always provide pointer to Init().
	_ = envconfig.Init(c)
	fmt.Printf("Configuration parsed successfully")
}

func NewConfig() *AppCfg {
	cfg := &AppCfg{}
	cfg.initialize()
	return cfg
}
