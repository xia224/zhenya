# 介绍常用的MQ
解耦、异步、削峰
## Kafka
1. 基础知识介绍
* Partition结构
    partition(分区)表现为服务器上的一个一个文件夹；
    partition文件夹下包含多组segment文件；
    每组segment文件包含：.index, .log and .timeindex文件；
    .log文件存储message，其他两个为索引文件。
    分区可提高读写性能；
* Replicated冗余
    容错和高可用需求；
    冗余是topic的分区粒度的；
* Producer
    ACK: acks = 0 or 1 or all(等待master sync到所有replicas之后，再确认，写成功)
    --bootstrap-server broker_host:port
* Persistence store
    Log segment flushed to disk: lazily flush;
    Rely on partition replication;
    Per-topic configuration parameters;
    Quote: `When a Broker receives data from a Producer, it is immediately written to a persistent log on the filesystem, however this does not mean it will be flushed to disk. The data will be transferred into the kernel’s page cache and it will be up to the operating system to decide when the flush should happen, i.e depending on the configured vm.dirty_ratio, vm.dirty_background_ratio and vm.swappiness kernel parameters.`
* Zero-copy
    传统：从disk文件读取内容，然后通过internet发送出去需要4次copy，disk --> read_buffer --> application --> socket_buffer --> nic;
    零拷贝：直接从read_buffer拷贝到socket_buffer，且都在内核态；通过unix系统调用函数sendfile(),在两个fd之间拷贝数据。
    网卡如果支持gather操作的话，数据从read_buffer --> nic;

2. 常见使用常见

3. 常见问题及解决方案

4. 其他
    用go实现kafka的producer和consumer库："github.com/Shopify/sarama"
## RabbitMQ

## RocketMQ

## Pulsar(云原生)

## MQ常见问题，以及解决方法
###