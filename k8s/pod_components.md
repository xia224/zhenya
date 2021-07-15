# Introduce Pod

## Pod间的亲和性与反亲和性
* Pod 间亲和性与反亲和性使你可以 基于已经在节点上运行的 Pod 的标签 来约束 Pod 可以调度到的节点，而不是基于节点上的标签。 规则的格式为“如果 X 节点上已经运行了一个或多个 满足规则 Y 的 Pod， 则这个 Pod 应该（或者在反亲和性的情况下不应该）运行在 X 节点”。
* 与节点不同，因为 Pod 是命名空间限定的（因此 Pod 上的标签也是命名空间限定的）， 因此作用于 Pod 标签的标签选择算符必须指定选择算符应用在哪个命名空间。
* 从概念上讲，X 是一个拓扑域，如节点、机架、云供应商可用区、云供应商地理区域等。 你可以使用 topologyKey 来表示它，topologyKey 是节点标签的键以便系统 用来表示这样的拓扑域。
* Pod 间亲和性通过 PodSpec 中 affinity 字段下的 podAffinity 字段进行指定。 而 Pod 间反亲和性通过 PodSpec 中 affinity 字段下的 podAntiAffinity 字段进行指定。

## Controller 控制器
虽然可以直接使用 Pod，但在 Kubernetes 中，更为常见的是使用控制器管理 Pod。
Deployment && StatefulSet && DaemonSet
控制器提供副本管理、滚动升级和集群级别的自愈能力。

## 了解Init容器
pod可以包含多个容器，这些容器里面可以有一个或多个先于应用容器启动的Init容器。
由于 Init 容器必须在应用容器启动之前运行完成，因此 Init 容器提供了一种机制来阻塞或延迟应用容器的启动，直到满足了一组先决条件。一旦前置条件满足，Pod 内的所有的应用容器会并行启动。

## Pod的重启策略
重启策略RestartPolicy应用于Pod内的所有容器，并且仅在 Pod 所处的 Node 上由 kubelet 进行判断和重启操作。Pod 的重启策略包括 Always、OnFailure 和 Never 三种，默认值为 Always。
kubelet 重启失效容器的时间间隔以 sync-frequency 乘以 2n 来计算，例如 1、2、4 等，最长延时 5min，并且在成功重启后的 10min 后重置该时间。

## 容器探针
探针就是由kubelet对容器执行的定期诊断， 帮助检测和保证Pod中服务正常运行。
要执行诊断，kubelet 调用由容器实现的 Handler，有三种类型的处理程序：
* ExecAction
    在容器内执行指定命令。如果命令退出时返回码为 0 则认为诊断成功。
* TCPSocketAction
    对指定端口上的容器的 IP 地址进行 TCP 检查。如果端口打开，则诊断被认为是成功的。
* HTTPGetAction
    对指定的端口和路径上的容器的 IP 地址执行 HTTP Get 请求。如果响应的状态码大于等于 200 且小于 400，则诊断被认为是成功的。

探针类型
Kubelet 可以选择是否执行在容器上运行的三种探针执行和做出反应：

* livenessProbe
    指示容器是否正在运行。如果存活探测失败，则 kubelet 会杀死容器，并且容器将受到其 重启策略 的影响。如果容器不提供存活探针，则默认状态为 Success。
    ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
            name: probe-tcp
            namespace: default
        spec:
            containers:
            - name: nginx
                image: escape/nginx-test:v1
                livenessProbe:
                    initialDelaySeconds: 5
                    timeoutSeconds: 1
                    tcpSocket:
                        port: 80
    ```
* readinessProbe
    指示容器是否准备好服务请求。如果就绪探测失败，端点控制器将从与 Pod 匹配的所有 Service 的端点中删除该 Pod 的 IP 地址。初始延迟之前的就绪状态默认为 Failure。如果容器不提供就绪探针，则默认状态为 Success。
* startupProbe
    指示容器中的应用是否已经启动。如果提供了启动探测(startup probe)，则禁用所有其他探测，直到它成功为止。如果启动探测失败，kubelet 将杀死容器，容器服从其重启策略进行重启。如果容器没有提供启动探测，则默认状态为成功Success。

## Pod 启动退出动作
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-demo
spec:
  containers:
    - name: lifecycle-demo-container
      image: nginx
      lifecycle:
        postStart:
          exec:
            command:
              - "/bin/sh"
              - "-c"
              - "echo Hello from the postStart handler > /usr/share/message"
        preStop:
          exec:
            command:
              - "/bin/sh"
              - "-c"
              - "echo Hello from the poststop handler > /usr/share/message"
```

