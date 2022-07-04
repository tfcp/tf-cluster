package watcher

import "time"

func ProcessCluster() [][]string {
	go WatchCluster()
	return [][]string{
		// watcher
		{"集群信息同步监控", "备注", "sujizhao", time.Now().Format("2006-01-02 15:04:05")},
	}
}
