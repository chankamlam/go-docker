package main
import(
	"github.com/spf13/cobra"
	"docker/container"
	"os"
)
func InitRunCmd() *cobra.Command{
	var runCmd = &cobra.Command{
		Use: "run",
		Short: "Run a command in a new container",
		Run: func(self *cobra.Command, args []string){
			is_tty,err := self.Flags().GetBool("tty")
			if err != nil{
				panic(err)
			}
			is_interactive,_ := self.Flags().GetBool("interactive")
			if err != nil{
				panic(err)
			}
			//fmt.Println(is_tty)
			//fmt.Println(is_interactive)
			cmd := container.CreateParentProcess(is_interactive,is_tty,args)
			if err := cmd.Start(); err != nil{
				panic(err)
			}
			cmd.Wait()
			os.Exit(-1)
		},
	}
	runCmd.Flags().BoolP("interactive","i",false,"Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty","t",false,"Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach","d",false,"Run container in background and print container ID")
	return runCmd
}
