package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"smoke-trees/uproar/auth/database"
	"smoke-trees/uproar/posts/post"
)

type Server struct {
	db database.Database
}

var s Server

func main() {

	// Read Port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	s.db := post.NewPostDB()

	log.Info("Database Connected")

	// Router Initialization
	router := httprouter.New()

	// Routes
	router.POST("/post/:postId/upvote", UpvoteHandler)
	router.POST("/post/:postId/downvote", DownvoteHandler)

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
