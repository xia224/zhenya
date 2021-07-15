# 介绍agora universal transport AUT

## 协议header生成脚本
[protocol开源工具](https://github.com/xia224/protocol)
例如： ./protocol "type:8,pkt_no:24,payload:32"
0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|      type     |                     pkt_no                    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            payload                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

type：packet type
pkt_no：packet sequence
payload：application payload

## Frame type
帧（Frame）是AUT内携带数据的最小单元，每个Frame根据自己的类型携带不同的信息：
StreamFrame：携带应用层数据
AckFrame：携带收到对端发出的包的信息。每个需要重传的数据包，都会被Ack，Ack根据网络情况，会动态调整是否延迟
ControlFrame：携带针对流的控制信息，如：
    WindowUpdate：表示接收端的流量控制接收窗口更新
    Finish：表示发送端数据发送完毕
    Blocked：表示发送端因流量控制阻塞而不能发送
CongestionFeedbackFrame：携带本方发送端的一些网络数据，如丢包率/发送带宽/抖动等，会定期发送，使用者可以设置最小发送间隔
CloseFrame：携带关闭某条流或整个Session的错误码和说明
除AckFrame外，所有的Frame都为可重传帧（Retransmittable Frame），AckFrame不可重传（Nonretransmittable Frame）。