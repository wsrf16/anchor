package httpproxy

import (
	"github.com/wsrf16/swiss/netkit/layer/app/http/httpkit"
	"github.com/wsrf16/swiss/netkit/layer/app/http/httpserver"
	"github.com/wsrf16/swiss/sugar/base/lambda"
	"github.com/wsrf16/swiss/sugar/base/stringkit"
	"net/http"
	"strings"
)

func ListenAndServe(listen string, forward string) error {
	patternHandlerMap := make(httpserver.PatternHandlerMap)
	patternHandlerMap[`^/`] = func(w http.ResponseWriter, r *http.Request) {
		url := lambda.If[string](len(forward) > 0, stringkit.JoinURL(addScheme(forward), r.URL.Path), r.RequestURI)

		if err := httpkit.ForwardTo(w, r, url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	_, err := httpserver.ListenAndServe(listen, patternHandlerMap)
	return err
}

func addScheme(url string) string {
	if len(url) > 0 {
		url = lambda.If[string](strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://"), url, "http://"+url)
	}
	return url
}
