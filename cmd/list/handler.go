package list

import (
	"ssh+/app/json"
)

func Show() [][]string {
	var connections json.Connections

	return connections.GetDataForListConnect()
}
