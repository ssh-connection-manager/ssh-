package create

import (
	"ssh+/app/json"
)

func Connect(alias, address, login, password string) {
	var connections json.Connections

	connect := json.Connect{
		Alias:    alias,
		Address:  address,
		Login:    login,
		Password: password,
	}

	connections.WriteConnectToJson(connect)
}
