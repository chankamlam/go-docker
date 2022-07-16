package container
import(
	"os"
	"os/exec"
	"syscall"
	"time"
	"math/rand"
	"io/ioutil"
	"path/filepath"
)
// const(
// 	MAX_CONTAINER_ID = 32
// 	IMAGE_FOLDER_PATH = "/var/lib/docker/images/base"
// 	ROOT_FOLDER_PATH_PREFEX = "/var/lib/docker/containers/"
// )
func CreateChildProcess(args []string) error{
	containerId := GenerateContainerId(MAX_CONTAINER_ID) 
	imageFolderPath := IMAGE_FOLDER_PATH
	rootFolderPath := ROOT_FOLDER_PATH_PREFEX + containerId
	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err){
		if err := CopyFileOrDirectory(imageFolderPath,rootFolderPath); err != nil{
			return err
		}
	}
	if err := syscall.Sethostname([]byte(containerId)); err != nil{
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
	path,err := exec.LookPath(args[0])
	if err != nil{
		return err
	}
	if err := syscall.Exec(path, args[0:], os.Environ()); err != nil {
		return err
	}
	return nil
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
func GenerateContainerId(n uint) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	length := len(letters)
        for i := range b {
            b[i] = letters[rand.Intn(length)]
        }
        return string(b)
}
