package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)


func clearScreen() error{
	switch runtime.GOOS {
	case "linux","darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err:= cmd.Run()
		if err != nil {
			fmt.Println("Failed to clear the screen;",err)
		}
	default:
		fmt.Println("Unsupported platform")
	}
	return nil
}
