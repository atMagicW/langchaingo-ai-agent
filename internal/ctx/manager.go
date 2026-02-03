package ctx

import (
	"langchaingo-ai-agent/internal/agent"
)

// Manager 负责构建和维护 Agent 的上下文
// ⚠️ 重要设计点：
// Orchestrator 不关心“上下文怎么来”
// 它只关心“状态流转”
type Manager struct {
	MaxHistory int
}

// NewManager 创建 ContextManager
func NewManager() *Manager {
	return &Manager{
		MaxHistory: 10, // 控制 token 的第一道闸
	}
}

// Build 在 THINKING 阶段构建 prompt
func (m *Manager) Build(a *agent.Agent) {
	// 1. system prompt（固定规则）
	system := agent.Message{
		Role:    "system",
		Content: "You are a task-oriented AI agent.",
	}

	// 2. 截断历史上下文（滑动窗口）
	history := m.trimHistory(a.Context)

	// 3. 重新组装
	a.Context = []agent.Message{system}
	a.Context = append(a.Context, history...)
}

// AppendAssistant 写入 LLM 返回结果
func (m *Manager) AppendAssistant(a *agent.Agent, content string) {
	a.Context = append(a.Context, agent.Message{
		Role:    "assistant",
		Content: content,
	})
}

// NeedExecuteTask 判断是否需要执行工具
// 为什么不是 Orchestrator 判断？
// 决策来自 LLM 语义
func (m *Manager) NeedExecuteTask(agent *agent.Agent) bool {
	if len(agent.Context) == 0 {
		return false
	}
	last := agent.Context[len(agent.Context)-1]
	return containsToolCall(last.Content)
}

// UpdateAfterTask 将工具执行结果写回上下文
func (m *Manager) UpdateAfterTask(a *agent.Agent) {
	a.Context = append(a.Context, agent.Message{
		Role:    "tool",
		Content: "task execution result",
	})
}

// trimHistory 控制历史长度
func (m *Manager) trimHistory(ctx []agent.Message) []agent.Message {
	if len(ctx) <= m.MaxHistory {
		return ctx
	}
	return ctx[len(ctx)-m.MaxHistory:]
}

// 简化判断（真实项目可用 JSON schema）
func containsToolCall(content string) bool {
	return len(content) > 0 && content[0] == '['
}
