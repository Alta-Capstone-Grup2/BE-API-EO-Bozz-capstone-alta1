package helper

import "encoding/json"

func ConvToJson(structData interface{}) string {

	jsonData, errConv := json.Marshal(structData)
	if errConv != nil {
		LogDebug("Tools - ConvToJson | Failed marshal.")
		return ""
	}
	return string(jsonData)
}
