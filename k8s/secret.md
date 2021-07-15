# 身份认证

## 拉取私有镜像，进行身份认证
登录私有镜像仓库，会创建或更新~/.docker/config.json文件维护授权token。

## 基于已经存在的docker凭据创建secret
```bash
kubectl create secret generic regcred \
     --from-file=.dockerconfigjson=<path/to/.docker/config.json> \
     --type=kubernetes.io/dockerconfigjson
```

## 用用户名和密码创建secret
```bash
kubectl create secret docker-registry myregcred \
     --docker-server=<your-registry-server> \
     --docker-username=<your-name> \
     --docker-password=<your-password> \
     --docker-email=<your-email> 
```

## 查看secret
kubectl get secret myregcred --output=yaml

## 用secret创建一个pod
```yaml
apiVersion: v1
 kind: Pod
 metadata:
   name: private-reg
 spec:
   containers:
   - name: private-reg-container
     image: hub.agoralab.co/uap/rtmp_pusher/rtmp_pusher-worker:release_20210630_1_0dd915bbb062
   imagePullSecrets:
   - name: myregcred
```