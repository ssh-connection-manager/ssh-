package convert

import (
	"os"
	"strings"
)

func StringToArray(str string) []string {
	resultArray := strings.Split(str, os.Getenv("SEPARATOR"))

	return resultArray
}

func ArrayToMap(arr []string) map[string]string {
	resultMap := make(map[string]string)

	firstElem := arr[0]

	for _, value := range arr {
		if strings.Contains(value, ";") {
			firstElem = value
		}

		resultMap[firstElem] = value
	}

	return resultMap
}
