package config

import "github.com/wsrf16/swiss/sugar/netkit/sshkit"

type RootConfig struct {
	Setting    *Setting          `json:"setting"`
	TCP        []TCPConfig       `json:"tcp"`
	UDP        []UDPConfig       `json:"udp"`
	NAT        []NATConfig       `json:"nat"`
	Link       []LinkConfig      `json:"link"`
	Socks      []SocksConfig     `json:"socks"`
	HTTP       []HTTPConfig      `json:"http"`
	SSH        []SSHConfig       `json:"ssh"`
	HttpServer *HttpServerConfig `json:"httpserver"`
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
	Local    string `json:"local"`
	Username string `json:"username"`
	Password string `json:"password"`
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
