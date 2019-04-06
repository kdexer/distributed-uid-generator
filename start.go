package main

import (
	"distributed-uid-generator/generator"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var nextId generator.NextId

func genIdHandler(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uid := nextId.GetNextId()
	fmt.Println(uid)
	fmt.Fprint(rw, uid)
}

// start up
func main() {

	dg := generator.New(1, "2019-04-05", 28, 22, 13)
	nextId = dg

	router := httprouter.New()
	router.GET("/gen/id", genIdHandler)

	startError := http.ListenAndServe(":8080", router)
	if nil != startError {
		panic("server start up error")
	}
}
