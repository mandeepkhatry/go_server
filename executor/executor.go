package executor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go_server/config"
	"github.com/go_server/logger"
	"github.com/go_server/routes"

	"github.com/go_server/store"
)

var (
	timeout = 15 * time.Second
)

var Store store.DataStore

type Executor struct {
	Config *config.AppConfig
}

func NewExecutor(config *config.AppConfig) *Executor {
	return &Executor{
		Config: config,
	}
}

func (ex *Executor) Execute() {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", ex.Config.Host, ex.Config.Port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
		Handler:      routes.RegisterRoutes(),
	}

	logger.Debug("Starting API Server")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
