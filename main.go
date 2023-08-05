package main

import (
	"anchor/entry/cmd"
	"time"
)

// var switchState = false
type Config struct {
	Verbose    bool
	UDPTimeout time.Duration
	TCPCork    bool
	TCP        bool
	UDP        bool
}

func main() {
	cmd.Start()

	// -s 'ss://AEAD_CHACHA20_POLY1305:123456@:8388' -verbose
	//cipher := "AEAD_CHACHA20_POLY1305"
	//cipher = "aes-256-gcm"
	//password := "123456"
	//addr := ":8388"
	//shadowsocksserver.Serve(addr, password, cipher)
}
