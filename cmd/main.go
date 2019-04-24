package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	payments "github.com/vstoianovici/paymentsapi"
	config "github.com/vstoianovici/paymentsapi/config"
)

func main() {

	// Define a cutom logger with the specified format
	logger := createLogger()
	startLogger := log.With(logger, "tag", "start")
	startLogger.Log("msg", "created logger")

	// define channel to monitor signals from os and handle gracefully any kind of shutdown
	var gracefulStopC = make(chan os.Signal)
	signal.Notify(gracefulStopC, syscall.SIGKILL)
	signal.Notify(gracefulStopC, syscall.SIGINT)
	signal.Notify(gracefulStopC, syscall.SIGQUIT)
	signal.Notify(gracefulStopC, syscall.SIGTERM)
	signal.Notify(gracefulStopC, syscall.SIGHUP)

	// define a monitor channel for http server errors
	monC := make(chan error)

	// get the postgres DB config and the application port number (can be passed in command line)
	dbConfigFile, appPort := config.ParseArgs()

	// create a new postgres DB connection
	db, err := payments.NewDBConnection(dbConfigFile)
	if err != nil {
		m, _ := fmt.Println("error when connecting to postgres:", err)
		startLogger.Log("err", m)
		os.Exit(0)
	}
	// defer closing postgres DB eventually
	defer payments.CloseDB(db)

	// make sure the right tables exist in the database
	payments.MigrateDB(db)

	// create a new Payments API service
	svc := payments.NewPaymentService(db)

	// add validator service
	svc, err = payments.NewValidator(svc)
	if err != nil {
		startLogger.Log("err", err)
		os.Exit(0)
	}

	// add a layer of logging on top of the core wallet service
	svc = payments.NewLogging(logger, svc)

	// create string for server address
	port := ":" + strconv.Itoa(appPort)

	// create a router
	router := payments.NewHTTPTransport(svc)

	// define http server
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	startLogger.Log("msg", "Welcome to the 'Payments REST API'")
	startLogger.Log("msg", "Payments API Endpoint: http://127.0.0.1"+port+"/v1/payments/")
	startLogger.Log("msg", "HTTP serving locally...", "port", port)

	// launch server in a go routine
	go func() {
		monC <- server.ListenAndServe()
	}()

	// ...and wait for either a failure to launch on monC or a signal to trigger the server shutdown on gracefulStopC
	select {
	// try to shutdown gracefully upon receiving defined signal
	case sig := <-gracefulStopC:
		m := fmt.Sprintf("Caught sig: %+v ", sig)
		startLogger.Log("msg", m)
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			m := fmt.Sprintf("HTTP server could not be stopped gracefully: %v", err)
			startLogger.Log("err", m)
			if err := server.Close(); err != nil {
				m := fmt.Sprintf("HTTP server could not be stopped: %v", err)
				startLogger.Log("err", m)
			}
		} else {
			startLogger.Log("msg", "Gracefully shutdown HTTP server.")
		}
		os.Exit(0)

	// monitor http server launch errors
	case err := <-monC:
		m, _ := fmt.Printf("err: Error starting server: %v ", err)
		startLogger.Log("err", m)
	}
}

// createLogger implements the disred log format
func createLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	return log.With(logger, "time", log.DefaultTimestampUTC())
}
