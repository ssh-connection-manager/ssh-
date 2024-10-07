package delete

import (
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/json"
)

func Connect(alias string) {
	var connects json.Connections

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")
	fileKey := viper.GetString("NameFileCryptKey")

	err := connects.DeleteConnectToJson(alias, filePath, fileName, fileKey)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
