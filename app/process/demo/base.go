package demo

import "time"

func ProcessDemo() [][]string {
	// demo
	go HelloProcess()
	go TestProcess()

	return [][]string{
		// watcher
		{"demo进程", "", "sujizhao", time.Now().Format("2006-01-02 15:04:05")},
	}
}
