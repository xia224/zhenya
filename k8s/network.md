# 希望了解从接收到用户的请求，到选择node，到选择service的pod，再到返回给用户这整个流程的网络走向

## 如何选择node

## 服务发现：如何选择service或deployment
Kubernetes支持两种主要的发现Service的模式 - 环境变量和DNS。
### 环境变量
当一个Pod运行在一个Node上的时候，kubelet会为每一个活跃的Service添加一组环境变量。它同时支持Docker links compatible变量和简单的如{SVCNAME}_SERVICE_HOST和{SVCNAME}_SERVICE_PORT一样的变量，在后边的这些变量名中服务的名称是大写的，然后横线（-）会被转换成下划线（_）。
例如，服务"redis-master"开放了6379的TCP端口，并且被分配了10.0.0.11的集群IP地址，那么生成的变量如下：
```bash
REDIS_MASTER_SERVICE_HOST=10.0.0.11
REDIS_MASTER_SERVICE_PORT=6379
REDIS_MASTER_PORT=tcp://10.0.0.11:6379
REDIS_MASTER_PORT_6379_TCP=tcp://10.0.0.11:6379
REDIS_MASTER_PORT_6379_TCP_PROTO=tcp
REDIS_MASTER_PORT_6379_TCP_PORT=6379
REDIS_MASTER_PORT_6379_TCP_ADDR=10.0.0.11
```
这确实隐含了一个对顺序的要求 - 任何一个Pod想要访问的Service必须在Pod自身被创建之前被创建好，否则环境变量无法被设置到这个Pod中去。DNS没有这种限制。

### DNS
一个可选的（也是我们强烈建议的）cluster add-on是一个DNS服务器。DNS会监控Kubernetes的新的Service并且会为其创建一组DNS记录。如果DNS在整个集群中都被启用的话，那所有的Pod应该能自动对Service进行命名解析。

例如，如果你有一个叫做my-service的Service，其位于"my-ns"的Namespace下，那么对于"my-service.my-ns"就会有一个DNS的记录被创建。位于名为"my-ns" Namespace下的Pod可以通过简单地使用"my-service"进行查找。而位于其他Namespaces的Pod必须把查找名称指定为"my-service.my-ns"。这些命名查找的结果是一个集群的IP地址。

对于命名的端口，Kubernetes也支持DNS SRV (service)记录。如果名为"my-service.my-ns"的Service 有一个协议为TCP的名叫"http"的端口，你可以对"_http._tcp.my-service.my-ns"做一次DNS SRV查询来发现”http”的端口号。


## 如何选择pods

## 虚拟IP和服务代理
* 每个集群中的node之上都运行一个kube-proxy，负责实现虚拟IP
* k8s v1.0 proxy运行在用户态，v1.1 新增了iptables proxy，但不是default的，
  v1.2 iptables proxy是default的。

### 用户态的kube-proxy工作过程
 kube-proxy watches the Kubernetes master for the addition and removal of Service and Endpoints objects. For each Service it opens a port (randomly chosen) on the local node. Any connections to this “proxy port” will be proxied to one of the Service’s backend Pods (as reported in Endpoints). Which backend Pod to use is decided based on the SessionAffinity of the Service. Lastly, it installs iptables rules which capture traffic to the Service’s clusterIP (which is virtual) and Port and redirects that traffic to the proxy port which proxies the backend Pod.
 翻译过程：
 1. kube-proxy监控k8s master上service和endpoints的增删，对每个service，它会在本地node上随机选择并绑定一个端口
 2. 任何连接到此端口的连接都会被代理到service后端的任何一个Pod
 3. 为上述过程添加iptables 规则，捕获发给Service 虚拟IP和port的流量，重定向到proxy port，然后发给真正的后端服务Pod
 By default, the choice of backend is round robin. Client-IP based session affinity can be selected by setting service.spec.sessionAffinity to "ClientIP" (the default is "None").

### iptables kube-proxy工作过程
kube-proxy watches the Kubernetes master for the addition and removal of Service and Endpoints objects. For each Service it installs iptables rules which capture traffic to the Service’s clusterIP (which is virtual) and Port and redirects that traffic to one of the Service’s backend sets. For each Endpoints object it installs iptables rules which select a backend Pod.

## 为什么不使用轮询式DNS？
一个经常不时冒出来的问题是，为什么我选择用虚拟IP而不是使用标准的轮询式DNS做法。这有一些原因：

    * DNS库不遵守DNS TTLs并且对命名的查询进行缓存已经有很长的历史了
    * 很多应用只做一次DNS查询然后缓存查询的结果
    * 即使应用和库做了正确的重新解析，每一个客户端反反复复的DNS重新解析会难以管理