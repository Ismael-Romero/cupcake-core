package rest

import (
	"context"
	"cupcake-core/src/logger"
	"cupcake-core/src/repository"
	"cupcake-core/src/services/rest/router"
	"cupcake-core/src/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(setting settings.RestService, repo repository.Repository, lrl logger.RestLogger, lhr logger.HTTPRequestLogger) {

	svr := &http.Server{
		Addr:              setting.Host,
		Handler:           router.New(lhr, repo),
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	cnxClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigint
		lrl.Info("", "Server Closed")
		if err := svr.Shutdown(context.Background()); err != nil {
			lrl.Error("", err.Error())
		}
		close(cnxClosed)
	}()

	switch setting.Https.Enable {
	case true:
		lrl.Info("", "Server HTTPs Running: ["+setting.Host+"]")
		err := svr.ListenAndServeTLS(setting.Https.Cert, setting.Https.Key)
		if err != http.ErrServerClosed {
			lrl.Error("", err.Error())
		}

	default:
		lrl.Info("", "Server HTTP Running: ["+setting.Host+"]")
		err := svr.ListenAndServe()
		if err != http.ErrServerClosed {
			lrl.Error("", err.Error())
		}
	}

	<-cnxClosed
}