## Deployment
应用场景
* 定义 Deployment 来创建 Pod 和 RS
* 滚动升级和回滚应用
滚动升级：
方法：kubectl edit deployment/* 编辑deployment配置文件 或者
     kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1
结果：新的rs和pods会用新的应用image生成
Deployment 可确保在更新时仅关闭一定数量的 Pods，默认情况下，它确保至少 75% 所需 Pods 是运行的，即有 25% 的最大不可用。
当第一次创建 Deployment 的时候，它创建了一个 ReplicaSet 并将其直接扩展至 3 个副本。更新 Deployment 时，它创建了一个新的 ReplicaSet ，并将其扩展为 1，然后将旧 ReplicaSet 缩小到 2，以便至少有 2 个 Pod 可用，并且最多创建 4 个 Pod。然后，它继续向上和向下扩展新的和旧的 ReplicaSet ，具有相同的滚动更新策略。最后，将有 3 个可用的副本在新的 ReplicaSet 中，旧 ReplicaSet 将缩小到 0。

回滚应用：
```bash
# 检查Deployment修改历史
$ kubectl rollout history deployment/nginx-deployment
deployments "nginx-deployment"
REVISION    CHANGE-CAUSE
1           kubectl apply --filename=https://k8s.io/examples/controllers/nginx-deployment.yaml --record=true
2           kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.9.1 --record=true
3           kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.91 --record=true

# 现在已决定撤消当前展开并回滚到以前的版本
$ kubectl rollout undo deployment/nginx-deployment
deployment/nginx-deployment

# 通过下面命令来回滚到特定修改版本
$ kubectl rollout undo deployment/nginx-deployment --to-revision=2
deployment/nginx-deployment
```
* 扩容和缩容
```bash
# 扩容(3->10)
$ kubectl scale deployment/nginx-deployment --replicas=10
deployment.apps/nginx-deployment scaled

# 水平缩放
$ kubectl autoscale deployment/nginx-deployment --min=10 --max=15 --cpu-percent=80
deployment/nginx-deployment scaled
```
* 暂停和继续 Deployment

## DaemonSet
* DaemonSet 确保全部或者一些 Node 上运行一个 Pod 的副本。当有 Node 加入集群时，也会为他们新增一个 Pod 。当有 Node 从集群移除时，这些 Pod 也会被回收。删除 DaemonSet 将会删除它创建的所有 Pod。
* 简单的用法是为每种类型的守护进程在所有的节点上都启动一个 DaemonSet。一个稍微复杂的用法是为同一种守护进程部署多个 DaemonSet；每个具有不同的标志， 并且对不同硬件类型具有不同的内存、CPU 要求。
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd-elasticsearch
  namespace: kube-system
  labels:
    k8s-app: fluentd-logging
spec:
  selector:
    matchLabels:
      name: fluentd-elasticsearch
  template:
    metadata:
      labels:
        name: fluentd-elasticsearch
    spec:
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
      containers:
        - name: fluentd-elasticsearch
          image: quay.io/fluentd_elasticsearch/fluentd:v2.5.2
          resources:
            limits:
              memory: 200Mi
            requests:
              cpu: 100m
              memory: 200Mi
          volumeMounts:
            - name: varlog
              mountPath: /var/log
            - name: varlibdockercontainers
              mountPath: /var/lib/docker/containers
              readOnly: true
      terminationGracePeriodSeconds: 30
      volumes:
        - name: varlog
          hostPath:
            path: /var/log
        - name: varlibdockercontainers
          hostPath:
            path: /var/lib/docker/containers
```

## StatfulSet
管理有状态应用
StatefulSet 不仅能管理 Pod 的对象，它还能够保证这些 Pod 的顺序性和唯一性，其会为每个 Pod 设置一个单独的持久标识 ID 号。

1. 如果是不需额外数据依赖或者状态维护的部署，或者 replicas 是 1，优先考虑使用 Deployment；
2. 如果单纯的要做数据持久化，防止 pod 宕掉重启数据丢失，那么使用 pv/pvc 就可以了；
3. 如果要打通 app 之间的通信，而又不需要对外暴露，使用 headlessService 即可；
4. 如果需要使用 service 的负载均衡，不要使用 StatefulSet，尽量使用 clusterIP 类型，用 serviceName 做转发；
5. 如果是有多 replicas，且需要挂载多个 pv 且每个 pv 的数据是不同的，因为 pod 和 pv 之间是一一对应的，如果某个 pod 挂掉再重启，还需要连接之前的 pv，不能连到别的 pv 上，考虑使用 StatefulSet；
6. 能不用 StatefulSet，就不要用；

## Job/CronJob
任务计划/周期性任务计划
```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: hello
              image: busybox
              args:
                - /bin/sh
                - -c
                - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```

## 让pod执行一条命令
kubectl exec -it $POD_NAME -n $POD_NAMESPACE -- /nginx-ingress-controller --version
