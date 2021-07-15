---
#ListWatch all pods added and deleted
`go run . -kubeconfig "/Users/zzy/.kube/config" -master "https://kubernetes.docker.internal:6443"`

#Fix "module k8s.io/client-go@latest found (v1.5.2), but does not contain package k8s.io/client-go/kubernetes"
`k8s.io/client-go v0.19.2`