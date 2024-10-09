package list

import (
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
)

func Show() [][]string {
	var connections json.Connections

	connectionsData, err := connections.GetDataForListConnect()
	if err != nil {
		output.GetOutError(err.Error())
	}

	return connectionsData
}
