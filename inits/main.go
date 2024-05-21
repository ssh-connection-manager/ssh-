package inits

import (
	"ssh+/app/file"

	"github.com/joho/godotenv"
)

func CreateFileConnects() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	file.CreateFile()
}
