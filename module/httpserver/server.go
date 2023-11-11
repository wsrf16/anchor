package httpserver

import (
	"github.com/wsrf16/swiss/netkit/layer/app/http/httpserver"
)

func ListenAndServe(addr string) {
	patternHandlerMap := make(httpserver.PatternHandlerMap)
	patternHandlerMap[`/ssh`] = sshHandler
	patternHandlerMap[`/shell`] = shellHandler
	httpserver.ListenAndServe(addr, patternHandlerMap)
}
