package components

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type TargetUpload interface {
	Target(*multipart.FileHeader, *[]byte) (filename string, basepath string, err error)
}

func Upload(r *http.Request, key string, target TargetUpload) (targetFilename string, targetPath string, err error) {
	r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile(key)
	if err != nil {
		return "", "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	filename, basepath, err := target.Target(fileHeader, &bytes)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	f, e := os.Stat(basepath)
	if !os.IsExist(e) || !f.IsDir() {
		os.MkdirAll(basepath, 0776)
	}

	err = ioutil.WriteFile(basepath+"/"+filename, bytes, 0664)
	if err != nil {
		return "", "", err
	}

	return filename, basepath, nil
}
