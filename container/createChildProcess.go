package container
import(
	"os"
	"os/exec"
	"syscall"
	"path/filepath"
	"fmt"
)
func CreateChildProcess(args []string) error{
	containerName := args[0]
	fmt.Println(containerName)
	fmt.Println(args)
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX,containerName,ROOTFS_NAME)
	fmt.Println(rootFolderPath)
	if err := syscall.Sethostname([]byte(containerName)); err != nil{
		return err
	}
	if err := syscall.Chroot(rootFolderPath); err != nil{
		return err
	}
	if err := syscall.Chdir("/"); err != nil{
		return err
	}
	if err := syscall.Mount("proc","/proc","proc",0,""); err != nil{
		return err
	}
	path,err := exec.LookPath(args[1])
	if err != nil{
		return err
	}
	if err := syscall.Exec(path, args[1:], os.Environ()); err != nil {
		return err
	}
	return nil
}


