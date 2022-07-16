package main
import(
	"github.com/spf13/cobra"
	"docker/container"
	"os"
	"fmt"
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
			cmd := container.CreateParentProcess(is_interactive,is_tty,args)
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
	return runCmd
}
