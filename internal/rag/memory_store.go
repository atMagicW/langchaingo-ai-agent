package rag

// MemoryStore 是最简单的向量存储实现
// 我们先验证 RAG 价值，而不是选型
type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Add(key, value string) {
	m.data[key] = value
}

func (m *MemoryStore) Search(query string) []string {
	// demo 级：直接返回全部
	// 重点是“检索 → 注入上下文”
	results := []string{}
	for _, v := range m.data {
		results = append(results, v)
	}
	return results
}
