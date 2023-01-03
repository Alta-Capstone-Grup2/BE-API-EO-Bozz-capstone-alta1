package helper

import (
	cfg "capstone-alta1/config"
	"errors"
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

func GetDateTimeFormatedStr(dateTimeStr string) string {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return ""
	}

	return dateTimeData.In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

<<<<<<< Updated upstream
=======
func GetDateTimeFormated(dateTimeInput time.Time) string {
	return dateTimeInput.In(location).Format(cfg.DEFAULT_DATETIME_LAYOUT)
}

func GetDateFormated(dateTimeInput time.Time) string {
	return dateTimeInput.In(location).Format(cfg.DEFAULT_DATE_LAYOUT)
}

func GetDateTimeFormatedToTime(dateTimeStr string) time.Time {
	dateTimeData, errParse := time.Parse(cfg.DEFAULT_DATETIME_LAYOUT, dateTimeStr)
	LogDebug("GetDateTimeFormatedToTime | input = ", dateTimeStr, " 	result = ", dateTimeData)
	if errParse != nil {
		LogDebug("Failed parse date time.")
		return time.Time{}
	}

	return dateTimeData.In(location)
}

func ValidateDateTimeFormatedToTime(dateTimeStr string, dateTimeLayout string) error {
	dateTimeData, errParse := time.Parse(dateTimeLayout, dateTimeStr)
	if errParse != nil {
		LogDebug("Failed Parse Datetime. GetDateTimeFormatedToTime | input = ", dateTimeStr, " format ", dateTimeLayout, " result = ", dateTimeData)
		return errors.New("Incorrect Datetime Format. Please check your input. Use format " + dateTimeLayout)
	}

	return nil
}

>>>>>>> Stashed changes
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
