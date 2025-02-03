package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"utils/lib/process"
)

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
	Use:   `repeat-cmd [flags] -- [command]`,
	Run: func(cmd *cobra.Command, args []string) {
		times, _ := cmd.Flags().GetInt("times")
		shouldClear, _ := cmd.Flags().GetBool("clear")
		verbose, _ := cmd.Flags().GetBool("verbose")
		ignoreFail, _ := cmd.Flags().GetBool("ignore-fail")

		command := getCommandToExecute()

		for i := 0; i < times; i++ {
			if shouldClear && (i > 0 || verbose) {
				_ = process.Run([]string{"clear"}, true)
			}

			if verbose && (i == 0 || shouldClear) {
				fmt.Print("Repeating: \"", strings.Join(command, " "), "\" for ", times, " times (idx=", i+1, "). ")
				if shouldClear {
					fmt.Print("Clearing screen after each command.")
				}
				println("")

			}

			if err := process.Run(command, true); err != nil && !ignoreFail {
				log.Print(err)
				os.Exit(err.Status())
			}

		}
	},
}

func main() {

	rootCmd.Flags().IntP("times", "t", 100, "Number of times to repeat the command")
	rootCmd.Flags().BoolP("clear", "c", true, "Clear the screen before each command")
	rootCmd.Flags().BoolP("verbose", "v", true, "Print the command before executing it")
	rootCmd.Flags().Bool("ignore-fail", false, "Do not fail if the command fails, but continue to the next iteration")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
