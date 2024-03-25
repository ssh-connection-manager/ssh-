package list

import (
	"encoding/json"
	"ssh+/app/file"
	"ssh+/output"
)

func GetConnectList() []string {
	fileConnects := file.ReadFile()
	var connections file.Connections

	err := json.Unmarshal([]byte(fileConnects), &connections)
	if err != nil {
		output.GetOutError("пустой файл")
	}

	result := connections.GetConnectionsAlias()

	return result
}
