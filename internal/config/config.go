package config

import (
	"net/url"

	"github.com/tagirmukail/forum/internal/utils/env"
)

type Config struct {
	LogLevel string
	API
	Postgres
}

type API struct {
	Addr         string
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
}

type Postgres struct {
	Username string
	Password string
	Hostname string
	Port     string
	Database string
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.LogLevel = env.ResolveString("LOG_LEVEL", "debug")

	cfg.API.Addr = env.ResolveString("API_ADDR", ":8080")
	cfg.API.WriteTimeout = env.ResolveInt("API_W_TIMEOUT", 15)
	cfg.API.ReadTimeout = env.ResolveInt("API_R_TIMEOUT", 15)
	cfg.API.IdleTimeout = env.ResolveInt("API_I_TIMEOUT", 15)

	cfg.Postgres.Username = env.ResolveString("POSTGRES_USERNAME", "postgres")
	cfg.Postgres.Password = env.ResolveString("POSTGRES_PASSWORD", "postgres")
	cfg.Postgres.Hostname = env.ResolveString("POSTGRES_HOSTNAME", "localhost")
	cfg.Postgres.Port = env.ResolveString("POSTGRES_PORT", "5432")
	cfg.Postgres.Database = env.ResolveString("POSTGRES_DATABASE", "postgres")

	return cfg
}

func (p *Postgres) ConnectionURI() string {
	if p == nil {
		return ""
	}

	host := p.Hostname
	if p.Port != "" {
		host += ":" + p.Port
	}

	var userinfo *url.Userinfo
	if p.Username != "" || p.Password != "" {
		userinfo = url.UserPassword(p.Username, p.Password)
	}

	uri := &url.URL{
		Scheme: "postgres",
		User:   userinfo,
		Host:   host,
		Path:   p.Database,
	}

	query := url.Values{"sslmode": []string{"disable"}}

	uri.RawQuery = query.Encode()

	return uri.String()
}
