package main

import (
	"anchor/entry/cmd"
	"github.com/wsrf16/swiss/sugar/logo"
)

func main() {
	logo.SetLevel(logo.DebugLevel)
	cmd.Start()
}

func main0() {
	// //////////////////

	// ifname := "MyNIC"
	// dev, err := tun.CreateTUN(ifname, 0)
	// if err != nil {
	//     panic(err)
	// }
	// defer dev.Close()
	// // 保存原始设备句柄
	// nativeTunDevice := dev.(*tun.NativeTun)
	//
	// // 获取LUID用于配置网络
	// link := winipcfg.LUID(nativeTunDevice.LUID())
	//
	// ip, err := netip.ParsePrefix("10.0.0.77/24")
	// if err != nil {
	//     panic(err)
	// }
	// err = link.SetIPAddresses([]netip.Prefix{ip})
	// if err != nil {
	//     panic(err)
	// }
	//
	// n := 2048
	// buf := make([]byte, n)
	//
	// // 读取ICMP
	// for {
	//     n = 2048
	//     n, err = dev.Read(buf, 0)
	//     if err != nil {
	//         panic(err)
	//     }
	//     const ProtocolICMP = 1
	//     header, err := ipv4.ParseHeader(buf[:n])
	//     if err != nil {
	//         continue
	//     }
	//     if header.Protocol == ProtocolICMP {
	//         log.Println("Src:", header.Src, " dst:", header.Dst)
	//         msg, _ := icmp.ParseMessage(ProtocolICMP, buf[header.Len:])
	//         log.Println(">> ICMP:", msg.Type)
	//         break
	//     }
	// }

}

// func wintun(dev tun.Device) {
//    n := 2048
//    bufs := make([][]byte, n)
//    for i := range bufs {
//        bufs[i] = make([]byte, 1024)
//    }
//
//    size := make([]int, n)
//    size[0] = 100
//    // 读取ICMP
//    for {
//        n = 2048
//        n, err := dev.Read(bufs, size, 1)
//        if err != nil {
//            panic(err)
//        }
//        const ProtocolICMP = 1
//        header, err := ipv4.ParseHeader(bufs[0])
//        if err != nil {
//            continue
//        }
//        if header.Protocol == ProtocolICMP {
//            log.Println("Src:", header.Src, " dst:", header.Dst)
//            msg, _ := icmp.ParseMessage(ProtocolICMP, bufs[0][header.Len:])
//            log.Println(">> ICMP:", msg.Type)
//            marshal, err := msg.Body.Marshal(n)
//            err = err
//            log.Println(">> ICMP:", string(marshal))
//            if body, ok := msg.Body.(*icmp.Echo); ok {
//                // do something with Body as an *icmp.Echo type
//                log.Println(">> ICMP2:", string(body.Data))
//
//            }
//            // break
//        } else if header.Protocol == 6 {
//            log.Println("Src:", header.Src, " dst:", header.Dst)
//            msg, _ := icmp.ParseMessage(ProtocolICMP, bufs[0][header.Len:])
//            log.Println(">> tcp:", msg.Type)
//            marshal, err := msg.Body.Marshal(n)
//            err = err
//            log.Println(">> tcp:", string(marshal))
//
//        }
//    }
//
//    time.Sleep(time.Second * 30)
// }
