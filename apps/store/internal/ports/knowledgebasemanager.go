package ports

type KnowledgeBaseManager interface {
	Run(knowledgeBase string)
	Stop(knowledgeBase string)
	Hydrate() error
}
