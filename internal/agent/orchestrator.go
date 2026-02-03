package agent

import (
	"context"
	"errors"
	"langchaingo-ai-agent/internal/ctx"
	"langchaingo-ai-agent/internal/llm"
	"langchaingo-ai-agent/internal/task"
	"time"
)

// Orchestrator 负责 Agent 的生命周期管理
// 它是“调度器”，不是执行器
type Orchestrator struct {
	LLMClient    llm.Client
	ContextMgr   *ctx.Manager
	TaskExecutor *task.Executor
}

// Run 启动一个 Agent
func (o *Orchestrator) Run(ctx context.Context, agent *Agent) error {
	agent.NextState(StateInit)
	agent.StartTime = time.Now()

	for {
		// 全局超时控制
		if agent.IsTimeout() {
			agent.NextState(StateTimeout)
			return errors.New("agent timeout")
		}

		switch agent.State {

		case StateInit:
			agent.NextState(StateThinking)

		case StateThinking:
			// 构建 prompt / ctx
			o.ContextMgr.Build(agent)
			agent.NextState(StateCallLLM)

		case StateCallLLM:
			resp, err := o.LLMClient.Generate(ctx, agent.Context)
			if err != nil {
				// 外部依赖失败，允许重试
				agent.NextState(StateFailed)
				return err
			}

			// 将 LLM 输出写入上下文
			o.ContextMgr.AppendAssistant(agent, resp)
			agent.NextState(StateExecTask)

		case StateExecTask:
			needTask := o.ContextMgr.NeedExecuteTask(agent)
			if !needTask {
				agent.NextState(StateDone)
				continue
			}

			if err := o.TaskExecutor.Execute(ctx, agent); err != nil {
				agent.NextState(StateFailed)
				return err
			}

			agent.NextState(StateUpdateCtx)

		case StateUpdateCtx:
			o.ContextMgr.UpdateAfterTask(agent)
			agent.NextState(StateThinking)

		case StateDone:
			return nil

		case StateFailed:
			return errors.New("agent failed")

		default:
			return errors.New("unknown agent state")
		}
	}
}
