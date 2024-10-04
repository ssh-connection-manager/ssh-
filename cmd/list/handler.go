package list

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/json"
	"ssh+/app/output"
)

func Show() [][]string {
	var connections json.Connections

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")

	connectionsData, err := connections.GetDataForListConnect(filePath, fileName)
	if err != nil {
		output.GetOutError(err.Error())
	}

	return connectionsData
}
