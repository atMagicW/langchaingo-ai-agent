package task

import (
	"context"
	"errors"
	"langchaingo-ai-agent/internal/agent"
)

// Executor 执行 Agent 决策的任务
type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

// Execute 执行外部动作
func (e *Executor) Execute(ctx context.Context, agent *agent.Agent) error {
	// 模拟失败
	if len(agent.Context)%2 == 0 {
		return errors.New("mock task failed")
	}
	return nil
}
