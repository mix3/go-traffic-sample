package main

import (
	"github.com/mix3/go-traffic-sample/model"
	"github.com/naoina/genmai"
	"github.com/pilu/traffic"
)

var router *traffic.Router
var db *genmai.DB

func init() {
	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/list", ListHandler)
	router.Post("/create", CreateHandler)
	router.Post("/switch/:id", SwitchHandler)
	router.Post("/delete/:id", DeleteHandler)
	router.ErrorHandler = ErrorHandler

	db = model.DBGet()
}

func main() {
	router.Run()
}
