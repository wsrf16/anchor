package tcproxy

import (
	"net"
)

var connectionPool = make(map[string]net.Conn, 0)

//func Register(raddress string, laddress string) error {
//    raddr, err := tcptrans.NewTCPAddr(raddress)
//    if err != nil {
//        return err
//    }
//    rconn, err := net.DialTCP("tcp", nil, raddr)
//    if err != nil {
//        return err
//    }
//
//    laddr, err := tcptrans.NewTCPAddr(laddress)
//    if err != nil {
//        return err
//    }
//    lconn, err := net.DialTCP("tcp", nil, laddr)
//    if err != nil {
//        return err
//    }
//
//    return socket.TransferThenClose(rconn, lconn)
//}
