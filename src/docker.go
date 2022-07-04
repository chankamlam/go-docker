package main
import(
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)
func main(){
	var versionCmd = &cobra.Command{
		Use: "version",
		Short: "Show the Docker version information",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("$$$$$$$$")
		},
	}
	var psCmd = &cobra.Command{
		Use: "ps",
		Short: "List containers",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println(strings.Join(args," "))
		},
	}
	psCmd.Flags().BoolP("all","a",false,"Show all containers (default shows just running)")
	psCmd.Flags().BoolP("quiet","q",false,"Only display container IDs")
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
	var execCmd = &cobra.Command{
		Use: "exec",
		Short: "Run a command in a running container",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("&&&&&")
		},
	}
	var startCmd = &cobra.Command{
		Use: "start",
		Short: "Start one or more stopped containers",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("!!!!!!")
		},
	}
	var stopCmd = &cobra.Command{
		Use: "stop",
		Short: "Stop one or more running containers",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("~~~~~")
		},
	}
	var logsCmd = &cobra.Command{
		Use: "logs",
		Short: "Fetch the logs of a container",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("~~~~~")
		},
	}
	var rootCmd = &cobra.Command{
		Use: "docker [Command]",
	}
	rootCmd.AddCommand(psCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.Execute()
}
