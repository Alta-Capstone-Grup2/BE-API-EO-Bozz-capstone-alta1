package helper

import (
	"encoding/json"

	"github.com/leekchan/accounting"
)

func ConvToJson(structData interface{}) string {

	jsonData, errConv := json.Marshal(structData)
	if errConv != nil {
		LogDebug("Tools - ConvToJson | Failed marshal.")
		return ""
	}
	return string(jsonData)
}

func FormatCurrencyIDR(moneyInput uint) string {
	currFormat := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return currFormat.FormatMoney(moneyInput)
}
