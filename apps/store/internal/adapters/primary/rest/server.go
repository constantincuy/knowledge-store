package rest

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/controller"
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/constantincuy/knowledgestore/internal/ports"
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

func New(knowledgeBaseManager ports.KnowledgeBaseManager, knowledgeService knowledgebases.Api, fileService files.Api) Server {
	r := mux.NewRouter()
	r.StrictSlash(true)

	knowledgeController := controller.NewKnowledgeBaseController(knowledgeBaseManager, knowledgeService, fileService)

	knowledgeController.Register(r)

	return Server{r}
}
