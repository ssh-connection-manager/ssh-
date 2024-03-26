package list

import (
	"ssh+/app/file"
)

func GetConnectsList() []string {
	var connections file.Connections

	return connections.GetConnectionsAlias()
}
