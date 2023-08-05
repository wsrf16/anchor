package shadowsocksserver

import "time"

var config Config

type Config struct {
	Verbose    bool
	UDPTimeout time.Duration
	TCPCork    bool
	TCP        bool
	UDP        bool
}
