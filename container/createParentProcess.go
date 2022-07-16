package container
import(
	"os"
	"os/exec"
	"syscall"
	"strings"
	"io/ioutil"
	"path/filepath"
)

func CreateParentProcess(containerName string,interactive bool,tty bool,args []string) *exec.Cmd {
	args = append([]string{containerName},args[0:]...)
	cmd := exec.Command("/proc/self/exe","child",strings.Join(args," "))
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
	imageFolderPath := IMAGE_FOLDER_PATH
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX,containerName,ROOTFS_NAME)
	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err){
		if err := CopyFileOrDirectory(imageFolderPath,rootFolderPath); err != nil{
			return err
		}
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
func CopyFileOrDirectory(src string, dst string) error{
	info,err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir(){
		if err := os.MkdirAll(dst,0777); err != nil{
			return err
		}
		if list,err := ioutil.ReadDir(src); err == nil {
			for _,item := range list{
				if err = CopyFileOrDirectory(filepath.Join(src,item.Name()),filepath.Join(dst,item.Name())); err != nil{
					return err
				}
			}
		}else{
			return err
		}
	}else{
		content,err := ioutil.ReadFile(src) 
		if err != nil{
			return err
		}
		if err := ioutil.WriteFile(dst,content,0777); err != nil{
			return err
		}
	}
	return nil
}
