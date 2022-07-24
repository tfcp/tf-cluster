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

type StatefulWatcher struct {
}

func (this *StatefulWatcher) Run(config *cluster2.Config) {
	informerFac, err := cluster.GetInformerFac(config.Config, config.ID)
	if err != nil {
		log.Logger.Errorf("StatefulInformerService Run GetInformerFacError: %v", err)
		return
	}
	statefulInformer := informerFac.Apps().V1().StatefulSets()
	informer := statefulInformer.Informer()
	defer runtime.HandleCrash()
	stopper := make(chan struct{})
	defer close(stopper)
	// 启动 informer，list & watch
	go cluster.InformerFac.Start(stopper)

	// 等待 Cache 都同步完毕,必不可少
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	log.Logger.Info("stateful informer start")
	// 使用自定义 handler
	informer.AddEventHandler(
		this,
	)
	// 创建 lister
	statefulLister := statefulInformer.Lister()

	// 从 lister 中获取所有 items
	_, err = statefulLister.List(labels.Everything())
	if err != nil {
		log.Logger.Errorf("StatefulInformerService Run Lister ListError: %v", err)
	}
	<-stopper
}

func (this *StatefulWatcher) OnDelete(obj interface{}) {

}

func (this *StatefulWatcher) OnUpdate(oldObj, newObj interface{}) {

}

func (this *StatefulWatcher) OnAdd(obj interface{}) {

}
