package process

import (
	"tf-cluster/app/process/watcher"
	"tf-cluster/library/utils"
)

var (
	processStdoutTitle    = []string{"进程名称", "备注", "负责人", "创建时间"}
	processStdoutContents = [][]string{}
)

// consumer process
func Process() {
	// process进程
	clusterStdout := watcher.ProcessCluster()
	processStdoutContents = append(processStdoutContents, clusterStdout...)

	//demoStdout := auth.ProcessDemo()
	//processStdoutContents = append(processStdoutContents, demoStdout...)

	// stdout cron list
	utils.TableStdout(processStdoutTitle, processStdoutContents)
}
