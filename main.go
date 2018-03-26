package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"yogo_rest/config"
	"yogo_rest/rest_api"
	"github.com/ricardo-ch/go-logger"
	tracing "github.com/ricardo-ch/go-tracing"
)

const appName = "yogo-rest"

func init() {
	// initialization (optional)
	logger.InitLogger(false)
}

func main() {

	//Zipkin Connection
	tracing.SetGlobalTracer(appName, config.SvcTracingZipkin)
	defer tracing.FlushCollector()

	// Errors channel
	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// rest_api endpoint
	rest_apiService := rest_api.NewService(rest_api.NewRepository(nil))
	rest_apiService = rest_api.NewTracing(rest_apiService)
	rest_apiHandler := rest_api.NewHandler(rest_apiService)

	go func() {

		httpAddr := ":" + config.AppPort
		router := mux.NewRouter()

		// index endpoint
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome to the demo-yogo API!")
		})

		// healthz endpoint
		router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
		})

		router.Handle("/rest_api/", tracing.HTTPMiddleware("rest_api-handler", http.HandlerFunc(rest_apiHandler.Get)))

		httpServer := &http.Server{
			Addr:    httpAddr,
			Handler: router,
		}

		logger.Info(fmt.Sprintf("The microservice demo-yogo is started on port %s", config.AppPort), zap.String("port", config.AppPort))
		errc <- httpServer.ListenAndServe()

	}()

	logger.Error("exit", zap.Error(<-errc))
}
