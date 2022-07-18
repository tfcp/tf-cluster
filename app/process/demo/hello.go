package demo

import (
	"fmt"
	"time"
)

func HelloProcess() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("this is auth-hello process")
	}
}

func TestProcess() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("this is auth-test process")
	}
}
