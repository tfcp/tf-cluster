package cluster

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"tf-cluster/internal/kube"
	"tf-cluster/internal/model/cluster"
	"tf-cluster/library/log"
	"time"
)

var (
	clientSet    *kubernetes.Clientset
	InformerFac  informers.SharedInformerFactory
	reSyncPeriod = 60 * time.Minute // informer watch period(全量同步时间)
)

func bootstrap() {
	var err error
	kubeConfigModel := new(cluster.Config)
	kubeConfig, _ := kubeConfigModel.OneConfig(map[string]interface{}{"id": 1})
	clientSet, err = kube.GetKubeClient(kubeConfig.Config, kubeConfig.Server)
	log.Logger.Info("k8s init client success")
	if err != nil {
		log.Logger.Fatalf("k8s init client Error:%s", err.Error())
		return
	}
	InformerFac = informers.NewSharedInformerFactory(clientSet, reSyncPeriod)
}
