package testutils

import (
	"time"

	"github.com/spf13/viper"
	"github.com/trenova-app/transport/internal/pkg/config"
)

func NewTestConfig() *config.Config {
	return &config.Config{
		App: config.AppConfig{
			Name:        "transport-test",
			Environment: "testing",
			LogLevel:    "error",
			Version:     "test",
		},
		DB: config.DatabaseConfig{
			Driver:          config.DatabaseDriverPostgres,
			Host:            "localhost", // Will be overridden by container
			Port:            5432,        // Will be overridden by container
			Username:        "postgres",  // Use standard postgres user for tests
			Password:        "postgres",  // Use simple password for tests
			SSLMode:         "disable",
			MaxConnections:  10,
			MaxIdleConns:    10,
			ConnMaxLifetime: 3600,
			Debug:           true,
		},
		Audit: config.AuditConfig{
			FlushInterval: 10,
			BufferSize:    1000,
		},
		Minio: config.MinioConfig{
			Endpoint:  "localhost:9000",
			AccessKey: "MBrqCQfiMVaO3jOHcAAs",
			SecretKey: "8uk1hH8PRuiFfxp4focyg3bjBfGo6s0EDFgLosTN",
			UseSSL:    false,
		},
		Cors: config.CorsConfig{
			AllowedOrigins:   "https://localhost:5173, http://localhost:5173, https://localhost:4173, http://localhost:4173, https://localhost:3000, http://localhost:3000, http://trenova.local, https://trenova.local",
			AllowedHeaders:   "Authorization, Origin, Content-Type, Accept, X-CSRF-Token, X-Idempotency-Key, Set-Cookie, Cookie, X-Forwarded-For, CF-Connecting-IP, X-Request-ID",
			AllowedMethods:   "GET, POST, PUT, DELETE, OPTIONS",
			AllowCredentials: true,
			MaxAge:           300,
		},
		Search: config.SearchConfig{
			Host:          "http://localhost:7700",
			APIKey:        "ufp2vpm0CRRCpTY-IxpiwLPQNLBKIlGYxI5vNHHbCMQ",
			IndexPrefix:   "trv",
			MaxBatchSize:  100,
			BatchInterval: 30,
		},
		Server: config.ServerConfig{
			SecretKey:               "test",
			ListenAddress:           ":3001",
			ReadTimeout:             30,
			WriteTimeout:            30,
			IdleTimeout:             60,
			PassLocalsToViews:       false,
			ReadBufferSize:          16384,
			WriteBufferSize:         16384,
			EnablePrefork:           false,
			EnablePrintRoutes:       false,
			Concurrency:             262144,
			DisableStartupMessage:   true,
			StreamRequestBody:       true,
			StrictRouting:           false,
			CaseSensitive:           true,
			Immutable:               false,
			EnableIPValidation:      true,
			EnableTrustedProxyCheck: true,
			ProxyHeader:             "X-Forwarded-For",
		},
		Auth: config.AuthConfig{
			SessionCookieName: "trv-session-id",
			CookiePath:        "/",
			CookieDomain:      "",
			CookieHTTPOnly:    false,
			CookieSecure:      false,
			CookieSameSite:    "Lax",
		},
		Redis: config.RedisConfig{
			Addr:         "localhost:6379",
			Password:     "",
			DB:           0,
			ConnTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			PoolSize:     10,
			MinIdleConns: 10,
		},
	}
}

func NewTestConfigManager(cfg *config.Config) *config.Manager {
	return &config.Manager{
		Cfg:   cfg,
		Viper: viper.New(),
	}
}