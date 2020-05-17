package routers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kdexer/distributed-uid-generator/generator"
	"net/http"
)

type GeneratorRouter struct {
	generator generator.NextId
}

func NewGeneratorRouter(generator generator.NextId) *GeneratorRouter {
	return &GeneratorRouter{generator: generator}
}

func (gr GeneratorRouter) GetHandlerFunction() httprouter.Handle {
	handler := func (rw http.ResponseWriter,re *http.Request, pa httprouter.Params) {
		uid := gr.generator.GetNextId()
		fmt.Fprint(rw, uid)
	}
	return handler
}