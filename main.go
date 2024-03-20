package main

import (
	"github.com/joho/godotenv"
	"os"
	"ssh+/app/file"
	"ssh+/cmd"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	file := file.File{
		Name: os.Getenv("NAME_FILE"),
		Path: os.Getenv("PATH_FILE"),
	}

	file.CreateFile()
}

func main() {
	cmd.Execute()
}
