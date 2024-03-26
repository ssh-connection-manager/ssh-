package create

import "ssh+/app/file"

func CreateConnect(alias, address, login, password string) {
	var connections file.Connections

	connect := file.Connect{
		Alias:    alias,
		Address:  address,
		Login:    login,
		Password: password,
	}

	connections.WriteConnectToJson(connect)
}
