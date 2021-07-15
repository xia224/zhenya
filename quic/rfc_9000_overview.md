# Log about reading RFC-9000 QUIC connection protocol

## Key points
1. quic 基于udp传输的连接层协议，client和server有状态交互
2. client可选立即发送data（0-RTT）
3. quic packets 包含frame
4. 应用层通过使用quic的streams交换信息，有两种类型的stream， bidirectional and unidirectional
5. 可靠传输和拥塞控制方面，quic提供feedback， 拥塞控制算法：QUIC-RECOVERY
6. quic连接不强绑定到单一的网络path，允许迁移到新的path
7. 提供多种终止连接的方式

### Stream related
1. quic不保证不同流中的字节有序， 多条流可同时发送
2. stream ID 62为正整数（0到2^62 - 1）且唯一
3. streamm ID 的最低位（0x01）奇数代表server端初始化的stream，否则偶数代表client初始化的stream
4. streamm ID 第二的最低位 0x02 偶数代表bidirectional，奇数代表unidirectional
5. stream frame封装data， 端拥stream ID和stream frame的偏移量来按序放置数据

### Why to use quic
QUIC 协议选择了 UDP，因为 UDP 本身没有连接的概念，不需要三次握手，优化了连接建立的握手延迟，同时在应用程序层面实现了 TCP 的可靠性，TLS 的安全性和 HTTP2 的并发性，只需要用户端和服务端的应用程序支持 QUIC 协议，完全避开了操作系统和中间设备的限制。

### congestion control algorithm
TCP 的拥塞控制实际上包含了四个算法：慢启动，拥塞避免，快速重传，快速恢复。

QUIC 协议当前默认使用了 TCP 协议的 Cubic 拥塞控制算法，同时也支持 CubicBytes, Reno, RenoBytes, BBR, PCC 等拥塞控制算法。

### quic stream滑动窗口
QUIC就算此前有些 packet 没有接收到，它的滑动只取决于接收到的最大偏移字节数。
针对stream的可用窗口 = 最大窗口数 - 接收到的最大数据偏移数
connection的可用窗口 = 各stream的可用窗口之和

## QUIC压缩算法

## QUIC连接迁移
那 QUIC 是如何做到连接迁移呢？很简单，任何一条 QUIC 连接不再以 IP 及端口四元组标识，而是以一个 64 位的随机数作为 ID 来标识，这样就算 IP 或者端口发生变化时，只要 ID 不变，这条连接依然维持着，上层业务逻辑感知不到变化，不会中断，也就不需要重连。

由于这个 ID 是客户端随机产生的，并且长度有 64 位，所以冲突概率非常低。