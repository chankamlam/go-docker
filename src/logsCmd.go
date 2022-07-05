package main
import(
	"fmt"
	"github.com/spf13/cobra"
)
func InitLogsCmd() *cobra.Command{
	var logsCmd = &cobra.Command{
		Use: "logs",
		Short: "Fetch the logs of a container",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println(args)
		},
	}
	return logsCmd
}
