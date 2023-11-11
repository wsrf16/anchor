package main

import (
	"anchor/support/config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/wsrf16/swiss/netkit/layer/transport/httptrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/sockstrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/tcptrans"
	"github.com/wsrf16/swiss/netkit/layer/transport/udptrans"
	"github.com/wsrf16/swiss/sugar/logo"
	"golang.org/x/net/proxy"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var ch = make(chan int)
var switchState = false
var srcChannel = make(chan net.Conn)

var tcpListenerChannel = make(chan *net.TCPListener, 4096)
var udpListenerChannel = make(chan net.Conn, 4096)
var buttonVector *widget.Button

var wg sync.WaitGroup
var m sync.Mutex

func setOff(listenerChannels ...chan *net.TCPListener) {
	go func(listenerChannels ...chan *net.TCPListener) {
		m.Lock()
		defer m.Unlock()
		if switchState {
			for _, listenerChannel := range listenerChannels {
				if listenerChannel != nil {
					listener := <-listenerChannel
					listener.Close()
					switchState = false
					time.Sleep(200 * time.Millisecond)
					buttonVector.SetText("Start")
				}
			}
		}
	}(listenerChannels...)
}

func setOn() {
	go func() {
		m.Lock()
		defer m.Unlock()
		if switchState == false {
			switchState = true
			time.Sleep(200 * time.Millisecond)
			buttonVector.SetText("Stop")
		}
	}()
}

func main() {

	//s := os.Args[0]
	a := app.New()
	w := a.NewWindow("anchor " + config.Version)
	selectList := widget.NewSelect([]string{"TCP", "UDP", "SOCKS", "SSH", "NAT", "LINK"}, func(s string) {
		if s == "SOCKS" {
		}
	})
	selectList.SetSelectedIndex(0)

	localLabel := widget.NewLabel("Local Address")
	localInput := widget.NewEntry()
	localInput.MultiLine = false
	localInput.PlaceHolder = "eg. :9090/127.0.0.1:9090"

	separator := widget.NewSeparator()

	remoteLabel := widget.NewLabel("Remote Address")
	remoteInput := widget.NewEntry()
	remoteInput.MultiLine = false
	remoteInput.PlaceHolder = "eg. 127.0.0.1:9090"

	//separator2 := widget.NewSeparator()

	usernameLabel := widget.NewLabel("Username")
	usernameInput := widget.NewEntry()
	usernameInput.MultiLine = false
	usernameInput.PlaceHolder = "username"

	//separator3 := widget.NewSeparator()

	passwordLabel := widget.NewLabel("Password")
	passwordInput := widget.NewEntry()
	passwordInput.MultiLine = false
	passwordInput.PlaceHolder = "password"

	button := widget.NewButton("Start", func() {
		local := strings.TrimSpace(localInput.Text)
		remote := strings.TrimSpace(remoteInput.Text)
		username := strings.TrimSpace(usernameInput.Text)
		password := strings.TrimSpace(passwordInput.Text)
		proxyType := selectList.Selected
		switch proxyType {
		case "UDP":
			if switchState == false {
				RunBackground(func() {
					setOn()
					if len(local) > 0 && len(remote) > 0 {
						transfer, err := udptrans.BuildTransfer(local, remote, false)
						if err != nil {
							logo.Error("", err)
						}
						err = udptrans.TransferFromListenToDialAddress(local, remote, false, transfer)
						if err != nil {
							setOff(nil)
							logo.Error("UDP连接错误或中断", err)
						}
					} else {
						setOff(nil)
						logo.Error("未填写Remote Address", nil)
					}
				})
			} else {
				setOff(nil)
			}

		case "TCP":
			if switchState == false {
				RunBackground(func() {
					setOn()
					if len(local) > 0 && len(remote) > 0 {
						transfer, err := tcptrans.BuildTransfer(local, remote, false)
						if err != nil {
							logo.Error("", err)
						}
						err = tcptrans.TransferFromListenToDialAddress(local, remote, false, transfer)
						if err != nil {
							setOff(tcpListenerChannel)
							logo.Error("TCP连接错误或中断", err)
						}
					} else if len(local) > 0 {
						transfer, err := httptrans.BuildTransfer(local, false)
						if err != nil {
							logo.Error("", err)
						}
						err = httptrans.TransferFromListenAddress(local, false, transfer)
						if err != nil {
							setOff(tcpListenerChannel)
							logo.Error("TCP连接错误或中断", err)
						}
					}
				})
			} else {
				setOff(tcpListenerChannel)
			}
		case "SOCKS":
			if switchState == false {
				RunBackground(func() {
					setOn()
					if len(local) > 0 {
						transfer, err := sockstrans.BuildTransfer(local, false)
						if err != nil {
							logo.Error("", err)
						}
						auth := &proxy.Auth{User: username, Password: password}
						err = sockstrans.TransferFromListenAddress(local, auth, false, transfer)
						if err != nil {
							setOff(tcpListenerChannel)
							logo.Error("SOCKS连接错误或中断", err)
						}
					}
				})
			} else {
				setOff(tcpListenerChannel)
			}
		case "SSH":
			if switchState == false {
				RunBackground(func() {
					setOn()
					transfer, err := tcptrans.BuildTransfer(local, remote, false)
					if err != nil {
						logo.Error("", err)
					}
					err = tcptrans.TransferFromListenToDialAddress(local, remote, false, transfer)
					if err != nil {
						setOff(tcpListenerChannel)
						logo.Error("SSH连接错误或中断", err)
					}
				})
			} else {
				setOff(tcpListenerChannel)
			}
		case "NAT":
			if switchState == false {
				RunBackground(func() {
					setOn()
					transfer, err := tcptrans.BuildTransfer(local, remote, false)
					if err != nil {
						logo.Error("", err)
					}
					err = tcptrans.TransferFromListenToListenAddress(local, remote, false, transfer)
					if err != nil {
						setOff(tcpListenerChannel)
						logo.Error("NAT连接错误或中断", err)
					}
				})
			} else {
				setOff(tcpListenerChannel, tcpListenerChannel)
			}

		case "LINK":
			if switchState == false {
				RunBackground(func() {
					setOn()
					err := tcptrans.TransferFromDialToDialAddress(local, remote, false)
					if err != nil {
						setOff(nil)
						logo.Error("LINK连接错误或中断", err)
					}
				})
			} else {
				setOff(nil)
			}

		}

	})
	buttonVector = button

	box := container.NewVBox(selectList, localLabel, localInput, separator, remoteLabel, remoteInput, separator, usernameLabel, usernameInput, separator, passwordLabel, passwordInput, separator, button)

	w.SetContent(box)
	w.Resize(fyne.NewSize(500, 400))
	//popWin := widget.NewPopUp(box, w.Canvas())
	//popWin.ShowAtPosition(fyne.NewPos(
	//    w.Canvas().Size().Width/2-popWin.MinSize().Width/2,
	//    w.Canvas().Size().Height/2-popWin.MinSize().Height/2))
	w.ShowAndRun()
}

func RunBackground(task func()) {
	go func() {
		task()
	}()
}

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}
