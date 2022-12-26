package repository

import (
	_discussion "capstone-alta1/features/discussion"
	"time"
)

// struct gorm model
type Discussion struct {
	ID        uint
	Comment   string
	PartnerID uint
	ClientID  uint
	ServiceID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _discussion.Core) Discussion {
	modelData := Discussion{
		Comment:   dataCore.Comment,
		PartnerID: dataCore.PartnerID,
		ClientID:  dataCore.ClientID,
		ServiceID: dataCore.ServiceID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Discussion) toCore() _discussion.Core {
	return _discussion.Core{
		ID:        dataModel.ID,
		Comment:   dataModel.Comment,
		PartnerID: dataModel.PartnerID,
		ClientID:  dataModel.ClientID,
		ServiceID: dataModel.ServiceID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Discussion) []_discussion.Core {
	var dataCore []_discussion.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
