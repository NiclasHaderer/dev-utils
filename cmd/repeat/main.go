package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func startProcess(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	args := os.Args
	if len(args) <= 2 {
		log.Fatalf("Usage: repeat <repeat_count> <command> [args...]\n")
	}

	repeatCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("The first argument must be an integer, got: %s\n", args[1])
	}

	// Start a process executing the command
	for i := 0; i < repeatCount; i++ {
		cmd := args[2:]
		err = startProcess(cmd)
		if err != nil {
			log.Fatalf("Error executing command: %v\n", err)
		}
	}
}
