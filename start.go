package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// start up
func main() {
	router := httprouter.New()
	startError := http.ListenAndServe(":8080", router)
	if nil != startError {
		panic("server start up error")
	}
}
