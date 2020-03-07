package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	log "github.com/sirupsen/logrus"
	"github.com/smoke-trees/uproar/auth/database"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
)

type Server struct {
	Database database.Database
}

var s Server

func main() {

	// Read Port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Router Initialization
	router := fasthttprouter.New()

	// Routes
	router.POST("/api/register", RegisterHandler)
	router.POST("/api/login", LoginHandler)

	// Serving
	go func() {
		log.Info(fmt.Sprintf("Starting Server on port: %s", port))
		err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", "", port), router.Handler)
		if err != nil {
			log.Fatalf("Error in serving to port")
		}
	}()

	// Listen For Signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan

	// Shutdown Routine
}


