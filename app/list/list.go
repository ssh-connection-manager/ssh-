package list

import (
	"ssh+/app/file"
)

func GetConnectsList() []string {
	var connections file.Connections
	dataConnectsInFile := file.ReadFile()

	return connections.GetConnectionsAlias(dataConnectsInFile)
}

func AddConnect(alias, address, login, password string) {
	var connections file.Connections

	connect := file.Connect{
		Alias:    alias,
		Address:  address,
		Login:    login,
		Password: password,
	}

	connections.WriteConnectToJson(connect)
}
