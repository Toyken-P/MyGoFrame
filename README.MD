# GoFrame Template For SingleRepo

基本配置：
- golang
- mysql、redis：分别在`hack/`和`manifest/config`下的`config.yaml`配置
命令行：
- 包下载和运行
  ```
  go env -w GO111MODULE=on
  go env -w GOPROXY=https://goproxy.io,direct
  go mod tidy
  go run main.go
  ```