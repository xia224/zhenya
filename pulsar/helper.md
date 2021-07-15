

## pulsar standalone
A pulsar broker + Zookeeper && Bookeeper components which running inside of s single JVM process.
### Install Java 8
[Oracle download center](https://www.oracle.com/index.html)

## Install pulsar using binary release
[pulsar 2.7.1 binary release](https://archive.apache.org/dist/pulsar/pulsar-2.7.1/apache-pulsar-2.7.1-bin.tar.gz)

## Untar tarball, it contains the below dir
bin: pulsar's command-line tools. such as pulsar and pulsar-admin
conf: configuration files for pulsar, including broker and zookeeper
data: the data storage directory used by zookeeper and bookeeper agter running pulsar

## Install builtin connectors(Optional)
[pulsar IO connectors 2.7.1 release](https://archive.apache.org/dist/pulsar/pulsar-2.7.1/connectors)
untar && mv to pulsar/connectors

## Install tiered storage offloaders(Optional)
[pulsar tiered storage offloaders 2.7.1 release](https://archive.apache.org/dist/pulsar/pulsar-2.7.1/apache-pulsar-offloaders-2.7.1-bin.tar.gz)
untar && mv to pulsar/offloaders

## Start pulsar standalone
bin/pulsar standalone

## Consume a message
bin/pulsar-client consume my-topic -s "first-subscription"

## Produce a message
bin/pulsar-client produce my-topic --messages "hello-pulsar"
