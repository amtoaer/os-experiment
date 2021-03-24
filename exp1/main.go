package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	manager := NewManager(4000)
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
			case "Release":
				manager.Release()
			}
		}
	}
}
