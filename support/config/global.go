package config

import (
	"github.com/wsrf16/swiss/netkit/layer/app/sshkit"
	"github.com/wsrf16/swiss/netkit/tun2socks/support/t2sconfig"
	"golang.org/x/net/proxy"
)

type RootConfig struct {
	Setting    *Setting          `json:"setting"`
	TCP        []TCPConfig       `json:"tcp"`
	UDP        []UDPConfig       `json:"udp"`
	NAT        []NATConfig       `json:"nat"`
	Link       []LinkConfig      `json:"link"`
	Socks      []SocksConfig     `json:"socks"`
	HTTP       []HTTPConfig      `json:"http"`
	SSH        []SSHConfig       `json:"ssh"`
	SS         []SSConfig        `json:"ss"`
	HttpServer *HttpServerConfig `json:"httpserver"`
	T2S        []T2SConfig       `json:"t2s"`
}
type Setting struct {
	Allows []string `json:"allows"`
}
type TCPConfig struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}
type UDPConfig struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}
type SocksConfig struct {
	proxy.Auth
	Local string `json:"local"`
	//Username string `json:"username"`
	//Password string `json:"password"`
}
type HTTPConfig struct {
	Local     string `json:"local"`
	Remote    string `json:"forward"`
	AddedHead string `json:"addedHead"`
}
type SSHConfig struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}

type SSConfig struct {
	Local     string `json:"local"`
	Password  string `json:"password"`
	Algorithm string `json:"algorithm"`
	TCP       bool   `json:"tcp"`
	UDP       bool   `json:"udp"`
}

type T2SConfig = t2sconfig.TunConfig

//type T2SConfig struct {
//    //Local    string `json:"local"`
//    //Password string `json:"password"`
//    //Algorithm   string `json:"cipher"`
//    //TCP      bool   `json:"tcp"`
//    //UDP      bool   `json:"udp"`
//}

//	type SSH struct {
//		ID         string `json:"id"`
//		Addr       string `json:"addr"`
//		User       string `json:"user"`
//		Password   string `json:"password"`
//		PrivateKey string `json:"privateKey"`
//	}
type HttpServerConfig struct {
	Local string               `json:"local"`
	Shell *ShellConfig         `json:"shell"`
	SSH   []sshkit.SSHProperty `json:"ssh"`
}

type ShellConfig struct {
	Enabled bool `json:"enabled"`
}
type NATConfig struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}
type LinkConfig struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}
