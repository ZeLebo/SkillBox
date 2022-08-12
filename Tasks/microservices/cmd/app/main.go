package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// TODO the output of all the programs to be logged
	fmt.Println("Setting up the environment...")
	fmt.Println("Starting the first server...")

	cmd1 := exec.Command("go", "run", "../proxy/first.go")
	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	go cmd1.Start()
	time.Sleep(time.Second * 2)

	fmt.Println("Starting the second server...")
	cmd2 := exec.Command("go", "run", "../proxy/second.go")
	stdout2, err := cmd2.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	go cmd2.Start()
	time.Sleep(time.Second * 2)
	fmt.Println("Starting the proxy server...")

	cmd3 := exec.Command("go", "run", "../proxy/proxy.go")
	stdout3, err := cmd3.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	go cmd3.Start()
	time.Sleep(time.Second * 2)

	for {
		tmp := make([]byte, 1024)
		_, err := stdout1.Read(tmp)
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
		_, err = stdout2.Read(tmp)
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
		_, err = stdout3.Read(tmp)
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
	}
}
