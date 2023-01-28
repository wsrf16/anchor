SETLOCAL

set CGO_ENABLED=0

set GOARCH=amd64
set GOOS=linux

@REM set GOARCH=arm64
@REM set GOOS=linux

@REM set GOARCH=amd64
@REM set GOOS=windows

@REM set GOARM=7

go build

ENDLOCAL