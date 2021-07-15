# introduction packets

## packet number
*使用递增的packet number保证可靠性， 类似于tcp的sequence
*QUIC 同样是一个可靠的协议，它使用 Packet Number 代替了 TCP 的 sequence number，并且每个 Packet Number 都严格递增，也就是说就算 Packet N 丢失了，重传的 Packet N 的 Packet Number 已经不是 N，而是一个比 N 大的值。
 根据ACK包的packet number可以确定RTT（没有歧义的那种的）

## Stream offset
*单纯依靠严格递增的Packet Number 肯定是无法保证数据的顺序性和可靠性。QUIC 又引入了一个 Stream Offset 的概念。
*一个 Stream 可以经过多个 Packet 传输，Packet Number 严格递增，没有依赖。
但是 Packet 里的 Payload 如果是 Stream 的话，就需要依靠 Stream 的 Offset 来保证应用数据的顺序。
Stream （offset:x) 和 Stream (offset:x+y) 按照顺序组织起来，交给应用程序处理。

## 基于stream和connection两类级别的流量控制
Quic支持多路复用，类比一条Connection上同时存在多条stream；
Stream类比为一条http请求；
Connection类比一条tcp连接；

## 没有对头阻塞的多路复用
QUIC的多路复用类似于HTTP2，都是在一条连接上并发发送多个http请求（stream）。
quic的优势是：一个连接上的多个stream之间没有依赖，即使stream2丢了udp packet，
也仅仅影响stream2的处理，对其他stream没有影响。

### HTTP2的情况：
HTTP2 在一个 TCP 连接上同时发送 4 个 Stream。其中 Stream1 已经正确到达，并被应用层读取。但是 Stream2 的第三个 tcp segment 丢失了，TCP 为了保证数据的可靠性，需要发送端重传第 3 个 segment 才能通知应用层读取接下去的数据，虽然这个时候 Stream3 和 Stream4 的全部数据已经到达了接收端，但都被阻塞住了。
不仅如此，由于 HTTP2 强制使用 TLS，还存在一个 TLS 协议层面的队头阻塞
Record 是 TLS 协议处理的最小单位，最大不能超过 16K，一些服务器比如 Nginx 默认的大小就是 16K。由于一个 record 必须经过数据一致性校验才能进行加解密，所以一个 16K 的 record，就算丢了一个字节，也会导致已经接收到的 15.99K 数据无法处理，因为它不完整。

### QUIC如何解决对头阻塞问题的呢
1. QUIC最基本的传输单元是Packet，不会超过MTU的大小，加密和认证过程也是基于packet的，
不会跨越多个Packet，这样就能避免TLS协议存在的阻塞。
2. Stream 之间相互独立，比如 Stream2 丢了一个 Pakcet，不会影响 Stream3 和 Stream4。不存在 TCP 队头阻塞。