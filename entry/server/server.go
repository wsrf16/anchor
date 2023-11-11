package server

import (
	"anchor/module/httpproxy"
	"anchor/module/httpserver"
	"anchor/support/config"
	"fmt"
	"github.com/wsrf16/swiss/netkit/layer/app/sshkit"
	"github.com/wsrf16/swiss/netkit/layer/transport/httptrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/sockstrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/sstrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/sstrans/sstcptrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/tcptrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/udptrans"
	"github.com/wsrf16/swiss/netkit/tun2socks"
	"github.com/wsrf16/swiss/sugar/base/collectorkit"
	"golang.org/x/net/proxy"
)

func Serve() {
	global, err := config.Global()
	if err != nil {
		panic(err)
	}

	if global.TCP != nil {
		for _, conf := range global.TCP {
			if len(conf.Remote) > 0 {
				go tcptrans.TransferFromListenToDialAddress(conf.Local, conf.Remote, true, nil)
			} else {
				go httptrans.TransferFromListenAddress(conf.Local, true, nil)
			}
		}
	}
	if global.NAT != nil {
		for _, conf := range global.NAT {
			go tcptrans.TransferFromListenToListenAddress(conf.Local, conf.Remote, true, nil)
		}
	}
	if global.Link != nil {
		for _, conf := range global.Link {
			go tcptrans.TransferFromDialToDialAddress(conf.Local, conf.Remote, true)
		}
	}
	if global.UDP != nil {
		for _, conf := range global.UDP {
			go udptrans.TransferFromListenToDialAddress(conf.Local, conf.Remote, true, nil)
		}
	}
	if global.Socks != nil {
		for _, conf := range global.Socks {
			auth := &proxy.Auth{User: conf.User, Password: conf.Password}
			go sockstrans.TransferFromListenAddress(conf.Local, auth, true, nil)
		}
	}
	if global.SSH != nil {
		for _, conf := range global.SSH {
			go tcptrans.TransferFromListenToDialAddress(conf.Local, conf.Remote, true, nil)
		}
	}
	if global.SS != nil {
		for _, conf := range global.SS {
			config, err := sstrans.BuildConfig(conf.Algorithm, nil, conf.Password)
			if err != nil {
				panic(err)
			}
			go sstcptrans.TransferFromListenAddress(conf.Local, config, true, nil)
		}
	}
	if global.T2S != nil {
		for _, conf := range global.T2S {
			go tun2socks.Serve(&conf)
		}
	}
	if global.HTTP != nil {
		for _, conf := range global.HTTP {
			if len(conf.Remote) > 0 {
				go httpproxy.ListenAndServe(conf.Local, conf.Remote)
			} else {
				go httptrans.TransferFromListenAddress(conf.Local, true, nil)
			}
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
