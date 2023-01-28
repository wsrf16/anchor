This is a tool to help you access external servers more efficiently

## 1.功能介绍
本软件针对跨区、跨网段等网络不通场景，或需要批量网络代理场景而开发，可实现：
- 基于http协议的转发（正、反向代理）
- 基于tcp协议的转发（正、反向代理）
- 基于udp协议的转发（反向代理），udp到tcp的伪装
- 基于socks5协议的代理
- 基于ssh协议的转发以及建立隧道，可用于通过ssh协议中转的方式搭建ssh代理机或http代理机
- 搭建http服务器，以http接口形式调用服务端shell脚本，或以服务器为中转机，访问远端ssh服务器
- 以伪终端的形式，访问远程ssh服务器，类似putty、xshell等工具





## 2.示意图
![img.png](img.png)


## 3.启动方式
### 3.1 命令行模式
适用于单一端口转发，启动方式简单，无需额外配置，可以通过无参数方式执行，查看帮助：
```
Help you access the server efficiently

Usage:                                                                  
  anchor [flags]                                                        
  anchor [command]                                                      
                                                                        
Available Commands:                                                     
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command                                    
  http        Start a http server         
  link        link to nat server
  nat         Start a nat server                              
  pty         Login ssh server                                         
  server      Start a anchor server                                     
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

#### 3.1.1 http转发
该类转发在应用层实现，仅适用于http协议，因为https需要证书。

·正向代理
```
# 将本机作为代理服务器。其他机器可以通过设置代理为192.168.0.100:8081访问其他网络（假设该服务器ip为192.168.0.100）
$ anchor http -L :8081
```

·反向代理
```
# 将本地8081端口接收到的http请求，转发到http://192.168.0.10:8081
$ anchor http -L :8081 -R http://192.168.0.10:8081
```

#### 3.1.2 tcp转发
该类转发在会话层实现，支持http、https、ssh等大部分基于tcp的协议。

·正向代理
```
# 将本机作为代理服务器。其他机器可以通过设置代理为192.168.0.100:8081访问其他网络（假设该服务器ip为192.168.0.100）
$ anchor tcp -L :8081
```

·反向代理
```
# 将本地8081端口接收到的tcp请求，转发到192.168.0.10的8081端口
$ anchor tcp -L :8081 -R 192.168.0.10:8081
```

#### 3.1.3 udp转发
该类转发在会话层实现，支持udp协议。

·反向代理
```
# 将本地8081端口接收到的udp请求，转发到192.168.0.10的8081端口
$ anchor udp -L :8081 -R 192.168.0.10:8081
```

#### 3.1.4 socks代理
该类转发在会话层实现，支持http、https、ssh、ftp等大部分基于tcp的协议。

·正向代理
```
# 将本地8081端口接收到的tcp请求，转发到192.168.0.10的8081端口
$ anchor socks -L :8081 -R 192.168.0.10:8081
```

#### 3.1.5 建立ssh隧道


#### 3.1.6 搭建http服务器，以http形式执行shell或访问远程ssh
本模式由于参数较多，仅支持配置文件方式启动。

#### 3.1.7 访问远程ssh服务器
```
$ pty 192.168.0.10 -u root -p 12345678
```

#### 3.1.8 内网穿透
##### 3.1.8.1 在外网服务器执行（内网ip：192.168.0.104 外网ip：10.172.0.104）
$ anchor nat -L 192.168.0.4:9090 -R 192.168.0.4:9091

##### 3.1.8.2 在内网服务器执行（内网ip：192.168.0.105）
$ anchor link -L 192.168.0.104:9091 -R 127.0.0.1:7777

##### 3.1.8.3 在外网客户端
在外网访问10.172.0.104:9090，即相当于访问内网192.168.0.105:7777

### 3.2 配置文件模式
以配置文件方式启动，一次启动同时支持多种转发方式，local为本地监听地址（必填），remote为转发目标地址（非必填）。

```
$ ./anchor server
```
$ cat config.yaml
```
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
  3ks33wGHRey/Rh3nM+CpfWyQkR34PVQP00A/L/9D5PuZ0kaU+mmt9am/iefouSJP
  uweo+jJaTMwZIaR5qbu+UOh/9FrvfCZQqBgslrIzVE37Lk58zCWRG9opY85RvQJr
  hQbyyYCSeKE/Pm+SqZO3ZXj4pvCH3YRSudW0sOLaiyro29I/Ttk0u3ATE3YnJu3t
  BSOI9WKJN259xXAZFJCYQeJIP74B8RSTpzDXVwIDAQABAoIBAQDU5nC/Gaqa7bpe
  wNwTh0FOX2nhv0txrjN+mDSPPKuB6C6XiNEnZEf0tO1HvkRer2Hb+t+tZRnT0cPM
  5s/qIQJM1HkE0R6tuf1TXKMSmb4Xl7Y1u+45udMgNw+jnHDW4JMu0/5cLviw7RvY
  q6Mgp7resCPUlXnKZM/Wlgoaa1+7dMY+RGHLznWVaYAwdgvR5sjmtfobmrHaIt4E
  e0XL7DdbtrLFhK2WVXyDlmB5Y/ji7UajRNTeZVEhihgUdBkTiyBMlJ8jA3sBG/7s
  EL+HS4WmjaoIh86wvUeh+Ue/R16p9CGfNS4QHocLM26qK3MwxOEb/il8un1/lOVB
  x7OWlYJZAoGBAO5URQSJsqCk2Dgitj+A3yPF650rsDsMhJcAoOtHPlwae/8RNh/9
  eNwqxPNrCqr+nbf4Eo22IxWmAUSRtq4NcWD0mhFyBb+MarD8P8NDByzllGlXAQok
  HK68KK6dmfnFMIa6seIJkQ2H0scP17GK/MWsjrKA6/+k4kuVwqLJkH2TAoGBAPf3
  hGb7M3BZqMWn/n6CWCyD3Jhtpfae9SZolTBrrC32uW+YxrLcJQcHQ4AqbYQihZap
  jnYUNUTB5ZpCB8Xt9vkmrJrkvvVmeonw8ebDNQbCHS0btU+06OBtwHahCj4mOLxI
  X4OtYKIoqyZzF1psPBP3i1jXmeA/3mpQM7TA+PmtAoGBAMq0mAM1Dt39dv2/4FyA
  wslNZInIh9qhWgCJA1U5wjtaN06JmAx7bRltsxgKWgd4CjNX5ofre0nFaNymRsch
  HebGBGiEH4l8553i3mS4lNjF54X0/yhd8JtZ5Awodbbv1ccx8Wv8kQf+FvNCcUVL
  npSOoMwIRTGpTuPU27fTUnbLAoGBANvRFxV3oGz7fmrh9vBNOj8UEc0qfpJ27ZnD
  4mVPPg5tBowQYiySjVyCqusGMZmOROB/pOO2HFa8kH7EBEIJJwRNrXuORl4jH4QD
  DV5lNC+uitC/HGbxV/QENzayofLgTTFT8yA25VA7PyXRpyUoh8Wnj2cNkcJCvDeA
  q6WT3+LtAoGAKHhFUXtSznBtajKaAc36rNR7V29/RctCTKWMHG9BrMXmrvOnuJw1
  r6/cWM1g7ez4qEsgzk0rRVRhsIoq01sXf4dc2N7Un2fgkmQOSqX7qqYBQgJUEFjB
  F4kdbHpJWTnF2l70iLkGIBu8Pxs7sUK658Q81iGJ/rvvaC8XAR5WM/M=
  -----END RSA PRIVATE KEY-----"
```

### 3.3 搭建http服务器
这里详细说一下[搭建http服务器](#httpserver)，以下配置内容说明：监听本地8080端口，客户端可访问本地8080端口执行shell命令，或通过该服务器访问其他ssh服务器执行shell命令。
> 除支持linux平台外，同时也支持windows平台的dos命令

#### 3.3.1 第一步：搭建http服务器
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

#### 3.3.2 远程调用ssh
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

