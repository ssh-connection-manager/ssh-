package create

import (
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/time"
)

func Connect(alias, address, login, password string) {
	var connections json.Connections

	timeNow := time.GetTime()

	connect := json.Connect{
		Alias:     alias,
		Address:   address,
		Login:     login,
		Password:  password,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")
	fileKey := viper.GetString("NameFileCryptKey")

	err := connections.WriteConnectToJson(connect, filePath, fileName, fileKey)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
