package cmd

import (
	"anchor/entry/server"
	"anchor/module/httpproxy"
	"anchor/support/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wsrf16/swiss/sugar/base/lambda"
	"github.com/wsrf16/swiss/sugar/logo"
	"github.com/wsrf16/swiss/sugar/netkit/socket/sockskit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/tcpkit"
	"github.com/wsrf16/swiss/sugar/netkit/socket/udpkit"
	"github.com/wsrf16/swiss/sugar/netkit/sshkit"
	"os"
	"strings"
)

func checkAllFlags(cmd *cobra.Command, flagKeys ...string) (T bool) {
	for _, flagKey := range flagKeys {
		flagVal, err := cmd.Flags().GetString(flagKey)
		check := lambda.If(len(flagVal) < 1 || err != nil, false, true)
		if !check {
			cmd.Help()
			os.Exit(1)
			return false
		}
	}
	return true
}

func hasFlag(cmd *cobra.Command, flagKey string) (T bool) {
	flagVal, err := cmd.Flags().GetString(flagKey)
	check := lambda.If(len(flagVal) < 1 || err != nil, false, true)
	if !check {
		cmd.Help()
		return false
	}
	return true
}

func Start() {
	var serverCMD = &cobra.Command{
		Use:   "server",
		Short: "Start a anchor server",
		// Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			server.Serve()
		},
	}

	var tcpCMD = &cobra.Command{
		Use:     "tcp",
		Short:   "Start a tcp server",
		Long:    "tcp -L <local-address> [-R <remote-address>]",
		Example: "tcp -L :8081 -R 192.168.0.103:8081",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				if !checkAllFlags(cmd, "local") {
					cmd.Help()
				}
				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")
				if len(local) > 0 && len(remote) > 0 {
					err := tcpkit.TransferFromListenToDialAddress(local, remote, true, nil)
					if err != nil {
						panic(err)
					}
				} else if len(local) > 0 {
					err := tcpkit.TransferFromListenAddress(local, true, nil)
					if err != nil {
						panic(err)
					}
				}
			} else {
				cmd.Help()
			}
		},
	}
	tcpCMD.Flags().StringP("local", "L", "", "<local-address>")
	tcpCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	var udpCMD = &cobra.Command{
		Use:     "udp",
		Short:   "Start a udp server",
		Long:    "udp -L <local-address> -R <remote-address>",
		Example: "udp -L :8081 -R 192.168.0.103:8081",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "local", "remote")

				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")
				err := udpkit.TransferFromListenToDialAddress(local, remote)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	udpCMD.Flags().StringP("local", "L", "", "<local-address>")
	udpCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	var socksCMD = &cobra.Command{
		Use:     "socks",
		Short:   "Start a socks server",
		Long:    "socks -L <local-address> [-U <username> -P <password>]",
		Example: "socks -L :8081 -U user -P 1234",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				if !checkAllFlags(cmd, "local") {
					cmd.Help()
				}
				local, _ := cmd.Flags().GetString("local")

				//_, err := cmd.Flags().GetString("password")
				//if err != nil {
				//    cmd.Help()
				//}

				username, _ := cmd.Flags().GetString("username")
				password, _ := cmd.Flags().GetString("password")
				config := sockskit.NewSocksConfig(username, password)

				if len(local) > 0 {
					err := sockskit.TransferFromListenAddress(local, config)
					if err != nil {
						panic(err)
					}
				}
			} else {
				cmd.Help()
			}
		},
	}
	socksCMD.Flags().StringP("local", "L", "", "<local-address>")
	socksCMD.Flags().StringP("username", "U", "", "<username>")
	socksCMD.Flags().StringP("password", "P", "", "<password>")

	var sshCMD = &cobra.Command{
		Use:     "ssh",
		Short:   "Start a ssh server",
		Long:    "ssh -L <local-address> -R <remote-address>",
		Example: "ssh -L :8081 -R 192.168.0.103:8081",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "local", "remote")

				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")
				err := tcpkit.TransferFromListenToDialAddress(local, remote, true, nil)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	sshCMD.Flags().StringP("local", "L", "", "<local-address>")
	sshCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	var sshPtyCMD = &cobra.Command{
		Use:     "pty",
		Short:   "Login ssh server",
		Long:    "pty <address> -U <username> -P <password>",
		Example: "pty 192.168.0.103:8081 -U root -P 12345678",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "username")

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
				addr := lambda.If[string](strings.Contains(args[0], ":"), args[0], args[0]+":22")
				err = sshkit.SimplePtyFlat(addr, username, password, "")
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	sshPtyCMD.Flags().StringP("username", "U", "", "<username>")
	sshPtyCMD.Flags().StringP("password", "P", "", "<password>")

	var httpCMD = &cobra.Command{
		Use:     "http",
		Short:   "Start a http server",
		Long:    "http -L <local-address> -R <remote-address>",
		Example: "http -L :8081 -R 192.168.0.103:8081",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "local", "remote")

				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")
				err := httpproxy.ListenAndServe(local, remote)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	httpCMD.Flags().StringP("local", "L", "", "<local-address>")
	httpCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	var natCMD = &cobra.Command{
		Use:     "nat",
		Short:   "Start a nat server",
		Long:    "nat -L <local-address> -R <remote-address>",
		Example: "nat -L :9090 -R :9091",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "local", "remote")

				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")

				err := tcpkit.TransferFromListenToListenAddress(local, remote)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	natCMD.Flags().StringP("local", "L", "", "<local-address>")
	natCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	var linkCMD = &cobra.Command{
		Use:     "link",
		Short:   "link to nat server",
		Long:    "link -L <local-address> -R <remote-address>",
		Example: "link -L 127.0.0.1:9091 -R 192.168.0.133:8080",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				checkAllFlags(cmd, "local", "remote")

				local, _ := cmd.Flags().GetString("local")
				remote, _ := cmd.Flags().GetString("remote")

				err := tcpkit.TransferFromDialToDialAddress(local, remote)
				if err != nil {
					panic(err)
				}
			} else {
				cmd.Help()
			}
		},
	}
	linkCMD.Flags().StringP("local", "L", "", "<local-address>")
	linkCMD.Flags().StringP("remote", "R", "", "<remote-address>")

	/*
	   anchor nat -A OUTPUT -p tcp --destination 2.2.2.2:2222 --to-destination 3.3.3.3:3333
	   -t nat
	   -j DNAT
	   iptables -t nat -A OUTPUT -p tcp -d 128.0.0.1 --dport 8080 -j DNAT --to-destination 127.0.0.1:8081

	   anchor iptables -A PREROUTING -p tcp --destination 2.2.2.2:2222 --to-destination 3.3.3.3:3333
	   -t nat
	   -j DNAT
	   iptables -t nat -A PREROUTING -p tcp -d 192.168.192.133 --dport 8081 -j DNAT --to-destination 127.0.0.1:8080

	   -A -I -D -R -L
	*/
	var iptablesCMD = &cobra.Command{
		Use:   "nat",
		Short: "Request remote",
		Long:  "nat -A <chain> -p <protocol> --destination <destination-address> --to-destination <to-destination-address>",
		Run: func(cmd *cobra.Command, args []string) {
			//if cmd.HasFlags() {
			//    if !checkAllFlags(cmd, "protocol", "destination", "to-destination") {
			//        cmd.Help()
			//    }
			//
			//    // var action = ""
			//    // var chain = ""
			//    const table = "nat"
			//    iptablesInstance, err := iptables.New()
			//    if err != nil {
			//        logo.Fatalf("", nil, "Failed to new up an IPtables intance. ERROR: %v", err)
			//    }
			//    if hasFlag(cmd, "Append") {
			//        protocol, _ := cmd.Flags().GetString("protocol")
			//        chain, _ := cmd.Flags().GetString("Append")
			//        err := iptablesInstance.Append(table, chain, "-p", protocol)
			//        if err != nil {
			//            logo.Fatalf("", nil, "Failed to List. ERROR: %v", err)
			//        }
			//    } else if hasFlag(cmd, "Insert") {
			//        // action := "-I"
			//        // chain, _ := cmd.Flags().GetString("Insert")
			//
			//    } else if hasFlag(cmd, "Delete") {
			//        action := "-D"
			//        chain, _ := cmd.Flags().GetString("Delete")
			//
			//    } else if hasFlag(cmd, "Replace") {
			//        // action = "-R"
			//        // chain, _ := cmd.Flags().GetString("Replace")
			//
			//    } else if hasFlag(cmd, "List") {
			//        action := "-L"
			//        chain, _ := cmd.Flags().GetString("List")
			//        list, err := iptablesInstance.List(table, chain)
			//        if err != nil {
			//            logo.Fatalf("", nil, "Failed to List. ERROR: %v", err)
			//        } else {
			//            logo.Info("", list, "")
			//        }
			//    } else {
			//        cmd.Help()
			//        os.Exit(1)
			//    }
			//
			//} else {
			//    cmd.Help()
			//}
		},
	}
	iptablesCMD.Flags().StringP("Append", "A", "", "<table>")
	iptablesCMD.Flags().StringP("Insert", "I", "", "<table>")
	iptablesCMD.Flags().StringP("Delete", "D", "", "<table>")
	iptablesCMD.Flags().StringP("Replace", "R", "", "<table>")
	iptablesCMD.Flags().StringP("List", "L", "", "<table>")
	iptablesCMD.Flags().StringP("protocol", "p", "", "<protocol>")
	iptablesCMD.Flags().StringP("destination", "", "", "<destination-address>")
	iptablesCMD.Flags().StringP("to-destination", "", "", "<to-destination-address>")

	var rootCMD = &cobra.Command{
		Use:   "anchor",
		Short: "Help you access the server efficiently",
		// Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.HasFlags() {
				version, _ := cmd.Flags().GetBool("version")
				if version {
					fmt.Printf("Version: %s\n", config.Version)
				} else {
					cmd.Help()
				}
			} else {
				cmd.Help()
			}
		},
	}
	//rootCMD.PersistentFlags().StringP("local", "L", "", "<local-address>")
	//rootCMD.PersistentFlags().StringP("remote", "R", "", "<remote-address>")
	rootCMD.Flags().BoolP("version", "v", false, "")
	rootCMD.AddCommand(tcpCMD)
	rootCMD.AddCommand(udpCMD)
	rootCMD.AddCommand(socksCMD)
	rootCMD.AddCommand(httpCMD)
	rootCMD.AddCommand(sshCMD)
	rootCMD.AddCommand(sshPtyCMD)
	rootCMD.AddCommand(serverCMD)
	rootCMD.AddCommand(linkCMD)
	rootCMD.AddCommand(natCMD)

	if err := rootCMD.Execute(); err != nil {
		fmt.Println()
		logo.Error("", err, "")
	}

}
