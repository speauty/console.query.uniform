@echo off
chcp 65001
::获取当前bat所在目录
::set workspace=%~dp0
set app=query
set mod=console.query.uniform

set /p opt=请输入操作(run-运行, build-编译):
if "%opt%" == "run" (
    call:run
) else if "%opt%" == "build" (
    call:build
) else (
    echo 暂无该操作
)

pause


:run
go run %mod%
goto:eof

:build
::GOOS/GOARCH 参考 https://github.com/goreleaser/goreleaser/issues/142
set GOOS=windows
set GOARCH=amd64
set /p "GOOS=设置编译目标操作系统(默认值:%GOOS%): "
set /p "GOARCH=设置编译目标架构(默认值:%GOARCH%): "
set "app=%app%-%GOOS%-%GOARCH%"
if "%GOOS%" == "windows" set "app=%app%.exe"
go build -ldflags "-s -w" -o "%app%"
echo 构建完成: %app%
goto:eof