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
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	content,err := ioutil.ReadAll(logFile)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(os.Stdout,string(content))

}
