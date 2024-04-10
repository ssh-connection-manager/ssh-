package connect

import (
	"ssh+/app/connect"
	"ssh+/app/file"
)

func ConsoleConnect(alias string) {
	var connections file.Connections

	connect.Connect(&connections, alias)
}
