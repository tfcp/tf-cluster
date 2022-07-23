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
	log.Logger.Infof("watchCluster start...")
	clusters.Listen()
}
