# langchaingo-agent

一个基于 Go 的 AI Agent 平台示例，包含 Agent 状态机、调度器（Orchestrator）、LLM 抽象、RAG、上下文管理等核心模块。

## 项目结构

> **说明**：请使用等宽代码块展示目录树，否则在 Markdown 渲染时容易被“压扁”。

```text
ai-agent-platform/
├── cmd/
│   └── server/
│       └── main.go              # 程序入口
├── internal/
│   ├── api/
│   │   └── http.go              # HTTP 接口层
│   ├── agent/
│   │   ├── agent.go             # Agent 核心结构
│   │   ├── state.go             # 状态机定义
│   │   └── orchestrator.go      # 调度器（最核心）
│   ├── llm/
│   │   ├── client.go            # LLM 接口抽象
│   │   ├── openai.go            # OpenAI 实现（可 mock）
│   │   └── fallback.go          # fallback / retry 逻辑
│   ├── rag/
│   │   ├── rag.go               # RAG 主逻辑
│   │   └── memory_store.go      # 简单向量存储（内存版）
│   ├── context/
│   │   └── manager.go           # 上下文管理（非常重要）
│   ├── task/
│   │   └── executor.go          # 工具 / 任务执行
│   ├── metrics/
│   │   └── metrics.go           # token / latency / cost
│   └── model/
│       └── request.go           # 通用请求/响应模型
├── pkg/
│   └── utils/
│       └── retry.go
├── go.mod
└── README.md
```

## 目录设计说明（简要）

* **cmd/**：程序入口，负责启动和组装依赖
* **internal/**：核心业务代码，对外不可直接依赖

    * **agent/**：Agent 定义、状态机与调度器
    * **llm/**：大模型抽象层，支持多实现与 fallback
    * **rag/**：检索增强生成（RAG）相关逻辑
    * **context/**：会话 / 长上下文管理（关键模块）
    * **task/**：工具与任务执行器
    * **metrics/**：性能与成本指标采集
* **pkg/**：可复用的公共工具包

