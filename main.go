package main

import (
	"cupcake-core/src/logger"
	"cupcake-core/src/repository"
	"cupcake-core/src/services/rest"
	"cupcake-core/src/settings"
	"os"
	"sync"
	"time"
)

func main() {

	settings.Launcher()
	s := settings.Loader()

	var mtx = &sync.RWMutex{}
	var RestLogger logger.RestLogger
	var HTTPRequestLogger logger.HTTPRequestLogger

	RestLogger = logger.New(
		s.GetConfigurationLogs().RestService.Path,
		s.GetConfigurationLogs().RestService.Name,
		24*time.Hour,
		mtx,
	)

	HTTPRequestLogger = logger.New(
		s.GetConfigurationLogs().HTTPRequest.Path,
		s.GetConfigurationLogs().HTTPRequest.Name,
		24*time.Hour,
		mtx,
	)

	go RestLogger.Backup()
	go HTTPRequestLogger.Backup()

	repo, err := repository.New("mysql", s.GetDatabaseEngines())
	if err != nil {
		RestLogger.Error("0", err.Error())
		os.Exit(0)
	}

	rest.Start(s.Services.Rest, repo, RestLogger, HTTPRequestLogger)
}
