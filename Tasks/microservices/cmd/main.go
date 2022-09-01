package main

import (
	"fmt"
	"os/exec"
)

// entry point
// start the whole instances
func main() {
	go func() {
		err := exec.Command("go", "run", "app/first.go").Run()
		if err != nil {
			fmt.Println("Cannot start the app", err.Error())
		}
	}()
	go func() {
		err := exec.Command("go", "run", "app/second.go").Run()
		if err != nil {
			fmt.Println("Cannot start the app", err.Error())
		}
	}()

	err := exec.Command("go", "run", "app/proxy.go").Run()
	if err != nil {
		fmt.Println("Cannot start the app", err.Error())
	}
}
