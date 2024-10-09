package delete

import (
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
)

func Connect(alias string) {
	var connects json.Connections

	err := connects.DeleteConnectToJson(alias)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
