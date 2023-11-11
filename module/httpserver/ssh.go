package httpserver

import (
	"anchor/support/api/results"
	"anchor/support/bean"
	"github.com/wsrf16/swiss/netkit/layer/app/http/httpkit"
	"github.com/wsrf16/swiss/netkit/layer/app/sshkit"
	"net/http"
)

func sshHandler(w http.ResponseWriter, r *http.Request) {
	// {"type":"ssh","shell":{"commands":["whoami", "find"]},"ssh":{"commands":["whoami", "find"],"serverId":"mecs.com:22"}}
	// {"commands":["whoami", "find"],"serverId":"mecs.com:22"}
	w.Header().Set("Content-Type", "application/json")

	t, err := httpkit.ParseRequestBody[bean.SSHRequest](r)
	if err != nil {
		results.WriteHttpFailed(w, err.Error())
	}

	run(t, w)
}

func run(t *bean.SSHRequest, w http.ResponseWriter) {
	property, err := sshkit.GetSingleton().Get(t.ServerId)
	if err != nil {
		results.WriteHttpFailed(w, err.Error())
	}

	client, err := sshkit.BuildClientByProperty(property)
	if err != nil {
		results.WriteHttpFailed(w, err.Error())
	}

	data, err := sshkit.RunBatch(client, t.Commands)
	if err != nil {
		results.WriteHttpFailedData(w, "操作失败", data)
	} else {
		results.WriteHttpSucceed(w, data)
	}
}
