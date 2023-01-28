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
)

func Serve() {
	global, err := config.Global()
	if err != nil {
		panic(err)
	}

	if global.TCP != nil {
		for _, conf := range global.TCP {
			if len(conf.Remote) > 0 {
				go tcpkit.TransferFromListenToDialAddress(conf.Local, conf.Remote)
			} else {
				go tcpkit.TransferFromListenAddress(conf.Local)
			}
		}
	}
	if global.NAT != nil {
		for _, conf := range global.NAT {
			go tcpkit.TransferFromListenToListenAddress(conf.Local, conf.Remote)
		}
	}
	if global.Link != nil {
		for _, conf := range global.Link {
			go tcpkit.TransferFromDialToDialAddress(conf.Local, conf.Remote)
		}
	}
	if global.UDP != nil {
		for _, conf := range global.UDP {
			go udpkit.TransferFromListenToDialAddress(conf.Local, conf.Remote)
		}
	}
	if global.Socks != nil {
		for _, conf := range global.Socks {
			go sockskit.TransferFromListenAddress(conf.Local)
		}
	}
	if global.SSH != nil {
		for _, conf := range global.SSH {
			go tcpkit.TransferFromListenToDialAddress(conf.Local, conf.Remote)
		}
	}
	if global.HTTP != nil {
		for _, conf := range global.HTTP {
			go httpproxy.ListenAndServe(conf.Local, conf.Remote)
		}
	}
	if global.HttpServer != nil && global.HttpServer.SSH != nil {
		ssh := collectorkit.ToPointerSlice(global.HttpServer.SSH)
		sshkit.GetSingleton().PutMulti(ssh)
		go httpserver.ListenAndServe(global.HttpServer.Local)
	}
	fmt.Println("start a anchor server")
	select {}
	fmt.Println("end a anchor server")

}
