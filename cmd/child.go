package cmd

import (
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

		must(syscall.Chroot("/"))
		must(syscall.Chdir("/"))
		defer syscall.Unmount("proc", syscall.MNT_DETACH)
		must(syscall.Mount("proc", "proc", "proc", 0, ""))
		must(syscall.Unshare(syscall.CLONE_NEWNS))
		must(command.Run())
		return
	},
}

func init() {
	rootCommand.AddCommand(childCommand)
}
