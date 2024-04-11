package delete

import "ssh+/app/file"

func DeleteConnect(alias string) {
	var connects file.Connections

	connects.DeleteConnectToJson(alias)
}
