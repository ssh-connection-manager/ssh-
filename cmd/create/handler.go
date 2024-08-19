package create

import (
	"ssh+/app/json"

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

	connections.WriteConnectToJson(connect)
}
