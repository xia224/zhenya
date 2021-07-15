# Introduce node components
Node components run on every node, maintaining running pods and providing the Kubernetes runtime environment.

## Kubelet
An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.
运行在集群中每个node上的代理，管理运行的容器。
The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy. The kubelet doesn't manage containers which were not created by Kubernetes.
kubelet已PodSpec为基准，确保容器健康正常运行。


## Kube-proxy
kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.
它是节点上的网络代理
kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.
它维护着网络规则，用于和集群内外通信的规则。
kube-proxy uses the operating system packet filtering layer if there is one and it's available. Otherwise, kube-proxy forwards the traffic itself.

## 亲和性和反亲和性
### 节点亲和性
* 目前有两种类型的节点亲和性，分别为 requiredDuringSchedulingIgnoredDuringExecution 和 preferredDuringSchedulingIgnoredDuringExecution。
前者指定了将 Pod 调度到一个节点上 必须满足的规则（就像 nodeSelector 但使用更具表现力的语法）， 后者指定调度器将尝试执行但不能保证的偏好。
* “IgnoredDuringExecution”部分意味着，类似于 nodeSelector 的工作原理， 如果节点的标签在运行时发生变更，从而不再满足 Pod 上的亲和性规则，那么 Pod 将仍然继续在该节点上运行。
* 节点亲和性通过 PodSpec 的 affinity 字段下的 nodeAffinity 字段进行指定。
* 新的节点亲和性语法支持下面的操作符： In，NotIn，Exists，DoesNotExist，Gt，Lt。 你可以使用 NotIn 和 DoesNotExist 来实现节点反亲和性行为，或者使用 节点污点 将 Pod 从特定节点中驱逐。