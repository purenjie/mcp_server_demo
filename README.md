# mcp_server_demo

这是一个 MCP（Model Context Protocol）服务的 demo 仓库，展示了 MCP 服务的多种实现方式。MCP 服务可以通过不同的编程语言和框架来实现，以下是目前包含的实现方式：

## 实现方式

### 1. 基于 UV Python STDIO 的实现
- **目录**: `uv_python_demo`
- **描述**: 使用 Python 和 UV 工具链实现的 MCP 服务。UV 是一个高效的 Python 包管理器，适用于快速开发和部署。
- **主要文件**:
  - `add.py`: 实现了 `add` 方法，用于演示基本的 MCP 服务功能。
  - `main.py`: 服务入口文件。
- **使用方法**:

```bash
// 没有安装的话执行
curl -LsSf https://astral.sh/uv/install.sh | sh
uv

cd uv_python_demo
uv venv
source .venv/bin/activate
uv add "mcp[cli]" httpx
mcp dev add.py
```

### 2. 基于 Golang 的 SSE 服务实现
- **目录**: `golang_sse_demo`
- **描述**: 使用 Golang 实现的 MCP 服务，利用 Server-Sent Events (SSE) 进行实时通信。
- **主要文件**:
  - `main.go`: 实现了 `add` 方法，用于演示基本的 MCP 服务功能。

- **使用方法**:

```bash
cd golang_sse_demo

// stdio 模式
go build
// 通过 command 方式，将二进制文件 golang_sse_demo 的路径作为执行路径
// e.g. /Users/my/Projects/mcp_server_demo/golang_sse_demo

// sse 模式
go run main.go
sse 地址: http://localhost:8080/sse
```


## 贡献
欢迎贡献更多的 MCP 服务实现方式！请 fork 本仓库，添加您的实现，并在 PR 中描述您的实现方式。

## 许可证
MIT
