package change

import (
	"github.com/spf13/viper"
	"ssh+/app/output"

	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/time"
)

var connects json.Connections

func Connect(
	oldAlias string,
	alias, address, login, password string) {
	var connections json.Connections

	timeNow := time.GetTime()

	connect := json.Connect{
		Alias:     alias,
		Address:   address,
		Login:     login,
		Password:  password,
		UpdatedAt: timeNow,
	}

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")
	fileKey := viper.GetString("NameFileCryptKey")

	err := connections.UpdateConnectJson(oldAlias, connect, filePath, fileName, fileKey)
	if err != nil {
		output.GetOutError("err update")
	}
}

func ExistByIndex(alias string) {
	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")
	fileKey := viper.GetString("NameFileCryptKey")

	_, err := connects.ExistConnectJsonByIndex(alias, filePath, fileName, fileKey)

	if err != nil {
		output.GetOutError("No connection found")
	}
}
