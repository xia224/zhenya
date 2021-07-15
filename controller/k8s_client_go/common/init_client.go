package common

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func InitClient(kubeconfigpath *string) (clientset *kubernetes.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigpath)
	if err != nil {
		return
	}

	if clientset, err = kubernetes.NewForConfig(config); err != nil {
		return
	}

	return
}

func GetHomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	return ""
}
