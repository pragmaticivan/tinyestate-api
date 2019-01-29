package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/pragmaticivan/tinyestate-api/healthcheck"
	log "github.com/sirupsen/logrus"
)

var build = "develop"

func main() {
	var cfg struct {
		Web struct {
			APIHost         string        `default:"0.0.0.0:3000" envconfig:"API_HOST"`
			DebugHost       string        `default:"0.0.0.0:4000" envconfig:"DEBUG_HOST"`
			ReadTimeout     time.Duration `default:"5s" envconfig:"READ_TIMEOUT"`
			WriteTimeout    time.Duration `default:"5s" envconfig:"WRITE_TIMEOUT"`
			ShutdownTimeout time.Duration `default:"5s" envconfig:"SHUTDOWN_TIMEOUT"`
		}
		Trace struct {
			Host         string        `default:"http://tracer:3002/v1/publish" envconfig:"HOST"`
			BatchSize    int           `default:"1000" envconfig:"BATCH_SIZE"`
			SendInterval time.Duration `default:"15s" envconfig:"SEND_INTERVAL"`
			SendTimeout  time.Duration `default:"500ms" envconfig:"SEND_TIMEOUT"`
		}
	}

	if err := envconfig.Process("TINYESTATE-API", &cfg); err != nil {
		log.Fatalf("main : Parsing Config : %v", err)
	}

	log.Infof("main : Started : Application Initializing version %q", build)
	defer log.Infoln("main : Completed")

	// Start DB

	// Start Debug Service

	// /debug/vars - Added to the default mux by the expvars package.
	// /debug/pprof - Added to the default mux by the net/http/pprof package.

	debug := http.Server{
		Addr:           cfg.Web.DebugHost,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    cfg.Web.ReadTimeout,
		WriteTimeout:   cfg.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Not concerned with shutting this down when the
	// application is being shutdown.
	go func() {
		log.Infof("main : Debug Listening %s", cfg.Web.DebugHost)
		log.Infof("main : Debug Listener closed : %v", debug.ListenAndServe())
	}()

	// Start API Service

	r := mux.NewRouter()
	r.HandleFunc("/_health", healthcheck.Handler).Methods("GET")
	r.Use(loggingMiddleware)

	api := http.Server{
		Addr:           cfg.Web.APIHost,
		Handler:        r,
		ReadTimeout:    cfg.Web.ReadTimeout,
		WriteTimeout:   cfg.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Infof("main : API Listening %s", cfg.Web.APIHost)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Stop API Service

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("main : Error starting server: %v", err)

	case <-osSignals:
		log.Infoln("main : Start shutdown...")

		// Create context for Shutdown call.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		if err := api.Shutdown(ctx); err != nil {
			log.Infof("main : Graceful shutdown did not complete in %v : %v", cfg.Web.ShutdownTimeout, err)
			if err := api.Close(); err != nil {
				log.Fatalf("main : Could not stop http server: %v", err)
			}
		}
	}

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Infof(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
