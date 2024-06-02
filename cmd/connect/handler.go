package connect

import (
	"ssh+/app/connect"
	"ssh+/app/json"
)

func Ssh(alias string) {
	var connections json.Connections

	connect.Ssh(&connections, alias)
}
