package cluster

import (
	"k8s.io/client-go/informers"
	"tf-cluster/internal/kube"
	"tf-cluster/library/log"
	"time"
)

var (
	//clientSet    *kubernetes.Clientset
	InformerFac  informers.SharedInformerFactory
	reSyncPeriod = 60 * time.Minute // informer watch period(全量同步时间)
)

func GetInformerFac(kubeConfig string, kubeId int) (informers.SharedInformerFactory, error) {
	var err error
	//kubeJson, err := yaml.YAMLToJSON([]byte(kubeConfig.Config))
	//if err != nil {
	//	log.Logger.Errorf("k8s bootstrap fail:%v", err)
	//	return
	//}
	clientSet, err := kube.GetKubeClient(kubeConfig, kubeId)
	log.Logger.Info("k8s init client success")
	if err != nil {
		log.Logger.Fatalf("k8s init client Error:%s", err.Error())
		return nil, err
	}
	return informers.NewSharedInformerFactory(clientSet, reSyncPeriod), err
}
