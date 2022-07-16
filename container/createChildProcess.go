package container
import(
	"os"
	"os/exec"
	"syscall"
	"path/filepath"
	"fmt"
	"strings"
)
func CreateChildProcess(args []string) error{
	// args len equal one like [xx yy zz], so here need to split
	cmdArr := strings.Split(args[0]," ")
	containerName := cmdArr[0]
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
	path,err := exec.LookPath(cmdArr[1])
	if err != nil{
		return err
	}
	if err := syscall.Exec(path, cmdArr[1:], os.Environ()); err != nil {
		return err
	}
	return nil
}


