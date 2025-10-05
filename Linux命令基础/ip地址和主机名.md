# ip地址和主机名

## 查看ip地址

使用ifconfig,找到ens333网卡，后面显示的就是当前的IP地址（ipv4）

格式为：a.b.c.d,其中abcd是0~255的数字

特殊ip有：127.0.0.1,本地回环ip,表示本机

0.0.0.0也表示本机，也可以在一些白名单中表示任意ip

## 修改主机名

可以通过命令：hostname查看主机的主机名

可以通过命名hostnamectl set-hostname 主机名，修改主机名（root）

更新完之后，重新登录FinalShell



## 域名解析

通过IP地址映射主机名，进行链接



## Linux固定IP地址配置

### 在VMware Workstation中配置固定ip

1. 在VMware Workstation中配置ip地址网管和网段（IP地址的范围）
2. 在Linux系统中手动修改配置文件，固定ip

使用vim编译器编辑/etc/sysconfig/network-scripts/ifcfg-ens33

```
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="8d1078b8-d3ac-4533-b118-34e2569f4511"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="192.168.88.130"
NETMASK="255.255.255.0"
GATEWAY="192.168.88.2"
DNS1="192.168.88.2"
```

改完之后重启一下网卡就行，systemctl restart network