package create

import (
	"ssh+/app/json"
	"ssh+/app/timeNow"
)

func Connect(alias, address, login, password string) {
	var connections json.Connections

	time := timeNow.GetTime()

	connect := json.Connect{
		Alias:     alias,
		Address:   address,
		Login:     login,
		Password:  password,
		CreatedAt: time,
		UpdatedAt: time,
	}

	connections.WriteConnectToJson(connect)
}
