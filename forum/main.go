package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/smoke-trees/uproar/forum/forum"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	Database forum.Database
}

var s Server

func main() {

	// Read Port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	var err error
	s.Database, err = forum.NewForumDB("mongodb://localhost:27017", "forum")
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	log.Info("Connected to Database")
	// Connect to database

	// Router Initialization
	router := httprouter.New()

	// Routes
	router.POST("/forum/register", UserRegisterHandler)
	router.POST("/forum/post/upvote", UserPostUpVoteHandler)
	router.POST("/forum/post/downvote", UserPostDownVoteHandler)
	router.GET("/forum/data/user", UserDataHandler)

	// CORS
	handler := cors.Default().Handler(router)

	// Serving
	go func() {
		log.Info(fmt.Sprintf("Starting Server on port: %s", port))
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", port), handler)
		if err != nil {
			log.Fatalf("Error in serving to port")
		}
	}()
	log.Info("Started the server")

	// Listen For Signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan

	// Shutdown Routine
}
