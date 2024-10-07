package connect

import (
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/connect"
	"github.com/ssh-connection-manager/json"
)

func Ssh(alias string) {
	var connections json.Connections

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")
	fileKey := viper.GetString("NameFileCryptKey")

	err := connect.Ssh(&connections, alias, filePath, fileName, fileKey)
	if err != nil {
		output.GetOutError("err connect")
	}
}
