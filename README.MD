[TOC]

# CampusWorkGuard
---
## 如何运行
> 前置条件: Redis、MySQL
1. 将`/config/config.yaml`内容替换为自己的配置
    其中jwt.secret需要自己定义一个复杂字符串，确保安全性
2. 将`/internal/initalize.go`第42行`v.SetConfigName("config.back")`替换为`v.SetConfigName("config")`
3. 运行如下命令
    ```zsh
    # mac系统
    /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
      --remote-debugging-port=9222 \
      --user-data-dir=/tmp/chrome-data \
      --no-first-run
    ```
    ```zsh
    # windows系统，需要先找到chrome.exe的路径
   "C:\Program Files\Google\Chrome\Application\chrome.exe" ^
    --remote-debugging-port=9222 ^
    --user-data-dir="C:\temp\chrome-data" ^
    --no-first-run
    ```
   ```zsh
   # linux系统，进入chrome所在目录
   google-chrome \
    --remote-debugging-port=9222 \
    --user-data-dir=/tmp/chrome-data \
    --no-first-run
    ```
    运行命令后会弹出一个新的chrome窗口，保持该窗口打开状态，输入`http://127.0.0.1:9222/json/version` ，复制第一个`webSocketDebuggerUrl`的值，替换到`/.env`中的`CHROME_WS_URL`字段
4. 配置高德API
    在`/.env`中配置字段AMAP_KEY，申请地址：https://console.amap.com/dev/key/app
5. 运行
    ```zsh
    go mod tidy
    go run cmd/CampusWorkGuard/main.go
    ```