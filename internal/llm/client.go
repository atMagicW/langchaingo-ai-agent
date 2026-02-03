package llm

import "context"

// Client 抽象 LLM 能力
// 为什么要抽接口？
// 1. 多模型支持
// 2. fallback
// 3. 测试可 mock
type Client interface {
	Generate(ctx context.Context, messages []Message) (string, error)
}

// Message 与 Agent Context 解耦
type Message struct {
	Role    string
	Content string
}
