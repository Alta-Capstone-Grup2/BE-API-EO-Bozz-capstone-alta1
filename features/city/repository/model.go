package repository

import (
	"capstone-alta1/features/city"

	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	CityName string
}

// mapping

// mengubah struct model gorm ke struct core
func (dataModel *City) toCore() city.Core {
	return city.Core{
		ID:       dataModel.ID,
		CityName: dataModel.CityName,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []City) []city.Core {
	var dataCore []city.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
