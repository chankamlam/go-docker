package main
import(
	"github.com/spf13/cobra"
	"docker/container"
	"docker/alert"
	"os"
	"fmt"
	"time"
	"math/rand"
)
func InitRunCmd() *cobra.Command{
	var runCmd = &cobra.Command{
		Use: "run",
		Short: "Run a command in a new container",
		Run: func(self *cobra.Command, args []string){
			is_tty,err := self.Flags().GetBool("tty")
			if err != nil{
				alert.Show(err,"001")
			}
			is_interactive,err := self.Flags().GetBool("interactive")
			if err != nil{
				alert.Show(err,"002")
			}
			is_detach,err := self.Flags().GetBool("detach")
			if err != nil{
				alert.Show(err,"003")
			}
			if is_detach && is_tty {
				alert.Show(nil,"004")
				return
			}
			containerName,err := self.Flags().GetString("name")
			if err != nil {
				alert.Show(err,"005")
			}
			if containerName == "" {
				containerName = GenerateContainerId(container.MAX_CONTAINER_ID)
			}
			cmd := container.CreateParentProcess(containerName,is_interactive,is_tty,args)
			if err := cmd.Start(); err != nil{
				alert.Show(err,"006")
			}
			if !is_detach {
				cmd.Wait()
			}
			os.Exit(-1)
		},
	}
	runCmd.Flags().BoolP("interactive","i",false,"Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty","t",false,"Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach","d",false,"Run container in background and print container ID")
	runCmd.Flags().StringP("name","n","","Assign a name to the container")
	return runCmd
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
