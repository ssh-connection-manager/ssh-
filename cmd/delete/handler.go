package delete

import "ssh+/app/file"

func DeleteConnect() {
	var connects file.Connections

	connects.DeleteConnectToJson()
}
