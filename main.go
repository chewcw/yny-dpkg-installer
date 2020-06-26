package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("usage: installer <fileName>")
			os.Exit(-1)
		}
	}()

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("usage: installer <fileName>")
		os.Exit(0)
	}

	fileName := os.Args[1]

	cmd := exec.Command("dpkg", "-i", fileName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error while trying to install, error: %s", err)
	}
}
