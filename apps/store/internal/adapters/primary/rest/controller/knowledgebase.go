package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/response"
	"github.com/constantincuy/knowledgestore/internal/core/service"
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type KnowledgeBaseController struct {
	service              knowledgebases.Api
	fileService          files.Api
	knowledgeBaseManager ports.KnowledgeBaseManager
}

type CreateKnowledgeBasePayload struct {
	Name string `json:"name"`
}

func (c KnowledgeBaseController) CreateKnowledgeBase(r *http.Request) response.Response {
	var createReq CreateKnowledgeBasePayload
	err := json.NewDecoder(r.Body).Decode(&createReq)

	if err != nil {
		return response.FromError(err)
	}

	_, err = c.service.Create(context.Background(), knowledgebases.CreateKnowledgeBaseReq{Name: createReq.Name})

	if err != nil {
		return response.FromError(err)
	}

	c.knowledgeBaseManager.Run(createReq.Name)
	return response.New()
}

func (c KnowledgeBaseController) SearchKnowledgeBase(r *http.Request) response.Response {
	vars := mux.Vars(r)
	base, ex := vars["base"]

	if !ex {
		return response.FromError(service.NewError(service.ErrBadRequest, errors.New("no knowledge base specified")))
	}

	query := r.URL.Query().Get("q")

	if strings.Trim(query, " \r\n") == "" {
		return response.FromError(service.NewError(service.ErrBadRequest, errors.New("no query specified")))
	}

	fis, err := c.fileService.Search(context.Background(), files.SearchFilesReq{KnowledgeBase: base, Query: query})

	if err != nil {
		return response.FromError(err)
	}

	return response.New().Json(fis)
}

func (c KnowledgeBaseController) Register(router *mux.Router) {
	router.HandleFunc("/knowledge-base/", Route(c.CreateKnowledgeBase)).Methods("POST")
	router.HandleFunc("/knowledge-base/{base}/files", Route(c.SearchKnowledgeBase)).Methods("GET")

}

func NewKnowledgeBaseController(knowledgeBaseManager ports.KnowledgeBaseManager, service knowledgebases.Api, fileService files.Api) KnowledgeBaseController {
	return KnowledgeBaseController{service, fileService, knowledgeBaseManager}
}
