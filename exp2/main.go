package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var size int
	fmt.Print("请输入缓冲区大小：")
	fmt.Scan(&size)
	bufModel := NewBuffer(size)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clear()
		fmt.Println(bufModel)
		fmt.Println("请输入采用的操作：")
		if scanner.Scan() {
			switch scanner.Text() {
			case "Produce":
				bufModel.Produce()
			case "Consume":
				bufModel.Consume()
			}
		}
	}
}
