package main
import(
	"github.com/spf13/cobra"
	"docker/container"
	"fmt"
)
func InitChildCmd() *cobra.Command{
	var childCmd = &cobra.Command{
		Use: "child",
		Run: func(self *cobra.Command,args []string){
			fmt.Println(args)
			err := container.CreateChildProcess(args)
			if err!=nil{
				panic(err)
			}
		},
	}
	return childCmd
}

