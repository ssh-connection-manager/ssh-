package delete

import (
	"ssh+/app/json"
)

func Connect(alias string) {
	var connects json.Connections

	connects.DeleteConnectToJson(alias)
}
