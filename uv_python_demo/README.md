## 1. 前置准备

```bash
// 安装 uv
curl -LsSf https://astral.sh/uv/install.sh | sh

// 验证
uv

// 创建项目文件夹（创建脚本时执行）
uv init test
cd test

// 创建虚拟环境并激活
uv venv
source .venv/bin/activate

// 安装依赖
uv add "mcp[cli]" httpx

// 调试脚本（调试脚本时执行）
mcp dev add.py

// 创建 server 文件（创建脚本时执行）
touch add.py
```

## 2. 编写脚本

`add.py`

```python
from typing import Any
import httpx
from mcp.server.fastmcp import FastMCP

# Initialize FastMCP server
mcp = FastMCP("grafana")

@mcp.tool()
def add(a: int, b: int) -> int:
    """Add two numbers together."""
    return a + b

if __name__ == "__main__":
    # Initialize and run the server
    mcp.run(transport='stdio')
```

## 3. 运行调试

执行命令 `mcp dev add.py`

> 需要保证 node version >=  18

![1745228559756_q0w4lepr29](https://pub-6bd5d0c20d254a3e9d8dea968e62938c.r2.dev/2025/4/1745228559756_q0w4lepr29)

点击链接跳转调试页面

1. 点击 Connect

![1745228695408_zjvfkz0n6q](https://pub-6bd5d0c20d254a3e9d8dea968e62938c.r2.dev/2025/4/1745228695408_zjvfkz0n6q)

2. 按照下图顺序点击进行调试

![1745229137020_vktjeo5gxi](https://pub-6bd5d0c20d254a3e9d8dea968e62938c.r2.dev/2025/4/1745229137020_vktjeo5gxi)

可以看到执行了我们定义的 add 方法