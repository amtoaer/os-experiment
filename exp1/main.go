package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var memory int
	fmt.Println("请输入模型总内存（单位MiB）")
	fmt.Scan(&memory)
	manager := NewManager(memory)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clear()
		fmt.Println(manager)
		fmt.Println("请输入采用的操作：")
		if scanner.Scan() {
			switch scanner.Text() {
			case "Create":
				manager.Create()
			case "EventWait":
				manager.EventWait()
			case "EventOccur":
				manager.EventOccur()
			case "Timeout":
				manager.Timeout()
			case "Release":
				manager.Release()
			}
		}
	}
}
