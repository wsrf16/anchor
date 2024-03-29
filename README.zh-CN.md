# Anchor

[![EN doc](https://img.shields.io/badge/document-English-blue.svg)](README.md) | 简体中文

[toc]




## 1.功能介绍

本软件针对跨区、跨网段等网络不通场景而开发，具备以下功能：
- [x] http协议：基于http协议的转发、正反向代理服务器
- [x] tcp协议：基于tcp协议的转发、正反向代理服务器
- [x] udp协议：基于udp协议的转发，udp到tcp的伪装
- [x] socks协议：基于socks5协议的代理服务器
- [x] ssh协议：基于ssh协议的转发以及搭建隧道，可用于通过ssh协议中转的方式搭建ssh代理机或http代理机
- [x] shadowsocks协议：支持搭建shadowsocks代理服务器，不详细解释，你懂
- [x] tun2socks客户端：基于socks的全局透明代理客户端连接socks服务器，不再需要手动设置应用程序（如浏览器）代理
- [x] shell/bat执行机：搭建http服务器，通过http接口的形式，在服务器执行shell脚本（windows执行bat批处理），或以服务器为跳板机，访问另一台远端ssh服务器
- [x] pty客户端：以伪终端的形式，访问远程ssh服务器，类似putty、xshell等工具
- [x] 内网穿透：在家访问公司电脑（需要一台公网服务器），异地组网





## 2.示意图
![img.png](img.png)




## 3.使用说明
本程序支持Linux、macOS、Windows平台，cli采用cobra编写，因此命令行风格与kubernetes相同，当不带参数执行时，可查看帮助信息。



### 3.1 命令行启动方式
通过命令行参数的方式启动程序，支持单一协议端口转发，不带参数直接执行可查看帮助：
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
  ss          Start a shadowsocks server                                        
  tcp         Start a tcp server                                        
  udp         Start a udp server                                        
                                                                        
Flags:                                                                  
  -h, --help             help for anchor                                
  -L, --local string    <local-address>                               
  -R, --remote string   <remote-address>                              
                                                                        
Use "anchor [command] --help" for more information about a command.
```



#### 3.1.1 http转发
该类转发在应用层实现，仅适用于http协议，因为https需要证书。

·正向代理
```sh
# 将本机作为代理服务器。其他机器可以通过设置代理为192.168.0.100:8081访问其他网络（假设该服务器ip为192.168.0.100）
$ ./anchor http -L :8081
```

·反向代理
```sh
# 将本地8081端口接收到的http请求，转发到http://192.168.0.10:8081
$ ./anchor http -L :8081 -R http://192.168.0.10:8081
```



#### 3.1.2 tcp转发
该类转发在会话层实现，支持http、https协议。

·正向代理
```sh
# 将本机作为代理服务器。其他机器可以通过设置代理为192.168.0.100:8081访问其他网络（假设该服务器ip为192.168.0.100）
$ ./anchor tcp -L :8081
```

·反向代理
```sh
# 将本地8081端口接收到的tcp请求，转发到192.168.0.10的8081端口
$ ./anchor tcp -L :8081 -R 192.168.0.10:8081
```



#### 3.1.3 udp转发
该类转发在会话层实现，支持udp协议。

·反向代理
```sh
# 将本地8081端口接收到的udp请求，转发到192.168.0.10的8081端口
$ ./anchor udp -L :8081 -R 192.168.0.10:8081
```



#### 3.1.4 socks代理
该类转发在会话层实现，支持http、https、ssh、ftp等大部分基于tcp的协议。

·正向代理
```sh
# 将本机作为socks5代理服务器。其他机器可以通过设置代理为192.168.0.100:1080访问其他网络（假设该服务器ip为192.168.0.100）
$ ./anchor socks -L :1080 -U user -P 1234
```



#### 3.1.5 ssh隧道


#### 3.1.6 shadowsocks代理服务器
```sh
# 搭建shadowsocks代理服务器
$ ./anchor ss -L :8388 -P 123456 -C aes-256-gcm
```

#### 3.1.7 tun2socks
```sh
# 简单模式
$ ./anchor t2s --proxyServer 192.168.0.103:1080
# 手动模式
$ ./anchor t2s --tunName MYNIC --tunAddress 10.255.0.2 --tunGateway 10.255.0.1 --tunMask 255.255.255.0 --tunDNS 8.8.8.8,8.8.4.4 --proxyType socks --proxyServer 192.168.0.103:1080
```

#### 3.1.8 远程执行机

搭建http服务器，以http形式执行shell或访问远程ssh。本模式由于参数较多，仅支持配置文件方式启动，详细内容参考[http远程执行机](#remote-server)部分。



#### 3.1.9 访问远程ssh服务器
```sh
$ ./anchor pty 192.168.0.10 -u root -p 12345678
```



#### 3.1.10 内网穿透
##### 3.1.10.1 公网服务器（有公网ip）

环境：公网服务器（ip: 110.168.0.104） 内网服务器（ip: 192.168.0.104）

```sh
$ ./anchor nat -L :9090 -R :9091
```

##### 3.1.10.2 内网服务器
环境：内网服务器（ip: 192.168.0.105）
```sh
$ ./anchor link -L 192.168.0.104:9091 -R :1234
```

##### 3.1.10.3 公网客户端

在公网访问110.168.0.104:9090，等同于访问内网192.168.0.105:1234

```sh
$ telnet 110.168.0.104 9090
```



### 3.2 配置文件模式
通过配置文件方式获取参数，local为本地监听地址（必填），remote为转发目标地址（非必填），一次启动可同时支持多种协议转发，不带参数直接执行可查看帮助：

```sh
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

ss:
  - local: :8388
    password: 123456
    algorithm: aes-256-gcm
    tcp: true
    udp: false

t2s:
  - tunName: MYNIC
    proxyType: socks
    proxyServer: micro.com:1080
    tunAddress: 10.255.0.4
    tunGateway: 10.255.0.1
    tunMask: 255.255.255.0
    tunDNS: 8.8.8.8,8.8.4.4
    defaultGateway: 192.168.0.1

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



### 3.3 搭建远程执行机

这里详细说一下搭建[http远程执行机](#remote-server)，以下配置内容说明：监听本地8080端口，客户端可通过调用http接口的方式执行shell命令，或通过该跳板服务器访问其他ssh服务器执行shell命令。

> 除支持linux平台外，同时也支持windows平台的dos命令



#### 3.3.1 搭建shell执行机
```
httpserver:
  # 监听本地端口
  local: :8080
  # 在服务器本地执行shell命令
  shell:
    enabled: true
  # 通过服务器连接其他ssh服务器执行shell命令
  ssh:
      # 自定义，用于标记该目标地址的唯一标识
    - id: 192.168.0.10:22
      # 目标服务器地址
      addr: 192.168.0.10:22
      # 登录该服务器的用户名
      user: root
      # 登录该服务器的密码
      password: 11
      # 也可配置私钥，与密码认证方式二选一即可
      privateKey: "-----BEGIN RSA PRIVATE KEY-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END RSA PRIVATE KEY-----"
```



#### 3.3.2 搭建ssh执行机
经过以上配置后，可以通过调用以下http接口，ssh连接到id为“mecs.com:22”的机器上执行命令
> $ curl -XPOST "http://localhost:8089/ssh" -H "Content-Type: applicaton/json" -d "{\"commands\":[\"whoami\", \"aaaa\", \"curl\"],\"serverId\":\"mecs.com:22\"}"
返回以下类似内容：
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
也可在服务器本地执行shell命令，使用方法与ssh类似，此时无需指定serverId
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

