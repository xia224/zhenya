# Ways to access k8s

## 1. Using kubectl proxy
kubectl proxy --port=8080 &
curl http://localhost:8080/api/
curl http://localhost:10081/proxyMode

## 2. Get authentication token and access API server directly
### Check all possible clusters, as your .KUBECONFIG may have multiple contexts:
kubectl config view -o jsonpath='{"Cluster name\tServer\n"}{range .clusters[*]}{.name}{"\t"}{.cluster.server}{"\n"}{end}'

### Select name of cluster you want to interact with from above output:
export CLUSTER_NAME="some_server_name"

### Point to the API server referring the cluster name
APISERVER=$(kubectl config view -o jsonpath="{.clusters[?(@.name==\"$CLUSTER_NAME\")].cluster.server}")

### Gets the token value
TOKEN=$(kubectl get secrets -o jsonpath="{.items[?(@.metadata.annotations['kubernetes\.io/service-account\.name']=='default')].data.token}"|base64 --decode)

### Explore the API with TOKEN
curl -X GET $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure

curl -X GET https://192.168.49.2:8443/api/v1/namespaces/kube-system/services/kube-dns --header "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6InNMS2dXYllKbFctWkdWajZuQXVvN0dpSDhfQUd4bDBhSnBQdng3VUJkcHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4tdzV0bG4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImE4ZWJmYmZkLTRiYjctNGE2Yi04ZTM0LWZjNWVhMjRkZGMxOSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.LC4SMnWjwTDR1w-xlG7Boaz_0LNNY-E_ULKog8x1xy5clruescLG9yLMKOdmMFhnBMxHXqGVFUnhjPldqyiuJfg3hTZV5LH5m2jLgRWGGpg7x6TvHQGPvK0lttC5H5kHLd7zu0qvc7xjAh-FmgH05uc5aUjQXUnFAJkkx_ZOSk8AF6MNgRBx5WHwXur4ezAjM8wZIN-w6W1btqBNvrYE4Z7G6dVz4SAALEX1zZ6XzFWNggGuvZWvmRRdHc0P8_QYlH63mBqo9jmKpffp3X00TJvR4Il9o4FHWTIm4ZQMSjfuAjLmB0OZHghVzRmCOW_ndfh3iJLiqqYwR2LhnySPqQ" --insecure

## 3. Go client

## 私有仓库的密码保存
$HOME/.docker/config.json
虽然在node上登录hub成功，但是集群里的pod仍然不能拉取到image，下面这种方式work：
第一步：创建secret
kubectl create secret docker-registry agoraregistry \
    --docker-server=hub.agoralab.co \
    --docker-username=xiazhenya@agora.io \
    --docker-password=Xia_2020_ioo \
    --docker-email=xiazhenya@agora.io

第二步：启动pod时指定imagePullSecrets
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: private-image-test-1
spec:
  containers:
    - name: uses-private-image
      image: hub.agoralab.co/uap/tools/k8s/echoserver:1.4
      imagePullPolicy: IfNotPresent
      command: [ "echo", "SUCCESS" ]
  imagePullSecrets:
    - name: agoraregistry
EOF


## 获得节点列表
*如果想要节点名称：nodes=$(kubectl get nodes -o jsonpath='{range.items[*].metadata}{.name} {end}')

*如果想要节点 IP ，nodes=$(kubectl get nodes -o jsonpath='{range .items[*].status.addresses[?(@.type=="ExternalIP")]}{.address} {end}')

## 暴露服务
kubectl expose deployment source-ip-app --name=clusterip --port=80 --target-port=8080