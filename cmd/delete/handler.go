package delete

import (
	"ssh+/app/output"

	"github.com/ssh-connection-manager/json"
)

func Connect(alias string) {
	var connects json.Connections

	err := connects.DeleteConnectToJson(alias)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
