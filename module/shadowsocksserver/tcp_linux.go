package shadowsocksserver

import (
	"net"

	"anchor/module/shadowsocksserver/nfutil"
	"anchor/module/shadowsocksserver/socks"
)

func getOrigDst(c net.Conn, ipv6 bool) (socks.Addr, error) {
	if tc, ok := c.(*net.TCPConn); ok {
		addr, err := nfutil.GetOrigDst(tc, ipv6)
		return socks.ParseAddr(addr.String()), err
	}
	panic("not a TCP connection")
}

// Listen on addr for netfilter redirected TCP connections
func redirLocal(addr, server string, shadow func(net.Conn) net.Conn) {
	logf("TCP redirect %s <-> %s", addr, server)
	tcpLocal(addr, server, shadow, func(c net.Conn) (socks.Addr, error) { return getOrigDst(c, false) })
}

// Listen on addr for netfilter redirected TCP IPv6 connections.
func redir6Local(addr, server string, shadow func(net.Conn) net.Conn) {
	logf("TCP6 redirect %s <-> %s", addr, server)
	tcpLocal(addr, server, shadow, func(c net.Conn) (socks.Addr, error) { return getOrigDst(c, true) })
}
