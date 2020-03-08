package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
}

var s Server

func main() {

	// Read Port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Connect to database

	// Router Initialization
	router := httprouter.New()

	// Routes
	router.POST("/user/register", UserRegisterHandler)
	router.POST("/user/login", UserLoginHandler)
	router.POST("/user/post/upvote", UserPostUpvoteHandler)
	router.POST("/user/post/downvote", UserPostDownvoteHandler)
	router.POST("/user/:userId", UserDataHandler)

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

	// Listen For Signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan

	// Shutdown Routine
}


