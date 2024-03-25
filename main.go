package main

import (
	"github.com/joho/godotenv"
	"ssh+/app/file"
	"ssh+/cmd"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	file.CreateFile()
}

func main() {
	cmd.Execute()
}
