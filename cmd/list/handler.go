package list

import (
	"ssh+/app/json"
)

func GetConnectsList() []string {
	var connections json.Connections

	return connections.GetConnectionsAlias()
}
