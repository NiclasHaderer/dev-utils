package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

func startProcess(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getCommandToExecute() []string {
	// Get the index of the "--" separator
	separatorIndex := 0
	for i, arg := range os.Args {
		if arg == "--" {
			separatorIndex = i
			break
		}
	}

	if separatorIndex == 0 {
		log.Fatal("Could not find delimiter '--' between flags and command")
	}

	command := os.Args[separatorIndex+1:]
	if len(command) == 0 {
		log.Fatal("No command provided")
	}

	return command
}

var rootCmd = &cobra.Command{
	Short: "Repeat a command a number of times",
	Use:   `repeat [flags] -- [command]`,
	Run: func(cmd *cobra.Command, args []string) {
		times, _ := cmd.Flags().GetInt("times")
		shouldClear, _ := cmd.Flags().GetBool("clear")
		verbose, _ := cmd.Flags().GetBool("verbose")

		command := getCommandToExecute()

		for i := 0; i < times; i++ {
			if shouldClear && i > 0 {
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				_ = cmd.Run()
			}

			if verbose && (i == 0 || shouldClear) {
				fmt.Print("Repeating: \"", strings.Join(command, " "), "\" for ", times, " times (idx=", i+1, "). ")
				if shouldClear {
					fmt.Print("Clearing screen after each command.")
				}
				println("")

			}

			if err := startProcess(command); err != nil {
				log.Fatal(err)
			}

		}
	},
}

func main() {

	rootCmd.Flags().IntP("times", "t", 2, "Number of times to repeat the command")
	rootCmd.Flags().BoolP("clear", "c", false, "Clear the screen before each command")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print the command before executing it")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
