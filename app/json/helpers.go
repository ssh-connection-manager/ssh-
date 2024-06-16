package json

import (
	"ssh+/app/file"

	"github.com/spf13/viper"
)

func GetPathToConnectFile() string {
	fullPath := file.GetFullPath(viper.GetString("NameFileConnects"))

	return fullPath
}

func CreateBaseJsonDataToFile() {
	connections := Connections{
		Connects: []Connect{},
	}

	file.WriteFile(GetPathToConnectFile(), connections.deserializationJson())
}
