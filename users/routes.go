package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/smoke-trees/uproar/users/user"
	"net/http"
)

func UserRegisterHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var body []byte
	request.Body.Read(body)
	var u user.User
	json.Unmarshal(body, u)
	fmt.Print(u)
}

func UserLoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func UserDataHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func UserPostDownvoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func UserPostUpvoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func UserNewPostHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
