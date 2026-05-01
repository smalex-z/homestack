package config_test

import (
	"testing"

	"homestack/internal/config"
)

func TestLoadDefaultsWhenEnvUnset(t *testing.T) {
	// Not parallel: mutates process env. Each var is unset before Load so the
	// fallback branch in getEnv runs deterministically.
	for _, k := range []string{"PORT", "DB_PATH", "CORS_ORIGIN", "APP_ENV"} {
		t.Setenv(k, "")
	}
	cfg := config.Load()
	if cfg.Port != "8080" {
		t.Errorf("Port = %q, want 8080", cfg.Port)
	}
	if cfg.DBPath != "./homestack.db" {
		t.Errorf("DBPath = %q, want ./homestack.db", cfg.DBPath)
	}
	if cfg.CORSOrigin != "*" {
		t.Errorf("CORSOrigin = %q, want *", cfg.CORSOrigin)
	}
	if cfg.Env != "production" {
		t.Errorf("Env = %q, want production", cfg.Env)
	}
}

func TestLoadOverridesFromEnv(t *testing.T) {
	t.Setenv("PORT", "9000")
	t.Setenv("DB_PATH", "/var/lib/homestack/x.db")
	t.Setenv("CORS_ORIGIN", "https://example.com")
	t.Setenv("APP_ENV", "staging")

	cfg := config.Load()
	if cfg.Port != "9000" {
		t.Errorf("Port = %q, want 9000", cfg.Port)
	}
	if cfg.DBPath != "/var/lib/homestack/x.db" {
		t.Errorf("DBPath = %q, want override", cfg.DBPath)
	}
	if cfg.CORSOrigin != "https://example.com" {
		t.Errorf("CORSOrigin = %q, want override", cfg.CORSOrigin)
	}
	if cfg.Env != "staging" {
		t.Errorf("Env = %q, want staging", cfg.Env)
	}
}
