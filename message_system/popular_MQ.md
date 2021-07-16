# 介绍常用的MQ
解耦、异步、削峰
## Kafka
[官方设计文档](https://kafka.apache.org/documentation/#design)
1. 基础知识介绍
* Partition结构:
    partition(分区)表现为服务器上的一个一个文件夹；
    partition文件夹下包含多组segment文件；
    每组segment文件包含：.index, .log and .timeindex文件；
    .log文件存储message，其他两个为索引文件。
    分区可提高读写性能；
* Replicated冗余:
    容错和高可用需求；
    冗余是topic的分区粒度的；
* Producer:
    ACK: acks = 0 or 1 or all(等待master sync到所有replicas之后，再确认，写成功)
    --bootstrap-server broker_host:port
* Consumer：
    One topic divided into partitions, every consumer has a single integer per partition.
* Persistence store:
    Log segment flushed to disk: lazily flush;
    Rely on partition replication;
    Per-topic configuration parameters;
    Quote: `When a Broker receives data from a Producer, it is immediately written to a persistent log on the filesystem, however this does not mean it will be flushed to disk. The data will be transferred into the kernel’s page cache and it will be up to the operating system to decide when the flush should happen, i.e depending on the configured vm.dirty_ratio, vm.dirty_background_ratio and vm.swappiness kernel parameters.`
    reads and appends to files for logging, read do not block writes or each other.
    Storage systems mix fast cached operations with slow physical disk operations,底层还使用了平衡B树，用于segment查询。
* Zero-copy：
    传统：从disk文件读取内容，然后通过internet发送出去需要4次copy，disk --> read_buffer --> application --> socket_buffer --> nic;
    零拷贝：直接从read_buffer拷贝到socket_buffer，且都在内核态；通过unix系统调用函数sendfile(),在两个fd之间拷贝数据。
    网卡如果支持gather操作的话，数据从read_buffer --> nic;
* Efficiency：
    多次写可以被消费者一次获取到，引入消息集抽象，避免多次小的IO disk操作；
    避免没效率的字节拷贝：引入标准化的binary消息格式，在producer，broker和consumer之间共享。
* 端到端批量压缩：
    有效的压缩：同时压缩多条消息， 而不是单个消息压缩。
    消息压缩形式存在于log文件中，消费者收到后会解压缩。
* 推/拉流模式：
    kafka follows大部分成熟消息系统的设计模式：producer推消息到broker， consumer从broker拉取消息。（主动推送到consumer的模式的问题：消费端是多种多样的）
* Semantics:
    `
    At most once—-Messages may be lost but are never redelivered.
    At least once—-Messages are never lost but may be redelivered.
    Exactly once—-this is what people actually want, each message is delivered once and only once.
    `
    Cases:
    At-Least-Once semantic: Producer failed to receive a response from broker, and it resend the message.
    At-Most-Once semantic: Consumer read messages, then save its position, and finally process messages, it process crashes after saving its position but before process messages.
    At-Least-Once semantic: Consumer read messages, then process messages, finally save its position, it process crashes after process messages but before saving its position.
    Exactly-Once semantic: when consuming from topic and producing to another topic.
* Node liveness
    Node maintain its session with ZooKeeper;
    Follower not fall  "too far" behind; 
    the list of "in sync" replicas;
* 消息顺序：
    topic是无序的；
    每个topic的partition内部是有序的；
* 创建topic:
    `bin/kafka-topics.sh --create --zookeeper kafka-00:21810, kafka-01:21810 --replication-factor 2 --partitions 3 --topic mytopic2`
* 查看broker个数:
    `zookeeper-shell localhost:2181 ls /brokers/ids`

2. 常见使用
* Expand kafka cluster:
    新的servers上不会自动被分配已有topic的partition，新的topic会；
    你可以尝试这主动迁移数据到新servers上；
* Datacenters:
    建议：每个数据中心部署一套本地kafka集群，集群间mirroring镜像数据；
* 消息挤压的监控报警：

3. 常见问题及解决方案
    * 同一机器上多次创建broker server，需要清理掉/tmp/kafka.logs/meta.properties文件；
    * 消费者消费数据失败（网络不稳定），且消费的数据有前后依赖关系，如何处理？
        同步重试机制--影响消费进度；
        异步重试机制--需要保存整个消息相关的信息，重试N次后通知相关人员；
    * 消息积压 （消费者消费不过来）
        消息从生产到消费至少需要2次网络IO和2次磁盘IO；
        消息体过大造成生产和消费过程缓慢，会出现积压；// 消息体中只包含搜索关键key，便于后续查询；


4. 其他
    用go实现kafka的producer和consumer库："github.com/Shopify/sarama"
## RabbitMQ

## RocketMQ

## Pulsar(云原生)

## MQ常见问题，以及解决方法
