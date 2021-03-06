package resource

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"tf-cluster/internal/enum"
	cluster2 "tf-cluster/internal/model/cluster"
	"tf-cluster/internal/service/cluster/watcher"
	"tf-cluster/library/log"
)

type NodeWatcher struct {
	// model Model
	nodeModel   cluster2.Node
	configModel *cluster2.Config
}

func (this *NodeWatcher) Run(config *cluster2.Config) {
	informerFac, err := cluster.GetInformerFac(config.Config, config.ID)
	if err != nil {
		log.Logger.Errorf("NodeInformerService Run GetInformerFacError: %v", err)
		return
	}
	this.configModel = config
	nodeInformer := informerFac.Core().V1().Nodes()
	informer := nodeInformer.Informer()
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
	log.Logger.Info("node informer start")
	// 使用自定义 handler
	informer.AddEventHandler(
		this,
	)
	// 创建 lister
	nodeLister := nodeInformer.Lister()

	// 从 lister 中获取所有 items
	_, err = nodeLister.List(labels.Everything())
	if err != nil {
		log.Logger.Errorf("NodeInformerService Run Lister ListError: %v", err)
	}
	<-stopper
}

func (this *NodeWatcher) OnDelete(obj interface{}) {
	if node, ok := obj.(*v1.Node); ok {
		fmt.Println("nodeDelete:", node)
	}
}

func (this *NodeWatcher) OnUpdate(oldObj, newObj interface{}) {
	//if Obj, ok := newObj.(*v1.Node); ok {
	//
	//}
}

func (this *NodeWatcher) OnAdd(obj interface{}) {
	if node, ok := obj.(*v1.Node); ok {
		unschedulable := enum.UnSchedulabledEnable
		if node.Spec.Unschedulable {
			unschedulable = enum.UnSchedulabledDisable
		}
		newNode := cluster2.Node{
			Ip:            node.Name,
			Labels:        gconv.String(node.Labels),
			Taints:        gconv.String(node.Spec.Taints),
			Unschedulable: unschedulable,
			Cluster:       this.configModel.Name,
			CreateAt:      node.CreationTimestamp.Time,
		}
		if err := this.nodeModel.CreateNode(newNode); err != nil {
			log.Logger.Errorf("NodeWatcher OnAdd CreateNodeError: %v", err)
		}
	}
}
