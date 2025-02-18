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
		if len(args) < 2 {
			os.Exit(2)
		}
		command := exec.Cmd{
			Path:   args[1],
			Args:   args[1:],
			Dir:    "/",
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			SysProcAttr: &syscall.SysProcAttr{
				Chroot:     "./alpine_root_fs/",
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
