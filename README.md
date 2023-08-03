# Anchor

[![EN doc](https://img.shields.io/badge/document-Chinese-blue.svg)](README.zh-CN.md)


[toc]




## 1. Feature introduction

This software is developed for scenarios such as cross-region, cross-network segment and other network barriers, or scenarios that require batches of network agents, and can achieve:
- Forwarding based on http protocol (forward and reverse proxy)
- Forwarding based on tcp protocol (forward and reverse proxy)
- Forwarding (reverse proxy) based on udp protocol, camouflage from udp to tcp
- Proxy based on socks5 protocol
- Based on ssh protocol forwarding and tunnel establishment, it can be used to build ssh agent or http agent through ssh protocol transfer
- Build an http server, call the server shell script in the form of http interface, or use the server as a transit machine to access the remote ssh server
- Access the remote ssh server in the form of a pseudo-terminal, similar to putty, xshell and other tools
- Intranet penetration, access to company computers at home, remote networking




## 2. Schematic diagram
![img.png](img.png)




## 3. Instructions for use
This program supports Linux, macOS, and Windows platforms. The cli is written in cobra, so the command line style is the same as that of kubernetes. When executed without parameters, you can view the help information.



### 3.1 Command line startup
Start the program with command line parameters, support port forwarding for a single protocol, and execute directly without parameters to view the help:
```
Help you access the server efficiently

Usage:
   anchor [flags]
   anchor [command]
                                                                        
Available Commands:                                                     
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command                                    
  http        Start a http server         
  link        Link to nat server
  nat         Start a nat server                              
  pty         Login ssh server                                         
  server      Start an anchor server                                     
  socks       Start a socks server                                      
  ssh         Start a ssh server                                        
  tcp         Start a tcp server                                        
  udp         Start a udp server                                        
                                                                        
Flags:                                                                  
  -h, --help             help for anchor                                
  -L, --local string    <local-address>                               
  -R, --remote string   <remote-address>                              
                                                                        
Use "anchor [command] --help" for more information about a command.
```



#### 3.1.1 http forwarding
This type of forwarding is implemented at the application layer and is only applicable to the http protocol, because https requires a certificate.

·Forward proxy
```
# Use this machine as a proxy server. Other machines can access other networks by setting the proxy to 192.168.0.100:8081 (assuming the server ip is 192.168.0.100)
$ anchor http -L :8081
```

·Reverse proxy
```
# Forward the http request received by the local port 8081 to http://192.168.0.10:8081
$ anchor http -L :8081 -R http://192.168.0.10:8081
```



#### 3.1.2 tcp forwarding
This type of forwarding is implemented at the session layer and supports most tcp-based protocols such as http, https, and ssh.

·Forward proxy
```
# Use this machine as a proxy server. Other machines can access other networks by setting the proxy to 192.168.0.100:8081 (assuming the server ip is 192.168.0.100)
$ anchor tcp -L :8081
```

·Reverse proxy
```
# Forward the tcp request received by the local port 8081 to port 8081 of 192.168.0.10
$ anchor tcp -L:8081 -R 192.168.0.10:8081
```



#### 3.1.3 udp forwarding
This type of forwarding is implemented at the session layer and supports the udp protocol.

·Reverse proxy
```
# Forward the udp request received by the local port 8081 to port 8081 of 192.168.0.10
$ anchor udp -L:8081 -R 192.168.0.10:8081
```



#### 3.1.4 socks proxy
This type of forwarding is implemented at the session layer and supports most TCP-based protocols such as http, https, ssh, and ftp.

·Forward proxy
```
# Forward the tcp request received by the local port 8081 to port 8081 of 192.168.0.10
$ anchor socks -L:8081 -R 192.168.0.10:8081
```



#### 3.1.5 Establish ssh tunnel




#### 3.1.6 Set up http server, execute shell or access remote ssh in http form
Due to the large number of parameters, this mode only supports configuration file startup.



#### 3.1.7 Access remote ssh server
```
$ anchor pty 192.168.0.10 -u root -p 12345678
```



#### 3.1.8 Intranet Penetration
##### 3.1.8.1 Execute on the external network server (intranet ip: 192.168.0.104 external network ip: 10.172.0.104)
```
$ anchor nat -L 192.168.0.4:9090 -R 192.168.0.4:9091
```

##### 3.1.8.2 Intranet server execution (intranet ip: 192.168.0.105)
```
$ anchor link -L 192.168.0.104:9091 -R 127.0.0.1:7777
```

##### 3.1.8.3 On the external network client
Accessing 10.172.0.104:9090 on the external network is equivalent to accessing 192.168.0.105:7777 on the internal network



### 3.2 Configuration file mode
Obtain parameters through configuration files. local is the local listening address (required), and remote is the forwarding target address (not required). One-time startup can support multiple protocol forwarding at the same time. Direct execution without parameters can view the help:

```
$ ./anchor server
$ cat config.yaml
tcp:
  - local: :8081
  - local: :8082
    remote: mecs.com:8080

udp:
  - local: :8083
    remote: localhost:8084

socks:
  - local: :1080

http:
  - local: :8087
  - local: :8088
    remote: http://mecs.com:8080
    addedHead: test_header

ssh:
  - local: :8022
    remote: mecs.com:22

httpserver:
  local: :8080
  shell:
    enabled: true
  ssh:
    - id: mecs.com:22
      addr: mecs.com:22
      user: root
      password: 11
      privateKey: "-----BEGIN RSA PRIVATE KEY-----
  MIIEpQIBAAKCAQEA5tm9KUtCqjSNMqZGENzyLYj5W/8fwghZVtta1CVv0ycgMW9G
  UKRnXkHR9mrUQ38W7JvMaY2G8Z5eijvIp20YtIe/jrvgs/ZWxmAZANz/CSTI5/Jt
  F4kdbHpJWTnF2l70iLkGIBu8Pxs7sUK658Q81iGJ/rvvaC8XAR5WM/M=
  -----END RSA PRIVATE KEY-----"
```



### 3.3 Build a remote execution machine
Let’s talk about [build http server](#httpserver) in detail here, the following configuration description: monitor the local port 8080, the client can execute shell commands by calling the http interface, or access other ssh servers to execute shell commands through the springboard server.
> In addition to supporting the linux platform, it also supports the dos command of the windows platform



#### 3.3.1 Building a shell execution machine
```
httpserver:
# Listen on local port
  local: :8080
  # Execute shell commands locally on the server
  shell:
    enabled: true
  # Connect to other ssh servers through the server to execute shell commands
  ssh:
      # Custom, used to mark the unique identifier of the target address
    - id: 192.168.0.10:22
      # Target server address
      addr: 192.168.0.10:22
      # Username to log in to the server
      user: root
      # Login password for this server
      password: 11
      # You can also configure the private key, and choose one of the password authentication methods
      privateKey: "-----BEGIN RSA PRIVATE KEY-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END RSA PRIVATE KEY-----"
```
#### 3.3.2 Build ssh execution machine
After the above configuration, you can connect to the machine with the id "mecs.com:22" through ssh to execute the command by calling the following http interface
> $ curl -XPOST "http://localhost:8089/ssh" -H "Content-Type: applicaton/json" -d "{\"commands\":[\"whoami\", \"aaaa\", \"curl\"],\"serverId\":\"mecs.com:22\"}"
returns something like:
```
{
    "spanId": "02063545-70ca-11ed-8f4d-f018980ebd48",
    "code": 0,
    "msg": "success",
    "data": {
        "results": [
            {
                "stdout": "root\n",
                "stderr": ""
            },
            {
                "stdout": "bash: aaaa: command not found\n",
                "stderr": "Process exited with status 127"
            },
            {
                "stdout": "curl: try 'curl --help' or 'curl --manual' for more information\n",
                "stderr": "Process exited with status 2"
            }
        ]
    },
    "dateTime": "2022-12-01T00:14:03.8019397+08:00",
    "timestamp": 1669824843801
}
```



#### 3.3.3 调用shell
The shell command can also be executed locally on the server. The method of use is similar to that of ssh, and there is no need to specify the serverId at this time
> $ curl -XPOST "http://localhost:8089/shell" -H "Content-Type: applicaton/json" -d "{\"commands\":[\"whoami\", \"aaaa\", \"curl\"]}"
```
{
    "spanId": "96478563-70cb-11ed-86be-000c297d3626",
    "code": 0,
    "msg": "success",
    "data": {
        "results": [
            {
                "stdout": "root\n",
                "stderr": ""
            },
            {
                "stdout": "/bin/bash: aaaa: command not found\n",
                "stderr": "exit status 127"
            },
            {
                "stdout": "curl: try 'curl --help' or 'curl --manual' for more information\n",
                "stderr": "exit status 2"
            }
        ]
    },
    "dateTime": "2022-12-01T00:25:22.032626684+08:00",
    "timestamp": 1669825522032
}
```

