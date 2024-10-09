package change

import (
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
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

	err := connections.UpdateConnectJson(oldAlias, connect)
	if err != nil {
		output.GetOutError("err update")
	}
}

func ExistByIndex(alias string) {
	_, err := connects.ExistConnectJsonByIndex(alias)

	if err != nil {
		output.GetOutError("No connection found")
	}
}
