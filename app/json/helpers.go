package json

import (
	"ssh+/app/file"

	"github.com/spf13/viper"
)

func GetPathToConnectFile() string {
	fullPath, err := file.GetFullPath(viper.GetString("NameFileConnects"))
	if err != nil {
		panic(err)
	}

	return fullPath
}

func CreateBaseJsonDataToFile() {
	connections := Connections{
		Connects: []Connect{},
	}

	file.WriteFile(GetPathToConnectFile(), connections.deserializationJson())
}
