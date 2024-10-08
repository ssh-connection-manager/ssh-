package list

import (
	"ssh+/app/output"

	"github.com/ssh-connection-manager/json"
)

func Show() [][]string {
	var connections json.Connections

	connectionsData, err := connections.GetDataForListConnect()
	if err != nil {
		output.GetOutError(err.Error())
	}

	return connectionsData
}
