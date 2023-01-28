@echo off
set CGO_ENABLED=0

set GOARCH=amd64
set GOOS=linux
go build -o .\output\anchor-linux-amd64\anchor
copy .\config.yaml .\output\anchor-linux-amd64\
7z a .\output\anchor-linux-amd64.7z .\output\anchor-linux-amd64
rd .\output\anchor-linux-amd64 /s /q

set GOARCH=arm64
set GOOS=linux
go build -o .\output\anchor-linux-arm64\anchor
copy .\config.yaml .\output\anchor-linux-arm64\
7z a .\output\anchor-linux-arm64.7z .\output\anchor-linux-arm64
rd .\output\anchor-linux-arm64 /s /q

set GOARCH=amd64
set GOOS=windows
go build -o .\output\anchor-windows-amd64\anchor.exe
copy .\config.yaml .\output\anchor-windows-amd64\
7z a .\output\anchor-windows-amd64.7z .\output\anchor-windows-amd64
rd .\output\anchor-windows-amd64 /s /q

