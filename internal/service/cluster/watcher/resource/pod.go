package resource

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	cluster2 "tf-cluster/internal/model/cluster"
	cluster "tf-cluster/internal/service/cluster/watcher"
	"tf-cluster/library/log"
)

type PodWatcher struct {
}

func (this *PodWatcher) Run(config *cluster2.Config) {
	informerFac, err := cluster.GetInformerFac(config.Config, config.ID)
	if err != nil {
		log.Logger.Errorf("PodInformerService Run GetInformerFacError: %v", err)
		return
	}
	podInformer := informerFac.Core().V1().Pods()
	informer := podInformer.Informer()
	defer runtime.HandleCrash()
	stopper := make(chan struct{})
	defer close(stopper)
	// 启动 informer，list & watch
	go informerFac.Start(stopper)

	// 等待 Cache 都同步完毕,必不可少
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	log.Logger.Info("pod informer start")
	// 使用自定义 handler
	informer.AddEventHandler(
		this,
	)
	// 创建 lister
	podLister := podInformer.Lister()

	// 从 lister 中获取所有 items
	_, err = podLister.List(labels.Everything())
	if err != nil {
		log.Logger.Errorf("podInformerService Run Lister ListError: %v", err)
	}
	<-stopper
}

func (this *PodWatcher) OnDelete(obj interface{}) {

}

func (this *PodWatcher) OnUpdate(oldObj, newObj interface{}) {

}

func (this *PodWatcher) OnAdd(obj interface{}) {

}
