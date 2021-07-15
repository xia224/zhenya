# Introduce kube-schedule module
Control plane component that watches for newly created Pods with no assigned node, and selects a node for them to run on.

Factors taken into account for scheduling decisions include: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference, and deadlines.

上述翻译为：
调度器为新创建的pod选择一个节点来运行程序。
调度策略丰富多样：有单一或组合资源需求的，有软硬件以及策略限制的，有喜好或反感标识的，有数据本地化，有内部负载干扰，
截止日期。

## 节点选择
### nodeSelector 最简单推荐形式，或者nodeName指定node
nodeSelector 是 PodSpec 的一个字段。 它包含键值对的映射。为了使 pod 可以在某个节点上运行，该节点的标签中 必须包含这里的每个键值对（它也可以具有其他标签）。
步骤一 
执行 kubectl label nodes <node-name> <label-key>=<label-value> 命令将标签添加到你所选择的节点上。
步骤二
在POD的配置文件里添加一个nodeSelector部分

当你之后运行 kubectl apply -f https://k8s.io/examples/pods/pod-nginx.yaml 命令， Pod 将会调度到将标签添加到的节点上。
