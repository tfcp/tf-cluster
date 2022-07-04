package kube

import (
	"github.com/gogf/gf/frame/g"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", g.Config().GetString("kube.configPath"))
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
