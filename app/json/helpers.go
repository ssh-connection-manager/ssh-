package json

import (
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/file"
)

func GetPathToConnectFile() string {
	fullPath := file.GetFullPath(
		viper.GetString("FullPathConfig"),
		viper.GetString("NameFileConnects"))

	return fullPath
}

func CreateBaseJsonDataToFile() {
	connections := Connections{
		Connects: []Connect{},
	}

	err := file.WriteFile(GetPathToConnectFile(), connections.deserializationJson())
	if err != nil {
		output.GetOutError("Error create base connect file")
	}
}
