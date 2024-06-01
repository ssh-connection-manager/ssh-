package delete

import (
	"ssh+/app/json"
)

func DeleteConnect(alias string) {
	var connects json.Connections

	connects.DeleteConnectToJson(alias)
}
