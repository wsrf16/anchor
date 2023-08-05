package shadowsocksserver

import (
	"anchor/module/shadowsocksserver/core"
	"encoding/base64"
	"os"
	"os/signal"
	"syscall"
)

func Serve(addr string, password string, cipher string, key string, conf Config) error {
	var err error

	config.Verbose = conf.Verbose
	tcp, udp := conf.TCP, conf.UDP
	if !tcp && !udp {
		tcp = true
	}

	var keyByte []byte
	if len(key) > 0 {
		keyByte, err = base64.URLEncoding.DecodeString(key)
		if err != nil {
			return err
		}
	}

	ciph, err := core.PickCipher(cipher, keyByte, password)
	if err != nil {
		return err
	}

	if udp {
		go RemoteUDP(addr, ciph.PacketConn)
	}
	if tcp {
		go RemoteTCP(addr, ciph.StreamConn)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	return nil
}
