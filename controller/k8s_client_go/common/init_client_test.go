package common

import (
	"testing"
)

func TestInitClient(t *testing.T) {
	kubeconfigfile := "/Users/zzy/.kube/config"
	if _, err := InitClient(&kubeconfigfile); err != nil {
		t.Errorf("Cannot parse kubeconfig path")
	}
}
