package alert
import (
	"fmt"
)

func Show(err error,num string){
	switch num{
		case "001": Display("Can not get is_tty value.")
	    case "002": Display("Can not get is_interactive value.")
	    case "003": Display("Can not get is_detach value.")
	    case "004": Display("Can not use '-it' and '-d' at the same time.")
	    case "005": Display("Can not get container name.")
	    case "006": Display("Can not start the cmd.")
	    case "007": Display("Failed on syscall.Sethostname.")
	    case "008": Display("Failed on syscall.Chroot.")
	    case "009": Display("Failed on syscall.Chdir.")
	    case "010": Display("Failed on syscall.Mount 'proc' folder.")
	    case "011": Display("Failed on exec.LookPath.")
	    case "012": Display("Failed on syscall.Exec cmd.")
	    case "013": Display("Failed on copy base image.")
	    case "014": Display("Can not create container.log file.")
	    case "015": Display("Can not open container.log file.")
	    case "016": Display("Can not write output to container.log file.")
	}
	panic(err)
}

func Display(msg string){
	fmt.Println("[ERROR] %s",msg)
}

func Println(msg string){
	fmt.Println(msg)
}
