# Kube-apiserver

kube-apiserver 主要通过对外提供 API 的方式与其他组件进行交互
## 重要的API
core group：主要在 /api/v1 下；
named groups：其 path 为 /apis/$NAME/$VERSION；
暴露系统状态的一些 API：如/metrics 、/healthz 等；

API 的 URL 大致以 /apis/group/version/namespaces/my-ns/myresource 组成

## 处理请求流程
最外面接入点是http server，接受http api请求，可以用nginx/tomcat收http；

当请求到达 kube-apiserver 时，kube-apiserver 首先会执行在 http filter chain 中注册的过滤器链，
该过滤器对其执行一系列过滤操作，主要有认证、鉴权等检查操作。当 filter chain 处理完成后，
请求会通过 route 进入到对应的 handler 中，handler 中的操作主要是与 etcd 的交互。

## kube-apiserver的组件
### Aggregator
暴露的功能类似于一个七层负载均衡，将来自用户的请求拦截转发给其他服务器，并且负责整个 APIServer 的 Discovery 功能.
### KubeAPIServer
负责对请求的一些通用处理，认证、鉴权等，以及处理各个内建资源的 REST 服务.
为 kubernetes 中众多 API 注册路由信息，暴露 RESTful API 并且对外提供 kubernetes service.
### APIExtensionServer
主要处理 CustomResourceDefinition（CRD）和 CustomResource（CR）的 REST 请求, 也是 Delegation 的最后一环.