package container
import(
	"os"
	"os/exec"
	"syscall"
	"path/filepath"
	"fmt"
	"strings"
)
func CreateChildProcess(args []string) (error,string){
	// args len equal one like [xx yy zz], so here need to split
	cmdArr := strings.Split(args[0]," ")
	containerName := cmdArr[0]
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX,containerName,ROOTFS_NAME)

	if err := syscall.Sethostname([]byte(containerName)); err != nil{
		return err,"007"
	}
	if err := syscall.Chroot(rootFolderPath); err != nil{
		return err,"008"
	}
	if err := syscall.Chdir("/"); err != nil{
		return err,"009"
	}
	if err := syscall.Mount("proc","/proc","proc",0,""); err != nil{
		return err,"010"
	}
	path,err := exec.LookPath(cmdArr[1])
	if err != nil{
		return err,"011"
	}
	if err := syscall.Exec(path, cmdArr[1:], os.Environ()); err != nil {
		return err,"012"
	}
	return nil,nil
}


