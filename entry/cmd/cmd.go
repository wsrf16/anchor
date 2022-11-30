package cmd

import (
	"anchor/entry/server"
	"anchor/module/httpproxy"
	"github.com/spf13/cobra"
	"github.com/wsrf16/swiss/sugar/base/control"
	"github.com/wsrf16/swiss/sugar/logo"
	"github.com/wsrf16/swiss/sugar/netkit/socket/sockskit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/tcpkit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/udpkit"
	"github.com/wsrf16/swiss/sugar/netkit/sshkit"
	"os"
)

func checkFlags(cmd *cobra.Command, flagKeys ...string) (T bool) {
	for _, flagKey := range flagKeys {
		flagVal, err := cmd.Flags().GetString(flagKey)
		check := control.If(len(flagVal) < 1 || err != nil, false, true)
		if !check {
			cmd.Help()
			os.Exit(1)
			return false
		}
	}
	return true
}

func Start() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start a anchor server",
		//Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			server.Serve()
		},
	}

	var tcpCmd = &cobra.Command{
		Use:   "tcp",
		Short: "Start a tcp server",
		Long:  "tcp -l <listen-address> [-f <forward-address>]",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				if !checkFlags(cmd, "listen") {
					cmd.Help()
				}
				listen, _ := cmd.Flags().GetString("listen")
				forward, _ := cmd.Flags().GetString("forward")
				if len(listen) > 0 && len(forward) > 0 {
					err := tcpkit.TransferToServe(listen, forward)
					if err != nil {
						panic(err)
					}
				} else if len(listen) > 0 {
					err := tcpkit.TransferToHostServe(listen)
					if err != nil {
						panic(err)
					}
				}
			} else {
				cmd.Help()
			}
		},
	}

	var udpCmd = &cobra.Command{
		Use:   "udp",
		Short: "Start a udp server",
		Long:  "udp -l <listen-address> -f <forward-address>",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkFlags(cmd, "listen", "forward")

				listen, _ := cmd.Flags().GetString("listen")
				forward, _ := cmd.Flags().GetString("forward")
				err := udpkit.TransferToServe(listen, forward)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	var socksCmd = &cobra.Command{
		Use:   "socks",
		Short: "Start a socks server",
		Long:  "tcp -l <listen-address>",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				if !checkFlags(cmd, "listen") {
					cmd.Help()
				}
				listen, _ := cmd.Flags().GetString("listen")
				//forward, _ := cmd.Flags().GetString("forward")
				if len(listen) > 0 {
					err := sockskit.TransferToHostServe(listen)
					if err != nil {
						panic(err)
					}
				}
			} else {
				cmd.Help()
			}
		},
	}
	var sshCmd = &cobra.Command{
		Use:   "ssh",
		Short: "Start a ssh server",
		Long:  "ssh -l <listen-address> -f <forward-address>",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkFlags(cmd, "listen", "forward")

				listen, _ := cmd.Flags().GetString("listen")
				forward, _ := cmd.Flags().GetString("forward")
				err := tcpkit.TransferToServe(listen, forward)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}

	var sshPtyCmd = &cobra.Command{
		Use:   "ssh-pty",
		Short: "Login remote ssh",
		Long:  "ssh-pty <address> -u <username> -p <password>",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkFlags(cmd, "username")

				_, err := cmd.Flags().GetString("password")
				if err != nil {
					cmd.Help()
					os.Exit(1)
				}

				if len(args) < 1 {
					cmd.Help()
					os.Exit(1)
				}

				username, _ := cmd.Flags().GetString("username")
				password, _ := cmd.Flags().GetString("password")
				err = sshkit.SimplePtyFlat(args[0], username, password, "")
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	sshPtyCmd.PersistentFlags().StringP("username", "u", "", "<username>")
	sshPtyCmd.PersistentFlags().StringP("password", "p", "", "<password>")

	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start a http server",
		Long:  "http -l <listen-address> -f <forward-address>",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkFlags(cmd, "listen", "forward")

				listen, _ := cmd.Flags().GetString("listen")
				forward, _ := cmd.Flags().GetString("forward")
				err := httpproxy.TransferToServeSpecial(listen, forward)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "anchor",
		Short: "Help you access the server efficiently",
		//Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.PersistentFlags().StringP("listen", "l", "", "<listen-address>")
	rootCmd.PersistentFlags().StringP("forward", "f", "", "<forward-address>")
	rootCmd.AddCommand(tcpCmd)
	rootCmd.AddCommand(udpCmd)
	rootCmd.AddCommand(socksCmd)
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(sshCmd)
	rootCmd.AddCommand(sshPtyCmd)
	rootCmd.AddCommand(serverCmd)

	if err := rootCmd.Execute(); err != nil {
		logo.Error("", err, "")
	}

}
