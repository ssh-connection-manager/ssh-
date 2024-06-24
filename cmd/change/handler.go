package change

import (
	"ssh+/app/json"
	"ssh+/app/output"
)

var connects json.Connections

func Connect(
	oldAlias string,
	alias, address, login, password string) {
	var connections json.Connections

	connect := json.Connect{
		Alias:    alias,
		Address:  address,
		Login:    login,
		Password: password,
	}

	connections.UpdateConnectJson(oldAlias, connect)
}

func ExistByIndex(alias string) {
	_, err := connects.ExistConnectJsonByIndex(alias)

	if err != nil {
		output.GetOutError("No connection found")
	}
}
