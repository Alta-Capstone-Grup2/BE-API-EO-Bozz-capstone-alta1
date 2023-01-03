package helper

import (
	cfg "capstone-alta1/config"
	"strings"
	"time"
)

var location, _ = time.LoadLocation(cfg.DEFAULT_DATETIME_LOCATION)

func GetDateNow() string {
	return time.Now().In(location).Format(cfg.DEFAULT_DATE_LAYOUT)
}

func GetDateTimeNow() string {
	return time.Now().In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

func GetDateTimeNowZUTC7() string {
	return time.Now().In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT_Z)
}

func GetDateTimeFormated(dateTimeStr string) string {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return ""
	}

	return dateTimeData.In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

func AddDateTimeFormated(dateTimeStr string, years, months, days int) string {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return ""
	}

	addedDateTime := dateTimeData.AddDate(years, months, days)

	return addedDateTime.Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

func AddDateTimeFormatedZUTC7(dateTimeStr string, years, months, days int) string {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return ""
	}

	addedDateTime := dateTimeData.AddDate(years, months, days)

	return addedDateTime.Format(cfg.DEFAULT_DATETIME_LAYOUT_Z)
}

func GetDateNowShort() string {
	return strings.Replace(GetDateNow(), "-", "", -1)
}
