package json

import (
	"os"
	"ssh+/app/file"
)

func GetPathToConnectFile() string {
	fullPath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		panic(err)
	}

	return fullPath
}
