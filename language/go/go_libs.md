# Record go useful library

## Gomega
Gomega is the Ginkgo BDD-style testing framework's preferred matcher library.
Gomega on Github: http://github.com/onsi/gomega

buffer:
"github.com/onsi/gomega/gbytes"

## Ginkgo
Ginkgo 是一个 Go 测试框架，旨在帮助你有效地编写富有表现力的全方位测试。
Ginkgo on Github: "github.com/onsi/ginkgo"

## client-go 
Go clients for talking to a kubernetes cluster
"k8s.io/client-go"
Informer是Client-go中的一个核心工具包。在Kubernetes源码中，如果Kubernetes的某个组件，需要List/Get Kubernetes中的Object，在绝大多 数情况下，会直接使用Informer实例中的Lister()方法（该方法包含 了 Get 和 List 方法）.

## convert uint64 to string
不能直接用string()
strconv.FormatUint(number, base)

## Web server 框架
"github.com/gin-gonic/gin"

## log 日志
log "github.com/sirupsen/logrus"