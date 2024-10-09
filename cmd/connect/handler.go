package connect

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/connect"
	"github.com/ssh-connection-manager/file"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
)

func Ssh(alias string) {
	var connections json.Connections

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")

	fileConnect := file.File{Path: filePath, Name: fileName}

	err := connect.Ssh(&connections, alias, fileConnect)
	if err != nil {
		output.GetOutError("err connect")
	}
}
