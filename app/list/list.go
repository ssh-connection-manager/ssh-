package list

import (
	"fmt"
	"ssh+/app/convert"
	"ssh+/app/file"
)

func getMapConnect() map[string]string {
	fileConnects := file.ReadFile()
	mapConnects := convert.ArrayToMap(convert.StringToArray(fileConnects))

	return mapConnects
}

func GetConnectList() string {
	mapConnects := getMapConnect()
	fmt.Println(mapConnects)
	return ""
}
