package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var childCommand = &cobra.Command{
	Use:   "child",
	Short: "run child",
	Long:  "run child command inside container!",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Cmd{
			Path:   args[0],
			Args:   args,
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		err := syscall.Chroot("/")
		if err != nil {
			fmt.Println(err.Error())
		}
		err = syscall.Chdir("/")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer syscall.Unmount("proc", syscall.MNT_DETACH)
		err = syscall.Mount("proc", "proc", "proc", 0, "")
		if err != nil {
			fmt.Println("mount error: ", err.Error())
			return
		}
		err = syscall.Unshare(syscall.CLONE_NEWNS)
		if err != nil {
			fmt.Println("unshare error: ", err.Error())
			return
		}
		err = command.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCommand.AddCommand(childCommand)
}
