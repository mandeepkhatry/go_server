package main

import (
	"flag"

	"github.com/go_server/config"
	"github.com/go_server/executor"
	"github.com/go_server/logger"
	"github.com/go_server/store"
)

func init() {

	//Define Log Level
	logger.NewLogger(logger.Config{
		Service: "domain",
		Level:   "debug",
	})

	//Parse Flags
	config.AppParams = &config.AppConfig{
		Host:    *flag.String("host", "127.0.0.1", "host address"),
		Port:    *flag.String("port", "5000", "port"),
		ApiBase: *flag.String("api-base", "/", "base api version"),
	}
	flag.Parse()

	logger.Debug("", logger.Field("App Params", config.AppParams))

	//Define Store
	store.Store = store.NewDataStoreClient(store.DatastoreOptions{
		StoreName: "sqlite",
		DbDir:     "data",
	})
}

func main() {
	logger.Debug("Running main")
	executor.NewExecutor(config.AppParams).Execute()
}
