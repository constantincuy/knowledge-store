package controller

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/response"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller interface {
	Register(router *mux.Router)
}

type RouteHandler func(r *http.Request) response.Response
type MuxHandler func(w http.ResponseWriter, r *http.Request)

func Route(handler RouteHandler) MuxHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		res := handler(r)
		err := res.Send(w)
		if err != nil {
			log.Printf("Could not handle unexpected error when responding! error: %s\n", err.Error())
		}
	}
}
