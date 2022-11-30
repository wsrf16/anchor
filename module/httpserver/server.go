package httpserver

import (
	"anchor/support/config"
	"fmt"
	"github.com/wsrf16/swiss/sugar/base/control"
	"github.com/wsrf16/swiss/sugar/encoding/jsonkit"
	"github.com/wsrf16/swiss/sugar/netkit/http/httpserver"
	"net/http"
	"time"
)

type HttpProxy struct {
	config.HttpServerConfig
}

func Serve(conf *config.HttpServerConfig) {
	proxy, err := jsonkit.DeepCloneSimilar[HttpProxy](conf)
	if err != nil {
		panic(err)
	}
	proxy.Serve(true)
}

func (p HttpProxy) Serve(autoReconnect bool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", p.dispatch)

	control.LoopAlways(autoReconnect, func() {
		if err := http.ListenAndServe(p.Listen, mux); err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Second)
		}
	})
}

func (p HttpProxy) dispatch(w http.ResponseWriter, r *http.Request) {
	patternHandlers := make(httpserver.PatternHandlers)
	patternHandlers[`/ssh`] = SSHHandler
	patternHandlers[`/shell`] = shellHandler

	httpserver.BindMatchHandlers(patternHandlers, w, r)
}
