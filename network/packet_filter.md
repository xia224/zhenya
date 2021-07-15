# Packet Filter

## Basic iptables/netfilter
prerouting, input, output, forward, postrouting;
1. 重定向：
iptables -t nat -I OUTPUT --src 0/0 --dst 192.168.1.5 -p tcp --dport 80 -j REDIRECT --to-ports 8123

2. 允许private ip to communicate with external public networks:
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE 

3. 允许公网用户访问自己搭建的私有http服务器
iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 80 -j DNAT \
      --to 172.31.0.23:80
iptables -A FORWARD -i eth0 -p tcp --dport 80 -d 172.31.0.23 -j ACCEPT

## DPDK
Development kit from intel.

## BPF or eBPF

## XDP
