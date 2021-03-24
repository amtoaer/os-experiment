package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getError(err string) {
	fmt.Printf("非法调用！%s，点击回车继续", err)
	fmt.Scanln()
}
