package agent

import (
	"time"
)

// Agent 表示一次 AI Agent 请求的完整上下文
// 类似“工作流实例”
type Agent struct {
	ID string

	// 当前状态
	State AgentState

	// 上下文信息（由 ContextManager 维护）
	Context []Message

	// 任务开始时间，用于超时判断
	StartTime time.Time

	// 最大执行时长
	Timeout time.Duration
}

// Message 表示一条上下文消息
// role 的划分是面试官会重点看的
type Message struct {
	Role    string // system / user / assistant / tool
	Content string
}

// IsTimeout 判断 Agent 是否超时
func (a *Agent) IsTimeout() bool {
	return time.Since(a.StartTime) > a.Timeout
}

// NextState 由 Orchestrator 控制
// Agent 自身不做决策，遵守单一职责
func (a *Agent) NextState(state AgentState) {
	a.State = state
}
