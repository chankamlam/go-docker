package main
import(
	"github.com/spf13/cobra"
	"docker/container"
	"docker/alert"
)
func InitChildCmd() *cobra.Command{
	var childCmd = &cobra.Command{
		Use: "child",
		Run: func(self *cobra.Command,args []string){
			err,num := container.CreateChildProcess(args)
			if err!=nil{
				alert.Show(err,num)
			}
		},
	}
	return childCmd
}

