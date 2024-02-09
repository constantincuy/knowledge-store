package knowledgebase

type KnowledgeBase struct {
	Name Name
}

func New(name Name) KnowledgeBase {
	return KnowledgeBase{name}
}
