package agent

// AgentState 表示 Agent 当前所处的阶段
// 为什么要拆状态？
// 答：为了可恢复、可观测、可重试
type AgentState string

const (
	StateInit      AgentState = "INIT"      //逻辑状态
	StateThinking  AgentState = "THINKING"  //逻辑状态
	StateCallLLM   AgentState = "CALL_LLM"  //不可控外部依赖状态
	StateExecTask  AgentState = "EXEC_TASK" //不可控外部依赖状态
	StateUpdateCtx AgentState = "UPDATE_CONTEXT"
	StateDone      AgentState = "DONE"   //终态
	StateFailed    AgentState = "FAILED" //终态
	StateTimeout   AgentState = "TIMEOUT"
)
