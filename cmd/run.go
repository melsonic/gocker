package cmd

import (
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
		args = append([]string{"/proc/self/exe", "child"}, args...)
		command := exec.Cmd{
			Path:   args[0],
			Args:   args,
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			SysProcAttr: &syscall.SysProcAttr{
				Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
				UidMappings: []syscall.SysProcIDMap{
					{
						ContainerID: 0,
						HostID:      1000,
						Size:        1,
					},
				},
				GidMappings: []syscall.SysProcIDMap{
					{
						ContainerID: 0,
						HostID:      1000,
						Size:        1,
					},
				},
			},
		}
		must(command.Run())
	},
}

func init() {
	rootCommand.AddCommand(runCommand)
}
