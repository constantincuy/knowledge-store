package rest

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/controller"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func (a Server) Run() error {
	server := &http.Server{}
	server.Handler = a.router
	server.Addr = ":8765"
	log.Println("Server is listening on port 8765")
	return server.ListenAndServe()
}

func New(knowledgeService knowledgebases.Api) Server {
	r := mux.NewRouter()
	r.StrictSlash(true)

	knowledgeController := controller.NewKnowledgeBaseController(knowledgeService)

	knowledgeController.Register(r)

	return Server{r}
}
