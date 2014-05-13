package main

import (
	"github.com/mix3/go-traffic-sample/model"
	"github.com/pilu/traffic"
)

var router *traffic.Router

func init() {
	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/list", ListHandler)
	router.Post("/create", CreateHandler)
	router.Post("/switch/:id", SwitchHandler)
	router.Post("/delete/:id", DeleteHandler)
	router.Post("/delete", DeleteAllHandler)
	router.ErrorHandler = ErrorHandler

	// for heroku
	if traffic.Env() == "production" {
		router.Use(traffic.NewStaticMiddleware(traffic.PublicPath()))
	}
}

func main() {
	db := model.GetDB()
	defer db.Close()

	router.Run()
}
