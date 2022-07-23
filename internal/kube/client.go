package kube

import (
	"encoding/json"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
)

func GetKubeClient(kubeConfig, master string) (*kubernetes.Clientset, error) {
	//config, err := clientcmd.BuildConfigFromFlags("", g.Config().GetString("kube.configPath"))
	configV1 := clientcmdapi.Config{}
	err := json.Unmarshal([]byte(kubeConfig), &configV1)
	if err != nil {
		fmt.Println(err)
	}
	configObject, _ := clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	configInternal := configObject.(*clientcmdapi.Config)

	clientConfig, _ := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmdapi.Cluster{Server: master},
	}).ClientConfig()
	//clientConfig.QPS = defaultQPS
	//clientConfig.Burst = defaultBurst
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(clientConfig)
}
