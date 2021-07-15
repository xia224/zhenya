# Kubernetes Service

* Service是一种抽象，它是一组Pods的逻辑集合和一个用于访问它们的策略。（大网的频道概念很像）
* 一个service的目标Pod集合通常是由Label Selector来决定的。

* 对于那些Kubernetes原生的应用，Kubernetes提供了一个简单的Endpoints API，会在Service中的Pod集合发生改变的时候更新。对于非Kubernetes原生的应用，Kubernetes为Service提供了一种基于虚拟IP的桥接方式使其重定向到后端的Pods。

* Service 是一个REST对象，可通过向apiServer发送增删改查service的命令。

## Service的一个简单实例定义
```json
{
    "kind": "Service",
    "apiVersion": "v1",
    "metadata": {
        "name": "my-service"
    },
    "spec": {
        "selector": {
            "app": "MyApp"
        },
        "ports": [
            {
                "protocol": "TCP",
                "port": 80,
                "targetPort": 9376
            }
        ]
    }
}
```
* 这个定义会创建一个新的Service对象，名字为”my-service”，它指向所有带有”app=MyApp”标签的Pod上面的9376端口。这个Service同时也会被分配一个IP地址（有时被称作”cluster ip”），它会被服务的代理所使用（见下面）。这个Service的选择器，会不断的对Pod进行筛选，并将结果POST到名字同样为“my-service”的Endpoints对象。

* 注意一个Service能将一个来源的端口映射到任意的targetPort。默认情况下，targetPort会被设置成与port字段一样的值。可能更有意思的地方在于，targetPort可以是一个字符串，能引用一个后端Pod中定义的端口名。实际指派给该名称的端口号在每一个Pod中可能会不同。这为部署和更新你的Service提供了很大的灵活性。例如，你可以在你的后端的下一个版本中更改开放的端口，而无需导致客户出现故障。

## 没有选择器的service
【这里可以做些东西，把集群内外连接起来】
Service一般是用来对Kubernetes　Pod的访问进行抽象，但是也可以用来对其他类型的后端进行抽象，如：
    * 你想在生产环境中有一个外部的数据库集群，但是在测试环境中你想使用自己的数据库
    * 你想把你的service指向位于另外一个Namespace或者集群下的service
    * 你正在把你的workload迁移到Kubernetes，并且一些backend后端运行在Kubernetes之外

如果你碰到这些场景之一，那么你可以定义一个没有选择器的Service。

因为没有选择器，相应的Endpoints对象不会被创建。你可以手动将service映射到自己特定的endpoint。
```json
{
    "kind": "Endpoints",
    "apiVersion": "v1",
    "metadata": {
        "name": "my-service"
    },
    "subsets": [
        {
            "addresses": [
                { "ip": "1.2.3.4" }
            ],
            "ports": [
                { "port": 9376 }
            ]
        }
    ]
}
```
注意：Endpoint的IP不能是本地回环接口 (127.0.0.0/8)，本地链路地址(169.254.0.0/16)或者本地链路多播地址((224.0.0.0/24)。

不用选择器来访问一个Service的表现就像它有一个选择器一样。流量会被路由到用户定义的endpoint（在这个例子是：1.2.3.4:9376）。

## 指定service的集群ip， 服务port
### 多port Services
spec.ports [
    {}
]
### 指定集群ip
设置spec.clusterIP字段，用户选择的IP地址必须是一个可用的IP地址，并且必须在API Server的service-cluster-ip-range启动参数所指定的CIDR范围内。

## Headless services
有时你不需要一个单独的服务IP地址，也不需要做负载均衡。在这种情况下，你可以创建一个”headless”的Service，只需要把集群IP(spec.clusterIP)指定为"None"即可。
这个选项让开发者可以减少对Kubernetes系统的耦合度，在他们想要的时候，能让他们自由决定如何用自己的方式去发现这些服务。应用仍然使用一个自注册的模式，并且其它服务发现的系统的适配器可以轻易的基于这个API被构建出来.

## 发布 services - service的类型
Kubernetes的ServiceTypes能让你指定你想要哪一种服务。默认的和基础的是ClusterIP，这会开放一个服务可以在集群内部进行连接。NodePort 和LoadBalancer是两种会将服务开放给外部网络的类型。

ServiceType字段的合法值是：

ClusterIP: 仅仅使用一个集群内部的IP地址 - 这是默认值，在上面已经讨论过。选择这个值意味着你只想这个服务在集群内部才可以被访问到。
NodePort: 在集群内部IP的基础上，在集群的每一个节点的端口上开放这个服务。你可以在任意<NodeIP>:NodePort地址上访问到这个服务。
LoadBalancer: 在使用一个集群内部IP地址和在NodePort上开放一个服务之外，向云提供商申请一个负载均衡器，会让流量转发到这个在每个节点上以<NodeIP>:NodePort的形式开放的服务上。
在使用一个集群内部IP地址和在NodePort上开放一个Service的基础上，还可以向云提供者申请一个负载均衡器，将流量转发到已经以NodePort形式开发的Service上。

注意尽管NodePort可以是TCP或者UDP的，对于Kubernetes 1.0来说，LoadBalancer还支持TCP。

## 集群外部访问集群-ingress对象
ingress-controller:
    将新加入的Ingress转化成Nginx的配置文件并使之生效
ingress:
    将Nginx的配置抽象成一个Ingress对象，每添加一个新的服务只需写一个新的Ingress的yaml文件即可

### ingress工作原理
* ingress controller通过和kubernetes api交互，动态的去感知集群中ingress规则变化，然后读取它，按照自定义的规则，规则就是写明了哪个域名对应哪个service，生成一段nginx配置，再写到nginx-ingress-control的pod里，这个Ingress controller的pod里运行着一个Nginx服务，控制器会把生成的nginx配置写入/etc/nginx.conf文件中，然后reload一下使配置生效。以此达到域名分配置和动态更新的问题。

* 第一步 安装ingress-nginx controller，官方下载对应的yaml文件
  tips：官方下载的yaml文件中容器image国内网络无法访问，我使用的办法是把image先下载下来（如何下载，自谋出路），然后推到私有仓库，kubelet访问私有仓库，需要创建对应的secret，方法参考：secret.md。
* 第二步 配置http代理
  具体就是：先用deployment起http 代理（nginx）的pod， 然后为pod配置servcie，最后启动Ingress，把相应的host/path路由到service。
  创建成功Ingress后， ingress controller会创建对应的server{}
* 第三步 配置/etc/hosts host到nodeip的地址映射
* 第四步 curl host:nodeport
* 待续 ssl + 规则重写