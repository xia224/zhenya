# Main Info for message system

## Request about message system
* 消息可靠性
    如何避免因消息丢失所造成状态不一致？
* 消息实时性
    第一时间通知状态变更事件
* 消息顺序性
    同一资源的事件需要带标记，用于区分
* 高性能

## kafka

### kafka consumer group 状态机
consumer是依赖于broker端的coordinator来完成组的管理的；
coordinator 实现组的管理，依赖的主要是consumer group的状态，仅有 Empty（组内没有任何active consumer）、PreparingRebalance（group 正在准备进行rebalance）、AwaitingSync（所有组员已经加入组并等待leader consumer发送分区的分配方案）、Stable（group开始正常消费）、Dead（该group 已经被废弃）这五个状态。

### kafka rebalance
rebalance也就是如何达成一致来分配订阅topic的所有分区。这个rebalance的代价还是不小的，我们是需要避免高频的rebalance的。常见的rebalance 场景有：新成员加入组、组内成员崩溃（这种场景无法主动通知，需要被动的检测才行，并且需要一个session.timeout 才检测到）、成员主动离组。

### kafka 消费数据考虑的问题
offset：
    * Last Committed Offset：consumer group最新一次 commit 的 offset，表示这个 group 已经把 Last Committed Offset 之前的数据都消费成功了。
    * Current Position：consumer group 当前消费数据的 offset，也就是说，Last Committed Offset 到 Current Position 之间的数据已经拉取成功，可能正在处理，但是还未 commit。
    * Log End Offset(LEO)：记录底层日志(log)中的下一条消息的 offset。,对producer来说，就是即将插入下一条消息的offset。
    * High Watermark(HW)：已经成功备份到其他 replicas 中的最新一条数据的 offset，也就是说 Log End Offset 与 High Watermark 之间的数据已经写入到该 partition 的 leader 中，但是还未完全备份到其他的 replicas 中，consumer是无法消费这部分消息(未提交消息)。