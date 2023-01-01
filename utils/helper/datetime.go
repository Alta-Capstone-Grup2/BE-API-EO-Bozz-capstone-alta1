package helper

import (
	cfg "capstone-alta1/config"
	"time"
)

var location, _ = time.LoadLocation(cfg.DEFAULT_DATETIME_LOCATION)

func GetDateNow() string {
	return time.Now().In(location).Format(cfg.DEFAULT_DATE_LAYOUT)
}

func GetDateTimeNow() string {
	return time.Now().In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

func GetDateTimeFormated(dateTimeStr string) string {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return ""
	}

	return dateTimeData.In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)

}
