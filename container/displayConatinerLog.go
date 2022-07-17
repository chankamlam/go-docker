package container

import(
	"os"
	"io/ioutil"
	"path/filepath"
	"docker/alert"
)

func DisplayContainerLog(containerName string){
	logFilePath := filepath.Join(ROOT_FOLDER_PATH_PREFEX,containerName,LOG_FILE_NAME)
	logFile,err := os.Open(logFilePath)
	defer logFile.Close()
	if err != nil {
		alert.Show(err,"015")
	}
	content,err := ioutil.ReadAll(logFile)
	if err != nil {
		alert.Show(err,"016")
	}
	fmt.Fprint(os.Stdout,string(content))

}
