package container
import(
	"os"
	"os/exec"
	"syscall"
)

func CreateParentProcess(interactive bool,tty bool,args []string) *exec.Cmd {
	cmd:=exec.Command("/proc/self/exe","child",args[0])
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:syscall.CLONE_NEWUTS|syscall.CLONE_NEWPID|syscall.CLONE_NEWNS|syscall.CLONE_NEWIPC|syscall.CLONE_NEWNET|syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	if tty{
		if interactive{
			cmd.Stdin = os.Stdin
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}else{
		// detach mode
		
	}
	return cmd
}
