package delivery

import (
	cfg "capstone-alta1/config"
	"capstone-alta1/utils/helper"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(structIntf interface{}) error {
	switch structIntf.(type) {
	case CheckAvailabilityRequest:
		obj := structIntf.(CheckAvailabilityRequest)
		return validation.ValidateStruct(&obj,
			validation.Field(&obj.StartDate, validation.Required, validation.Date(cfg.DEFAULT_DATE_LAYOUT).Min(time.Now().AddDate(0, 0, -1)).Max(helper.GetDateTimeFormatedToTime(obj.EndDate+" 00:00:00"))),
			validation.Field(&obj.EndDate, validation.Required, validation.Date(cfg.DEFAULT_DATE_LAYOUT).Min(helper.GetDateTimeFormatedToTime(obj.StartDate+" 00:00:00"))),
		)
	}
	return nil
}

// validation.When(or.ServiceID == 0).Error(`Service ID cannot 0`))
