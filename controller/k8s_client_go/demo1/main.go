package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"controller/k8s_client_go/common"

	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		pods      *core_v1.PodList
		err       error
	)

	var kubeconfigfile string
	if home := common.GetHomeDir(); home != "" {
		flag.StringVar(&kubeconfigfile, "kubeconfig", filepath.Join(home, ".kube", "config"), "(Optional)Absolute path of k8s config file")
	} else {
		flag.StringVar(&kubeconfigfile, "kubeconfig", "", "(Optional)Absolute path of k8s config file")
	}

	if clientset, err = common.InitClient(&kubeconfigfile); err != nil {
		goto FATL
	}

	if pods, err = clientset.CoreV1().Pods("default").List(context.TODO(), meta_v1.ListOptions{}); err != nil {
		goto FATL
	}
	fmt.Println(*pods)

FATL:
	fmt.Println(err)
}
