package config

import "github.com/wsrf16/swiss/sugar/netkit/sshkit"

type RootConfig struct {
	Setting    *Setting          `json:"setting"`
	TCP        []TCPConfig       `json:"tcp"`
	UDP        []UDPConfig       `json:"udp"`
	Socks      []SocksConfig     `json:"socks"`
	HTTP       []HTTPConfig      `json:"http"`
	SSH        []SSHConfig       `json:"ssh"`
	HttpServer *HttpServerConfig `json:"httpserver"`
}
type Setting struct {
	Allows []string `json:"allows"`
}
type TCPConfig struct {
	Listen  string `json:"listen"`
	Forward string `json:"forward"`
}
type UDPConfig struct {
	Listen  string `json:"listen"`
	Forward string `json:"forward"`
}
type SocksConfig struct {
	Listen string `json:"listen"`
}
type HTTPConfig struct {
	Listen    string `json:"listen"`
	Forward   string `json:"forward"`
	AddedHead string `json:"addedHead"`
}
type SSHConfig struct {
	Listen  string `json:"listen"`
	Forward string `json:"forward"`
}
type SSH struct {
	ID         string `json:"id"`
	Addr       string `json:"addr"`
	User       string `json:"user"`
	Password   string `json:"password"`
	PrivateKey string `json:"privateKey"`
}
type HttpServerConfig struct {
	Listen string               `json:"listen"`
	Shell  *ShellConfig         `json:"shell"`
	SSH    []sshkit.SSHProperty `json:"ssh"`
}

type ShellConfig struct {
	Enabled bool `json:"enabled"`
}
