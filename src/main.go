package main
import(
//	"fmt"
	"github.com/spf13/cobra"
//	"strings"
)
func main(){
	var rootCmd = &cobra.Command{
		Use: "docker [Command]",
	}
	rootCmd.AddCommand(InitRunCmd())
	rootCmd.AddCommand(InitChildCmd())
	rootCmd.AddCommand(InitLogsCmd())
	rootCmd.Execute()
}
