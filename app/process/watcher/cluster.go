package watcher

import (
	cluster "tf-cluster/internal/service/cluster/watcher"
	"tf-cluster/internal/service/cluster/watcher/resource"
	"tf-cluster/library/log"
)

func WatchCluster() {
	clusters := cluster.NewClusters()
	clusters.Add(
		// 节点监听
		&resource.NodeWatcher{},
	//&resource.PodWatcher{},
	//&resource.DeployWatcher{},
	//&resource.ServiceWatcher{},
	)
	if err := clusters.Listen(); err != nil {
		log.Logger.Errorf("watchCluster error:%v", err)
		return
	}
	log.Logger.Infof("watchCluster start...")
}
