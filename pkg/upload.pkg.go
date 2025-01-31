package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func HandleUpload(multiPartFile multipart.File, directoryName string, multiPartHandler *multipart.FileHeader) (fileName string, err error) {
	fileName = GenerateUUID()
	var extensionFile string = filepath.Ext(multiPartHandler.Filename)
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	targetFile, err := os.OpenFile(filepath.Join(dir, directoryName, fmt.Sprintf("%s%s", fileName, extensionFile)), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer targetFile.Close()
	fileName = fileName + extensionFile

	if _, err = io.Copy(targetFile, multiPartFile); err != nil {
		return
	}
	return
}
