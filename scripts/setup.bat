@echo off

set GOPROXY=https://goproxy.io
set GO111MODULE=on

set command=%1

if "%command%"=="build" (
  go build -o build/amazing-proxy src/main.go
) else (
  go run src/main.go
)
