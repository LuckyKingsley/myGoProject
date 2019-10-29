package mvcMain

import (
	"fmt"
	"geometry/mvc/controller"
	"geometry/mvc/framework"
	"net/http"
	"time"
)

func MvcMain() {

	framework.InitDB()
	framework.CreateTable()

	server := &http.Server{
		Addr:        ":8080",
		Handler:     framework.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(framework.Router)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("start server error")
	}
	fmt.Println("start server success")
}

func RegiterRouter(handler *framework.RouterHandler) {
	new(controller.UserConterller).Router(handler)
}
