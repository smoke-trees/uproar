package main

import (
	"encoding/json"
	"github.com/pascaldekloe/jwt"
	"github.com/smoke-trees/uproar/auth/database"
	"github.com/valyala/fasthttp"
)

type RegisterResponse struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
}

func LoginHandler(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("email"))
	password := string(ctx.FormValue("password"))
	if username == "" || password == "" {
		authResponse := database.AuthResponse{
			Status: database.WrongUsername,
			JWT:    nil,
		}
		js, _ := json.Marshal(&authResponse)
		ctx.Write(js)
		return

	}
	user := database.User{
		Username: username,
		Email:    username,
		Password: password,
	}
	authResponse, err := s.Database.Authenticate(user)
	if err != nil {
		js, _ := json.Marshal(&authResponse)
		ctx.Write(js)
		return
	}
	claims := jwt.Claims{Set: map[string]interface{}{
		"email":    user.Email,
		"username": user.Username,
		"password": user.Password,
	}}
	jwToken, err1 := claims.HMACSign("SHA256", []byte("smoketrees"))
	if err1 != nil {
		authResponse.Status = 5
		return
	}
	authResponse.JWT = jwToken

	js, _ := json.Marshal(&authResponse)
	ctx.Write(js)
}

func RegisterHandler(ctx *fasthttp.RequestCtx) {
	email := string(ctx.FormValue("email"))
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))
	if username == "" || password == "" || email == "" {
		rr := RegisterResponse{
			Status:  1,
			Message: "all fields not mentioned",
		}
		js, _ := json.Marshal(&rr)
		ctx.Write(js)
	}
	user := database.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	err := s.Database.AddUser(user)

	if err != nil {
		// Send error
		rr := RegisterResponse{
			Status:  2,
			Message: err.Error(),
		}
		js, _ := json.Marshal(&rr)
		ctx.Write(js)
	}
	rr := RegisterResponse{
		Status:  0,
		Message: "success",
	}
	js, _ := json.Marshal(&rr)
	ctx.Write(js)
}
