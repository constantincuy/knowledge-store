package actor

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/vladopajic/go-actor/actor"
)

type Manager struct {
	actors     map[string]actor.Actor
	fileRepo   ports.FileRepo
	docService documents.Api
	embedding  ports.EmbeddingExtractor
	storage    ports.Storage
	baseRepo   ports.KnowledgeBaseRepo
}

func (m *Manager) Run(knowledgeBase string) {
	_, ex := m.actors[knowledgeBase]
	if !ex {
		m.actors[knowledgeBase] = NewKnowledgeBase(knowledgeBase, m.fileRepo, m.docService, m.embedding, m.storage)
		m.actors[knowledgeBase].Start()
	}
}

func (m *Manager) Stop(knowledgeBase string) {
	_, ex := m.actors[knowledgeBase]
	if !ex {
		m.actors[knowledgeBase].Stop()
		delete(m.actors, knowledgeBase)
	}

}

func (m *Manager) Hydrate() error {
	kbs, err := m.baseRepo.GetAll(context.Background())
	if err != nil {
		return err
	}

	for _, kb := range kbs {
		m.Run(kb)
	}

	return nil
}

func NewManager(baseRepo ports.KnowledgeBaseRepo, fileRepo ports.FileRepo, docService documents.Api, embedding ports.EmbeddingExtractor, storage ports.Storage) Manager {
	return Manager{make(map[string]actor.Actor), fileRepo, docService, embedding, storage, baseRepo}
}
