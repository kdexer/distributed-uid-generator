package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
  project start up method
 */
func main() {
	router := httprouter.New()
	startError := http.ListenAndServe(":8080", router)
	if nil != startError {
		panic("server start up error")
	}
}
