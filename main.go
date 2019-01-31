package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pragmaticivan/tinyestate-api/adapters/web"
	_cityRepository "github.com/pragmaticivan/tinyestate-api/city/repository"
	_cityUsecase "github.com/pragmaticivan/tinyestate-api/city/usecase"
	"github.com/pragmaticivan/tinyestate-api/schema"
	_stateRepository "github.com/pragmaticivan/tinyestate-api/state/repository"
	_stateUsecase "github.com/pragmaticivan/tinyestate-api/state/usecase"
	log "github.com/sirupsen/logrus"
)

var build = "develop"

func main() {

	flag.Parse()

	var cfg struct {
		Web struct {
			APIPort         string        `default:"3000" envconfig:"PORT"`
			DebugPort       string        `default:"4000" envconfig:"DEBUG_PORT"`
			APIHost         string        `default:"0.0.0.0" envconfig:"API_HOST"`
			DebugHost       string        `default:"0.0.0.0" envconfig:"DEBUG_HOST"`
			ReadTimeout     time.Duration `default:"5s" envconfig:"READ_TIMEOUT"`
			WriteTimeout    time.Duration `default:"5s" envconfig:"WRITE_TIMEOUT"`
			ShutdownTimeout time.Duration `default:"5s" envconfig:"SHUTDOWN_TIMEOUT"`
		}
		DB struct {
			DBHost string `default:"localhost" envconfig:"DB_HOST"`
			DBPort string `default:"5432" envconfig:"DB_PORT"`
			DBUser string `default:"tinyestate" envconfig:"DB_USER"`
			DBPass string `default:"" envconfig:"DB_PASS"`
			DBName string `default:"tinyestate" envconfig:"DB_NAME"`
		}
	}

	if err := envconfig.Process("TINYESTATE-API", &cfg); err != nil {
		log.Fatalf("main : Parsing Config : %v", err)
	}

	log.Infof("main : Started : Application Initializing version %q", build)
	defer log.Infoln("main : Completed")

	// Start DB
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, cfg.DB.DBPass, cfg.DB.DBName)

	dbConn, err := sql.Open("postgres", connection)

	if err != nil {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Start migration or seed
	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(dbConn); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return

	case "seed":
		if err := schema.Seed(dbConn); err != nil {
			log.Println("error seeding database", err)
			os.Exit(1)
		}
		log.Println("Seed data complete")
		return
	}

	// Start Debug Service

	// /debug/vars - Added to the default mux by the expvars package.
	// /debug/pprof - Added to the default mux by the net/http/pprof package.

	debug := http.Server{
		Addr:           cfg.Web.DebugHost + ":" + cfg.Web.DebugPort,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    cfg.Web.ReadTimeout,
		WriteTimeout:   cfg.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Not concerned with shutting this down when the
	// application is being shutdown.
	go func() {
		log.Infof("main : Debug Listening %s", cfg.Web.DebugHost+":"+cfg.Web.DebugPort)
		log.Infof("main : Debug Listener closed : %v", debug.ListenAndServe())
	}()

	// Temporarely load dependencies here
	timeoutContext := time.Duration(10000 * 5)
	cityRepository := _cityRepository.NewPostgresCityRepository(dbConn)
	stateRepository := _stateRepository.NewPostgresStateRepository(dbConn)
	stateUsecase := _stateUsecase.NewStateUsecase(stateRepository, timeoutContext)
	cityUsecase := _cityUsecase.NewCityUsecase(cityRepository, timeoutContext)

	// Start API Service
	r := web.NewWebAdapter(stateUsecase, cityUsecase)

	api := http.Server{
		Addr:           cfg.Web.APIHost + ":" + cfg.Web.APIPort,
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
		log.Infof("main : API Listening %s", cfg.Web.APIHost+":"+cfg.Web.APIPort)
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
