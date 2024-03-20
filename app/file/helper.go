package file

import (
	"os"
)

func getPathToFile() string {
	return os.Getenv("PATH_FILE") +
		os.Getenv("SEPARATOR") +
		os.Getenv("NAME_FILE")
}

func GetFullPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := homePath + getPathToFile()
	return path, nil
}
