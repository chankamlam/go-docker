package main
import(
	"fmt"
	"github.com/spf13/cobra"
)
func InitRunCmd() *cobra.Command{
	var runCmd = &cobra.Command{
		Use: "run",
		Short: "Run a command in a new container",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println(args)
		},
	}
	runCmd.Flags().BoolP("interactive","i",false,"Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty","t",false,"Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach","d",false,"Run container in background and print container ID")
	return runCmd
}
