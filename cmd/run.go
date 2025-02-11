package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "run a command",
	Long:  "run a command with arguments in gocker!",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command(args[1], args[2:]...)
		output, err := command.Output()
		if err != nil {
			exitError := err.(*exec.ExitError)
			// print the actual error message returned by the command
			fmt.Println(string(exitError.Stderr))
			// return the process exit code
			os.Exit(err.(*exec.ExitError).ProcessState.ExitCode())
		}
		fmt.Println(string(output))
	},
}

func Execute() {
	if err := runCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
