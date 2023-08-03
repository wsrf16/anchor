@echo on

rd .\output /s /q

set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
md .\output\anchor-linux-amd64
go build -o .\output\anchor-linux-amd64\anchor main.go
rem set CGO_ENABLED=1
rem go build -o .\output\anchor-linux-amd64\anchor-ui ui.go
copy .\config.yaml .\output\anchor-linux-amd64\
7z a .\output\anchor-linux-amd64.7z .\output\anchor-linux-amd64
rd .\output\anchor-linux-amd64 /s /q

set GOARCH=arm64
set GOOS=linux
set CGO_ENABLED=0
md .\output\anchor-linux-arm64
go build -o .\output\anchor-linux-arm64\anchor main.go
rem set CGO_ENABLED=1
rem go build -o .\output\anchor-linux-arm64\anchor-ui ui.go
copy .\config.yaml .\output\anchor-linux-arm64\
7z a .\output\anchor-linux-arm64.7z .\output\anchor-linux-arm64
rd .\output\anchor-linux-arm64 /s /q

set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=0
md .\output\anchor-windows-amd64
go build -o .\output\anchor-windows-amd64\anchor.exe main.go
set CGO_ENABLED=1
go build -o .\output\anchor-windows-amd64\anchor-ui.exe ui.go
copy .\config.yaml .\output\anchor-windows-amd64\
7z a .\output\anchor-windows-amd64.7z .\output\anchor-windows-amd64
rd .\output\anchor-windows-amd64 /s /q

