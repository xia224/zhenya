package main

import (
	"context"
	"controller/k8s_client_go/common"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = apps_v1.Deployment{}
		replicas   int32
		err        error
		mark       int32
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

	// Read YAML
	if deployYaml, err = ioutil.ReadFile("/Users/zzy/Documents/k8s/nginx_deployment.yaml"); err != nil {
		goto FATL
	}

	// Convert yaml to json
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		goto FATL
	}

	// Convert Json to struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		goto FATL
	}
	fmt.Println(deployment)
	// Modify replica number
	replicas = 1
	deployment.Spec.Replicas = &replicas

	// Check if delloyment existed
	mark = 1
	if _, err = clientset.AppsV1().Deployments("default").Get(context.TODO(), deployment.Name, meta_v1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			goto FATL
		}

		mark = 2
		if _, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), &deployment, meta_v1.CreateOptions{}); err != nil {

			goto FATL
		}
	} else {
		mark = 3
		if _, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), &deployment, meta_v1.UpdateOptions{}); err != nil {
			goto FATL
		}
	}

	fmt.Println("Deployment Operation successful")
	fmt.Println(mark)
	return
FATL:
	fmt.Println(err)
	fmt.Println(mark)
}
