package inits

import (
	"github.com/joho/godotenv"
	"ssh+/app/file"
)

func CreateFileConnects() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	file.CreateFile()
}
