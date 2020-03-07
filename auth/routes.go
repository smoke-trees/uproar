package main

import (
	"encoding/json"
	"github.com/pascaldekloe/jwt"
	"github.com/smoke-trees/uproar/auth/database"
	"github.com/valyala/fasthttp"
)

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
	jwt, err1 := claims.HMACSign("SHA256", []byte("smoketrees"))
	if err1 != nil {
		// Add response for no sign
		return
	}
	authResponse.JWT = jwt

	js, _ := json.Marshal(&authResponse)
	ctx.Write(js)
}

func RegisterHandler(ctx *fasthttp.RequestCtx) {
	email := string(ctx.FormValue("email"))
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))
	if username == "" || password == "" || email == "" {
		// Send error
	}
	user := database.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	err := s.Database.AddUser(user)

	if err != nil {
		// Send error
	}

	// Send the Respond
}
