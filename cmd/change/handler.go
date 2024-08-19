package change

import (
	"ssh+/app/json"
	"ssh+/app/output"

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

	connections.UpdateConnectJson(oldAlias, connect)
}

func ExistByIndex(alias string) {
	_, err := connects.ExistConnectJsonByIndex(alias)

	if err != nil {
		output.GetOutError("No connection found")
	}
}
