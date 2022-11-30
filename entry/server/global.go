package server

import (
	"anchor/module/httpproxy"
	"anchor/module/httpserver"
	"anchor/support/config"
	"fmt"
	"github.com/wsrf16/swiss/sugar/base/collectorkit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/sockskit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/tcpkit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/udpkit"
	"github.com/wsrf16/swiss/sugar/netkit/sshkit"
	"time"
)

func Serve() {
	global, err := config.Global()
	if err != nil {
		panic(err)
	}

	if global.TCP != nil {
		for _, conf := range global.TCP {
			if len(conf.Forward) > 0 {
				go tcpkit.TransferToServe(conf.Listen, conf.Forward)
			} else {
				go tcpkit.TransferToHostServe(conf.Listen)
			}
		}
	}
	if global.UDP != nil {
		for _, conf := range global.UDP {
			go udpkit.TransferToServe(conf.Listen, conf.Forward)
		}
	}
	if global.Socks != nil {
		for _, conf := range global.Socks {
			go sockskit.TransferToHostServe(conf.Listen)
		}
	}
	if global.SSH != nil {
		for _, conf := range global.SSH {
			go tcpkit.TransferToServe(conf.Listen, conf.Forward)
		}
	}
	if global.HTTP != nil {
		for _, conf := range global.HTTP {
			go httpproxy.TransferToServeConf(conf)
		}
	}
	if global.HttpServer != nil && global.HttpServer.SSH != nil {
		ssh := collectorkit.ToPointerArray(global.HttpServer.SSH)
		sshkit.GetSingleton().PutMulti(ssh)
		go httpserver.Serve(global.HttpServer)
	}
	fmt.Println("start a anchor server")
	for {
		time.Sleep(5 * time.Second)
	}

}
