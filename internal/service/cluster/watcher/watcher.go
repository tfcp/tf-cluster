package cluster

import (
	"tf-cluster/internal/model/cluster"
)

type Watcher interface {
	OnDelete(obj interface{})
	OnUpdate(oldObj, newObj interface{})
	OnAdd(obj interface{})
	Run(config *cluster.Config)
}

type Clusters struct {
	watcherList []Watcher
}

func NewClusters() *Clusters {
	clusters := new(Clusters)
	return clusters
}

func (this *Clusters) Listen() error {
	var err error
	kubeConfigModel := new(cluster.Config)
	kubeConfigs, err := kubeConfigModel.ListConfig(nil, 1, 1000)
	if err != nil {
		return err
	}
	for _, kubeConfig := range kubeConfigs {
		for _, watcher := range this.watcherList {
			go watcher.Run(kubeConfig)
		}
	}
	return err
}

func (this *Clusters) Add(watch ...Watcher) {
	if len(watch) > 0 {
		this.watcherList = append(this.watcherList, watch...)
	}
}
