# 介绍Nginx

## nginx的模块类型
* Event module: 事件处理模块
* Phase handler: 处理客户端请求并响应
* Output filter: 负责对输出的内容进行处理，对输出进行修改
* Upstream: 实现反向代理的功能，将请求转发达到后端服务器，并从后端服务器上读取响应，返回给客户端
* Load-balancer: 负载均衡，选择后端服务器

## nginx的网络IO模型
IO多路复用，epoll