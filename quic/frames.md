# Introduction Frames
QUIC的包payload是由一系列完整的frame组成的(至少一个)，但是（Version Negotiation, Stateless Reset, and Retry packets）
版本协商、无状态重置和重试包不含frames。
Packet Payload {
    Frame (8...)...,
}
如果包不含有frame， 终端返回连接错误：protocol_violation, 一个frame不可跨越多个pkt。

Frame {
    Frame Type (i),
    Type-Dependent Fields (..),
}
## Frame type
[Types](/Users/zzy/Desktop/frame_t.png "Quic frame types table")

0-RTT：
A 0-RTT packet is used to carry "early" data from the client to the server as part of the first flight, prior to handshake completion. As part of the TLS handshake, the server can accept or reject this early data.

Handshake Packet：
It is used to carry cryptographic handshake messages and acknowledgments from the server and client.

Retry Packet：
It is used by a server that wishes to perform a retry。

1-RTT Packet:
It is used after the version and 1-RTT keys are negotiated.

### Window_update frame
告诉对端自己可以接收的字节数，这样发送方就不会发送超过这个数量的数据。

### Block frame
告诉对端由于流量控制被阻塞，无法发送数据。

### Padding frame
type=0x00
没有具体语义，仅用来填充，增加packet的大小。
譬如：增加initial包满足最小尺寸要求；保护包免遭流量分析。

### Ping frame
type=0x01
用来心跳保活

### ACK frame
type=0x02 and 0x03
接收端通知发送者它收到的和正在处理的包

### Reset_stream frame
type=0x04
终止流的发送端

### Stop_sending frame
type=0x05
终端告诉对端，不要再给我发送数据，我不收了。
即使发给我，我也会丢掉。

### Crypto frame
type=0x06
用于传输加密握手消息，不能通过0-RTT包发送

### New_token frame
type=0x07
server提供给client的token，用于未来的连接


