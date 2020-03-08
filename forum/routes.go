package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
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

func PostDownVoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PostUpVoteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func NewPostHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := request.FormValue("username")
	postContent := request.FormValue("postContent")

	u, _ := s.Database.GetUserFromUserName(username)

	sha := sha256.New()
	sha.Write([]byte(postContent))

	post := forum.Post{
		PostId:      hex.EncodeToString(sha.Sum(nil)),
		UserId:      u.UserId,
		PostContent: postContent,
		PostUp:      0,
		PostDown:    0,
		Rel:         u.Cred,
	}
	uPost := forum.UserPost{
		PostId: hex.EncodeToString(sha.Sum(nil)),
		Rel:    u.Cred,
	}

	err := s.Database.AddUserPost(uPost, u)
	if err != nil {
		log.Warn(err)
		res, _ := json.Marshal(&Response{
			Status:  StatusFail,
			Message: err.Error(),
		})
		writer.Write(res)
		return
	}

	err = s.Database.NewPost(post)
	if err != nil {
		log.Warn(err)
		res, _ := json.Marshal(&Response{
			Status:  StatusFail,
			Message: err.Error(),
		})
		writer.Write(res)
	}

	res, _ := json.Marshal(&Response{
		Status:  StatusSuccess,
		Message: SuccessMessage,
	})
	writer.Write(res)
}
