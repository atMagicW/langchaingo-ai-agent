package rag

import "langchaingo-ai-agent/internal/agent"

// RAGService 将检索结果注入 Agent Context
type RAGService struct {
	store *MemoryStore
}

func NewRAGService(store *MemoryStore) *RAGService {
	return &RAGService{store: store}
}

// EnrichContext 将知识补充到上下文
func (r *RAGService) EnrichContext(a *agent.Agent, query string) {
	docs := r.store.Search(query)

	for _, doc := range docs {
		a.Context = append(a.Context, agent.Message{
			Role:    "system",
			Content: "Relevant knowledge: " + doc,
		})
	}
}
