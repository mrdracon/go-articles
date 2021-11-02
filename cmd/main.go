package main

import (
	"flag"
	"fmt"
	"go-articles/internal/article"
	"go-articles/internal/user"

	"go-articles/store"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go-articles/internal/config"
	"go-articles/pkg/log"

	"net/http"
	"os"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	rand.Seed(time.Now().UnixNano())

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	err, st := configureStore(cfg.DSN)
	if  err != nil {
		logger.Error("error configuring db: %s", err)
		os.Exit(-1)
	}

	defer st.Close()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, st, cfg),
	}

	// start the HTTP server with graceful shutdown
	//go routing.GracefulShutdown(hs, 10*time.Second, logger.Infof)
	logger.Infof("server %v is running at %v", Version, address)

	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, st *store.Store, cfg *config.Config) http.Handler {
	router := gin.Default()


	//router.Use(
	//	accesslog.Handler(logger),
	//	errors.Handler(logger),
	//	content.TypeNegotiator(content.JSON),
	//	cors.Handler(cors.AllowAll),
	//)

	//healthcheck.RegisterHandlers(router, Version)

	//rg := router.Group("/v1")
	apiV1 := router.Group("/v1")

	apiV1.POST("/users", func(c *gin.Context) {
		user.HandleCreateUser(c, st, logger)
	})
	apiV1.GET("/users/:userId", func(c *gin.Context) {
		user.HandleGetUser(c, st, logger)
	})

	apiV1.POST("/articles", func(c *gin.Context) {
		article.HandleCreateArticle(c, st, logger)
	})

	apiV1.GET("/articles/:articleID", func(c *gin.Context) {
		article.HandleGetArticle(c, st, logger)
	})

	apiV1.GET("/articles", func(c *gin.Context) {
		article.HandleListArticle(c, st, logger)
	})

	return router
}

func configureStore(DSN string) (error, *store.Store) {
	st := store.New(DSN)
	if err := st.Open(); err != nil {
		return err, nil
	}

	return nil, st
}