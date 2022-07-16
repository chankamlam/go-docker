package container

import(
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func DisplayContainerLog(containerName string){
	logFilePath := filepath.Join(ROOT_FOLDER_PATH_PREFEX,containerName,LOG_FILE_NAME)
	logFile,err := os.Open(logFilePath)
	defer logFile.close()
	if err != nil {
		fmt.Errorf("Can not find the container.log file.")
		return 
	}
	content,err := ioutil.ReadAll(logFile)
	if err != nil {
		fmt.Errorf("can not read the container.log file.")
		return
	}
	fmt.Fprint(os.Stout,string(content))

}
