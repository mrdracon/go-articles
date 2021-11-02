package store_test

import (
	"flag"
	"go-articles/internal/config"
	"go-articles/pkg/log"
	"math/rand"
	"os"
	"testing"
	"time"
)

var TestCfg *config.Config

func TestMain(m *testing.M){
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", "TESTING")
	rand.Seed(time.Now().UnixNano())

	// load application configurations
	cfg, err := config.Load("../config/qa.yml", logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	TestCfg = cfg
	os.Exit(m.Run())
}
