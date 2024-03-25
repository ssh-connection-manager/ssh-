package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFile() {
	filePath, err := GetFullPath()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(filePath)

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		createdFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}

		defer createdFile.Close()
	}
}

func ReadFile() string {
	path, err := GetFullPath()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(fContent)
}
