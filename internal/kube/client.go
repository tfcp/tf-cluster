package kube

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient(kubeConfig string, kubeId int) (*kubernetes.Clientset, error) {
	var err error
	configPath := fmt.Sprintf("%s%d-config", g.Config().GetString("kube.configPath"), kubeId)
	if gfile.Exists(configPath) {
		if err = gfile.Remove(configPath); err != nil {
			return nil, err
		}
	}
	if err = gfile.PutContents(configPath, kubeConfig); err != nil {
		return nil, err
	}
	clientConfig, err := clientcmd.BuildConfigFromFlags("", configPath)

	// TODO 版本不兼容问题(json marshal异常)
	//configV1 := clientcmdapi.Config{}
	//err := json.Unmarshal(kubeJson, &configV1)
	//if err != nil {
	//	return nil, err
	//}
	//configObject, _ := clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	//configInternal := configObject.(*clientcmdapi.Config)
	//
	//clientConfig, _ := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
	//	ClusterDefaults: clientcmdapi.Cluster{Server: master},
	//}).ClientConfig()
	//clientConfig.QPS = defaultQPS
	//clientConfig.Burst = defaultBurst
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(clientConfig)
}
