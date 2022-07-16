package main
import(
	"github.com/spf13/cobra"
	"docker/container"
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
			fmt.Println(args)
			is_tty,err := self.Flags().GetBool("tty")
			if err != nil{
				panic(err)
			}
			is_interactive,err := self.Flags().GetBool("interactive")
			if err != nil{
				panic(err)
			}
			is_detach,err := self.Flags().GetBool("detach")
			if err != nil{
				panic(err)
			}
			if is_detach && is_tty {
				fmt.Errorf("Can not use -it and -d in the same time.")
			}
			containerName,err := self.Flags().GetString("name")
			if err != nil {
				panic(err)
			}
			if containerName == "" {
				containerName = GenerateContainerId(container.MAX_CONTAINER_ID)
			}
			cmd := container.CreateParentProcess(containerName,is_interactive,is_tty,args)
			if err := cmd.Start(); err != nil{
				panic(err)
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
