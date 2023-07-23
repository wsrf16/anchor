package main

import (
	"anchor/support/config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/wsrf16/swiss/sugar/logo"
	"github.com/wsrf16/swiss/sugar/netkit/socket/sockskit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/tcpkit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/udpkit"
	"os"
	"strings"
	"time"
)

func main() {
	//s := os.Args[0]
	a := app.New()
	w := a.NewWindow("anchor网络转发工具 " + config.Version)
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

	var buttonVector *widget.Button
	button := widget.NewButton("Start", func() {
		local := strings.TrimSpace(localInput.Text)
		remote := strings.TrimSpace(remoteInput.Text)
		username := strings.TrimSpace(usernameInput.Text)
		password := strings.TrimSpace(passwordInput.Text)
		proxyType := selectList.Selected
		buttonVector.SetText("Stop")
		switch proxyType {
		case "UDP":
			if len(local) > 0 && len(remote) > 0 {
				err := udpkit.TransferFromListenToDialAddress(local, remote)
				if err != nil {
					logo.Error("UDP启动错误", err)
				}
			}
			recover(buttonVector)
		case "TCP":
			if len(local) > 0 && len(remote) > 0 {
				err := tcpkit.TransferFromListenToDialAddress(local, remote)
				if err != nil {
					logo.Error("TCP启动错误", err)
				}
			} else if len(local) > 0 {
				err := tcpkit.TransferFromListenAddress(local)
				if err != nil {
					logo.Error("TCP启动错误", err)
				}
			}
			recover(buttonVector)
		case "SOCKS":
			config := sockskit.NewSocksConfig(username, password)

			if len(local) > 0 {
				err := sockskit.TransferFromListenAddress(local, config)
				if err != nil {
					logo.Error("SOCKS启动错误", err)
				}
			}
			recover(buttonVector)
		case "SSH":
			err := tcpkit.TransferFromListenToDialAddress(local, remote)
			if err != nil {
				logo.Error("SSH启动错误", err)
			}
			recover(buttonVector)
		case "NAT":
			err := tcpkit.TransferFromListenToListenAddress(local, remote)
			if err != nil {
				logo.Error("NAT启动错误", err)
			}
			recover(buttonVector)
		case "LINK":
			err := tcpkit.TransferFromDialToDialAddress(local, remote)
			if err != nil {
				logo.Error("LINK启动错误", err)
			}
			recover(buttonVector)
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

func recover(buttonVector *widget.Button) {
	time.Sleep(500 * time.Millisecond)
	buttonVector.SetText("Start")
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
