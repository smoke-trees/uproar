package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"smoke-trees/uproar/posts/post"
)

func UpvoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	userId := params.ByName("userId")

	//Post and user gotten from DB using ID's

	postById, userById := post.Post{}, post.User{}

	postById = post.UpdatePostOnUp(postById, userById)
}

func DownvoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	userId := params.ByName("userId")

	//Post and user gotten from DB using ID's

	postById, userById := , post.User{}

	postById = post.UpdatePostOnDown(postById, userById)

}
