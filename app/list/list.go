package list

import (
	"encoding/json"
	"log"
	"ssh+/app/file"
)

func GetConnectList() []string {
	fileConnects := file.ReadFile()
	var connections file.Connections

	err := json.Unmarshal([]byte(fileConnects), &connections)
	if err != nil {
		log.Fatal("Ошибка при разборе JSON:", err)
	}

	result := connections.GetConnectionsAlias()

	return result
}
