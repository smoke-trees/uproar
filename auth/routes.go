package main

import (
	"github.com/smoke-trees/uproar/auth/database"
	"github.com/valyala/fasthttp"
)

func LoginHandler(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("email"))
	password := string(ctx.FormValue("password"))
	if username == "" || password == "" {
		// Send error
	}
	user := database.User{
		Username: username,
		Email:    username,
		Password: password,
	}
	_, err := s.Database.Authenticate(user)
	if err != nil {
		// Send error
	}

	//Jwt
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
