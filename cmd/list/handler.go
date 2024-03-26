package list

import (
	"ssh+/app/file"
)

func GetConnectsList() []string {
	var connections file.Connections
	dataConnectsInFile := file.ReadFile()

	return connections.GetConnectionsAlias(dataConnectsInFile)
}
