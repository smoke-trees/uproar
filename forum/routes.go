package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/smoke-trees/uproar/forum/forum"
	"net/http"
)

type ResponseStatus int8

const SuccessMessage string = "success"

const (
	//StatusSuccess
	StatusSuccess ResponseStatus = iota
	StatusFail
)

type Response struct {
	Status  ResponseStatus `json:"status"`
	Message interface{}    `json:"msg"`
}

func UserRegisterHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	u := forum.User{
		UserId:    request.FormValue("userId"),
		Name:      request.FormValue("name"),
		UserName:  request.FormValue("username"),
		Phone:     request.FormValue("phone"),
		Email:     request.FormValue("email"),
		Address1:  request.FormValue("ad1"),
		Address2:  request.FormValue("ad2"),
		City:      request.FormValue("city"),
		State:     request.FormValue("State"),
		UserLevel: 0,
		Cred:      0.5,
		RelUp:     nil,
		RelDown:   nil,
		Posts:     nil,
	}
	err := s.Database.NewUserRegister(u)
	var p []byte
	if err != nil {
		p, _ = json.Marshal(&Response{
			Status:  1,
			Message: err.Error(),
		})
	} else {
		p, _ = json.Marshal(&Response{
			Status:  0,
			Message: SuccessMessage,
		})
	}
	writer.Write(p)
}

func UserDataHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := request.URL.Query().Get("id")
	u, err := s.Database.GetUserFromUserId(id)
	var p []byte
	if err != nil {
		p, _ = json.Marshal(&Response{
			Status:  1,
			Message: err.Error(),
		})
	} else {
		p, _ = json.Marshal(&Response{
			Status:  0,
			Message: u,
		})
	}
	writer.Write(p)

}

func UserPostDownVoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
 panic("implement")
}

func UserPostUpVoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	panic("implement")
}

func UserNewPostHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
