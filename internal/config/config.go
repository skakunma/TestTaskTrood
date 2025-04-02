package config

import (
	"errors"
	"os"

	"github.com/skakunma/TestTaskTrood/internal/storage"
	"go.uber.org/zap"
)

var (
	ErrEnvServerAddr = errors.New("environment variable SERVER_ADDRESS not found")
	ErrEnvDbUrl      = errors.New("environment variable DATABASE_URL not found")
	ErrEnvSpay       = errors.New("environment variables HOST_SPACY and PORT_SPACY not found")
)

type Config struct {
	Store         storage.Storage
	AddrRunServer string
	DatabaseURL   string
	SpacyHost     string
	Sugar         *zap.SugaredLogger
}

func LoadEnvs(cfg *Config) error {
	addrRunServer := os.Getenv("SERVER_ADDRESS")
	databaseURL := os.Getenv("DATABASE_URL")
	spacyHost := os.Getenv("HOST_SPACY")
	spacyPort := os.Getenv("PORT_SPACY")

	if addrRunServer == "" {
		return ErrEnvServerAddr
	}

	if databaseURL == "" {
		return ErrEnvDbUrl
	}

	if spacyHost == "" || spacyPort == "" {
		return ErrEnvSpay
	}

	cfg.AddrRunServer = addrRunServer
	cfg.DatabaseURL = databaseURL
	cfg.SpacyHost = "http://" + spacyHost + ":" + spacyPort

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := LoadEnvs(cfg)

	if err != nil {
		return nil, err
	}

	store, err := storage.NewPostgresStorage(cfg.DatabaseURL)

	if err != nil {
		return nil, err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	cfg.Sugar = logger.Sugar()

	cfg.Store = store

	return cfg, nil
}
