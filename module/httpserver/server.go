package httpserver

import (
	"github.com/wsrf16/swiss/sugar/netkit/http/httpserver"
)

func ListenAndServe(addr string) {
	patternHandlerMap := make(httpserver.PatternHandlerMap)
	patternHandlerMap[`/ssh`] = sshHandler
	patternHandlerMap[`/shell`] = shellHandler
	httpserver.ListenAndServe(addr, patternHandlerMap)
}
