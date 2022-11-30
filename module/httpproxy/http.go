package httpproxy

import (
	"anchor/support/config"
	"github.com/wsrf16/swiss/sugar/base/stringkit"
	"github.com/wsrf16/swiss/sugar/encoding/jsonkit"
	"github.com/wsrf16/swiss/sugar/netkit/http/httpkit"
	"github.com/wsrf16/swiss/sugar/netkit/http/httpserver"
	"net/http"
)

type HttpProxy struct {
	config.HTTPConfig
}

func TransferToServeConf(conf config.HTTPConfig) error {
	proxy, err := jsonkit.DeepCloneSimilar[HttpProxy](conf)
	if err != nil {
		panic(err)
	}
	return proxy.Serve(true)
}

func TransferToServeSpecial(listen string, forward string) error {
	proxy := &HttpProxy{}
	proxy.Listen = listen
	proxy.Forward = forward
	return proxy.Serve(true)
}

func (p HttpProxy) Serve(autoReconnect bool) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", p.dispatch)

	return http.ListenAndServe(p.Listen, mux)
}

func (p HttpProxy) dispatch(w http.ResponseWriter, r *http.Request) {
	patternHandlers := make(httpserver.PatternHandlers)
	patternHandlers[`^/`] = p.httpHandler

	httpserver.BindMatchHandlers(patternHandlers, w, r)
}

func (p HttpProxy) httpHandler(w http.ResponseWriter, r *http.Request) {
	url := stringkit.JoinURL(r.RequestURI, p.Forward)

	if err := httpkit.ForwardTo(w, r, url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
