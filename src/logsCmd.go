package main
import(
	"github.com/spf13/cobra"
	"docker/container"
)
func InitLogsCmd() *cobra.Command{
	var logsCmd = &cobra.Command{
		Use: "logs",
		Short: "Fetch the logs of a container",
		Run: func(cmd *cobra.Command, args []string){
			container.DisplayContainerLog(args[0])
		},
	}
	return logsCmd
}
