package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"ssh+/app/output"
)

func CreateFile() {
	filePath, err := getFullPath()
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
	path, err := getFullPath()
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

func WriteFile(rowData []byte) {
	path, err := getFullPath()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, rowData, 0)

	if err != nil {
		output.GetOutError("ошибка при записи в файл")
	}
}
