package rest

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/controller"
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *mux.Router
}

func (a Server) Run() error {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("EMBD_STORE_ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server := &http.Server{}
	server.Handler = handlers.CORS(originsOk, headersOk, methodsOk)(a.router)
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
