package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lab259/cors"
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
		port = "3001"
	}

	// Connect to database
	db, err := database.NewAuthDB("mongodb://localhost:27017", "auth")
	if err != nil {
		log.Fatal("Error in connecting to database")
	}
	s.Database = db
	log.Info("Database Connected")

	// Router Initialization
	router := fasthttprouter.New()

	// Routes
	router.POST("/register", RegisterHandler)
	router.POST("/login", LoginHandler)

	// CORS
	handler := cors.Default().Handler(router.Handler)

	// Serving
	go func() {
		log.Info(fmt.Sprintf("Starting Server on port: %s", port))
		err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", "", port), handler)
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
