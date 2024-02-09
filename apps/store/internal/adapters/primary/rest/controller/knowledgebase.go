package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest/response"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrDuplicate = errors.New("duplicate entry")
)

type KnowledgeBaseController struct {
	service knowledgebases.Api
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

	return response.New()
}

func (c KnowledgeBaseController) Register(router *mux.Router) {
	router.HandleFunc("/knowledge-base/", Route(c.CreateKnowledgeBase)).Methods("POST")

}

func NewKnowledgeBaseController(service knowledgebases.Api) KnowledgeBaseController {
	return KnowledgeBaseController{service}
}
