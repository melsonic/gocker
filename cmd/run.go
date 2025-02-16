package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "run a command",
	Long:  "run a command with arguments in gocker!",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Cmd{
			Path:   "/bin/bash",
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			SysProcAttr: &syscall.SysProcAttr{
				Cloneflags: syscall.CLONE_NEWUTS,
			},
		}
		err := command.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func Execute() {
	if err := runCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
