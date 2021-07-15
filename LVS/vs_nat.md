# Study how to use vs-nat ipvs

## Help Doc url
[vs_NAT](http://www.linuxvirtualserver.org/VS-NAT.html)

## Kernel network tools
旧的方式
1. ipfwadm 
config route, e.g. ipfwadm -F -a m -S 172.16.0.0/24 -D 0.0.0.0/0
2. ippfvsadm
config ipvs rules, e.g.
ippfvsadm -A -t 202.103.106.5:80 -R 172.16.0.2:80 -w 1
ippfvsadm -A -t 202.103.106.5:80 -R 172.16.0.3:8000 -w 2
ippfvsadm -A -t 202.103.106.5:21 -R 172.16.0.3:21 

kernel 2.2.x以后的版本
3. ipchains
echo 1 > /proc/sys/net/ipv4/ip_forward
    ipchains -A forward -j MASQ -s 172.16.0.0/24 -d 0.0.0.0/0
4. ipvsadm
增加虚拟服务器，并指定调度策略
ipvsadm -A -t 202.103.106.5:80 -s wlc  (Weighted Least-Connection scheduling)
ipvsadm -A -t 202.103.106.5:21 -s wrr  (Weighted Round Robing scheduling )
具体选择rules：
ipvsadm -a -t 202.103.106.5:80 -r 172.16.0.2:80 -m 
ipvsadm -a -t 202.103.106.5:80 -r 172.16.0.3:8000 -m -w 2 
ipvsadm -a -t 202.103.106.5:21 -r 172.16.0.2:21 -m 