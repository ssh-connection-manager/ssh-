package connect

import (
	"ssh+/app/connect"
	"ssh+/app/json"
)

func ConsoleConnect(alias string) {
	var connections json.Connections

	connect.Connect(&connections, alias)
}
