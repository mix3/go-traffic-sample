package main

import (
	"fmt"
	"strconv"

	"github.com/mix3/go-traffic-sample/model"

	"github.com/pilu/traffic"
)

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func ErrorHandler(w traffic.ResponseWriter, r *traffic.Request, err interface{}) {
	w.WriteJSON(&ResponseData{
		Success: false,
		Message: fmt.Sprintf("%+v", err),
		Result:  nil,
	})
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Render("index")
}

func ListHandler(w traffic.ResponseWriter, r *traffic.Request) {
	result, err := model.TodoList(db)
	if err != nil {
		panic(err)
	}
	w.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func CreateHandler(w traffic.ResponseWriter, r *traffic.Request) {
	title := r.FormValue("title")
	if err := model.TodoCreate(db, title); err != nil {
		panic(err)
	}
	result, err := model.TodoList(db)
	if err != nil {
		panic(err)
	}
	w.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func SwitchHandler(w traffic.ResponseWriter, r *traffic.Request) {
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	if err = model.TodoSwitch(db, id); err != nil {
		panic(err)
	}
	result, err := model.TodoList(db)
	if err != nil {
		panic(err)
	}
	w.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func DeleteHandler(w traffic.ResponseWriter, r *traffic.Request) {
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	if err = model.TodoDelete(db, id); err != nil {
		panic(err)
	}
	result, err := model.TodoList(db)
	if err != nil {
		panic(err)
	}
	w.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}
