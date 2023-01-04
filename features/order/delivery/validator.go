package delivery

import (
	cfg "capstone-alta1/config"
	"capstone-alta1/utils/helper"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(structIntf interface{}) error {
	switch structIntf.(type) {
	case OrderRequest:
		obj := structIntf.(OrderRequest)
		return validation.ValidateStruct(&obj,
			validation.Field(&obj.EventName, validation.Required, validation.Length(5, 50)),
			validation.Field(&obj.StartDate, validation.Required, validation.Date(cfg.DEFAULT_DATE_LAYOUT).Min(time.Now()).Max(helper.GetDateTimeFormatedToTime(obj.EndDate+" 00:00:00"))),
			validation.Field(&obj.EndDate, validation.Required, validation.Date(cfg.DEFAULT_DATE_LAYOUT).Min(helper.GetDateTimeFormatedToTime(obj.StartDate+" 00:00:00"))),
			validation.Field(&obj.EventLocation, validation.Required, validation.Length(5, 50)),
			validation.Field(&obj.EventAddress, validation.Required, validation.Length(5, 300)),
			validation.Field(&obj.NotesForPartner, validation.Length(0, 50)),
			validation.Field(&obj.PaymentMethod, validation.Required, validation.In("va permata", "va bca", "va bni", "va bri")),
			validation.Field(&obj.ServiceID, validation.Required),
		)
	}
	return nil
}

// validation.When(or.ServiceID == 0).Error(`Service ID cannot 0`))
