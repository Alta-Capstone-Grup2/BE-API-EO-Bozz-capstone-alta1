package helper

import (
	cfg "capstone-alta1/config"
	"errors"
	"fmt"
	"strings"
)

func ServiceErrorMsg(errData error) error {
	if strings.Contains(errData.Error(), "table") {
		return errors.New("Failed. Error on process. Please contact your administrator.")
	} else if strings.Contains(errData.Error(), "found") {
		return errors.New("Failed. Data not found. Please check input again.")
	} else if strings.Contains(errData.Error(), "failed on the 'required' tag") {
		return errors.New("Failed. Required field is empty. Please check input again.")
	} else if strings.Contains(errData.Error(), "foreign key constraint fails ") {
		return errors.New("Failed. Reference ID not found. Please check input again.")
	} else {
		return errors.New("Failed. Other Error. Please contact your administrator.")
	}
}

func HandlerErrorMsg(errData error) error {
	if strings.Contains(errData.Error(), "table") {
		return errors.New("Failed. Error on process. Please contact your administrator.")
	} else if strings.Contains(errData.Error(), "found") {
		return errors.New("Failed. Data not found. Please check input again.")
	} else {
		return errors.New("Failed. Other Error. Please contact your administrator.")
	}
}

// Change SHOW_LOGS to false to hide logs accross this app
func LogDebug(msg ...interface{}) {
	if cfg.SHOW_LOGS {
		fmt.Println("\n\n", msg, "\n\n")
	}
}
