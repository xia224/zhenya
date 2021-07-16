# k8s helper doc

## 书籍
《深入剖析kubernetes》张磊

## kubectl related
// load kubectl binary file
`curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"`

// load kubectl checksum file
`curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"`

// check
```echo "$(<kubectl.sha256) kubectl" | sha256sum --check
```
`kubectl version --client`

## Introduce minikube
minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.
`minikube start`
[minikube doc url](https://minikube.sigs.k8s.io/docs/start/)

`minikube kubectl -- get pods -A`

`minikube start --nodes 2 -p multinode-demo` // 启动多个nodes

## Helm related
`brew install helm` // install helm on macos
[helm doc url](https://helm.sh/zh/docs/intro/install/)

// add new chart repo
`helm repo add apache https://pulsar.apache.org/charts`

// check chart repo
`helm search repo apache`

## 资源概念
k8s系统中，所有内容都抽象为资源，资源实例化之后就叫做对象。
Kubernetes 对象是持久化的实体，Kubernetes 使用这些实体去表示整个集群的状态。
当创建 Kubernetes 对象时，必须提供对象的规约，用来描述该对象的期望状态，以及关于对象的一些基本信息。

## 获取集群节点的IP地址
kubectl --kubeconfig [USER_CLUSTER_KUBECONFIG] get nodes --output wide

## 获得用户集群的ssh密钥
kubectl --kubeconfig [ADMIN_CLUSTER_KUBECONFIG] get secrets -n [USER_CLUSTER_NAME] ssh-keys \
-o jsonpath='{.data.ssh\.key}' | base64 -d > \
~/.ssh/[USER_CLUSTER_NAME].key && chmod 600 ~/.ssh/[USER_CLUSTER_NAME].key

## minikube集群ssh登录
ssh -i ~/.minikube/machines/minikube/id_rsa docker@$(minikube ip)

## dig服务地址发现
dig -t A my-service-1.default.svc.cluster.local.