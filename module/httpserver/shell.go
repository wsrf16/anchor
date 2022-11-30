package httpserver

import (
	"anchor/support/api/results"
	"anchor/support/bean"
	"github.com/wsrf16/swiss/sugar/console/shellkit"
	"github.com/wsrf16/swiss/sugar/netkit/http/httpkit"
	"net/http"
)

func shellHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	t, err := httpkit.ParseRequestBody[bean.ShellRequest](r)
	if err != nil {
		results.WriteHttpFailed(w, err.Error())
	}

	runShell(t, w)
}

func runShell(t *bean.ShellRequest, w http.ResponseWriter) {
	result, err := shellkit.ExecuteBatch(t.Commands)
	if err != nil {
		results.WriteHttpFailedData(w, "操作失败", result)
	} else {
		results.WriteHttpSucceed(w, result)
	}
}
