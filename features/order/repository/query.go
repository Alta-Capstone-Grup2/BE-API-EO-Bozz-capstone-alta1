package repository

import (
	_order "capstone-alta1/features/order"
	"capstone-alta1/utils/helper"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) _order.RepositoryInterface {
	return &orderRepository{
		db: db,
	}
}

func (repo *orderRepository) Create(inputOrder _order.Core, inputDetail []_order.DetailOrder) (data _order.Core, err error) {
	orderGorm := fromCore(inputOrder)
	detailorderGorm := fromDetailOrderList(inputDetail)

	// // datetime layout
	// layoutDefault := "2006-01-02 15:04:05"
	// // //init the loc
	// // loc, _ := time.LoadLocation("Asia/Jakarta")
	// // //set timezone,
	// // // now := time.Now().In(loc).Format(layoutDefault)

	// var errParse error
	// orderGorm.PayoutDate, errParse = time.Parse(layoutDefault, "0000-00-00 00:00:00.000")
	// if errParse != nil {
	// 	helper.LogDebug("Order - query - Create | Error parse = ", errParse)
	// 	return errors.New("Failed insert. Parse payoutdate failed.")
	// }

	// orderGorm.PayoutDate = time.Time{}

	//Check Service is exist
	cx := repo.db.Find(&Service{}, orderGorm.ServiceID)
	if cx.Error != nil {
		helper.LogDebug("Order - query - Create | Check Service is exist. Error  = ", cx.Error)
		return _order.Core{}, cx.Error
	}

	// Insert Order Process
	tx := repo.db.Create(&orderGorm) // proses insert data
	if tx.Error != nil {
		helper.LogDebug("Order - query - Create | Error execute query order. Error  = ", tx.Error)
		if strings.Contains(tx.Error.Error(), "Cannot add or update a child row: a foreign key constraint fails") {
			return _order.Core{}, errors.New("Service Data or Additional Data Not Found. Please Check your input.")
		}
		return _order.Core{}, tx.Error
	}
	helper.LogDebug("Order - query - create | Row Affected query order : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return _order.Core{}, errors.New("insert order failed")
	}

	// Insert order id to detail id struct model
	for idx := range detailorderGorm {
		detailorderGorm[idx].OrderID = orderGorm.ID
	}

	helper.LogDebug("Order - query - create | Add order id to order detail slice. []Detail Order : ", helper.ConvToJson(detailorderGorm))

	yx := repo.db.Create(&detailorderGorm) // proses insert data
	if yx.Error != nil {
		helper.LogDebug("Order - query - Create | Error execute query detail order. Error  = ", yx.Error)
		if strings.Contains(yx.Error.Error(), "Cannot add or update a child row: a foreign key constraint fails") {
			return _order.Core{}, errors.New("Service Data or Additional Data Not Found. Please Check your input.")
		}
		return _order.Core{}, yx.Error
	}
	helper.LogDebug("Order - query - Create | Row Affected query detail order : ", yx.RowsAffected)
	if yx.RowsAffected == 0 {

		return _order.Core{}, errors.New("insert detail order failed")
	}

	zx := repo.db.Model(&orderGorm).Updates(Order{GrossAmmount: orderGorm.GrossAmmount, OrderStatus: orderGorm.OrderStatus}) // proses insert data
	if zx.Error != nil {
		helper.LogDebug("Order - query - Create | Error execute query update gross amount. Error  = ", zx.Error)
		return _order.Core{}, zx.Error
	}
	helper.LogDebug("Order - query - Create | Row Affected query update gross amount : ", zx.RowsAffected)
	if yx.RowsAffected == 0 {

		return _order.Core{}, errors.New("update gross ammount failed")
	}

	data = orderGorm.toCore()

	return data, nil
}

func (repo *orderRepository) GetServiceByID(serviceID uint) (data _order.Service, err error) {
	var serviceData Service
	tx := repo.db.Find(&serviceData, serviceID)

	if tx.Error != nil {
		helper.LogDebug("Order - query - GetServiceID | Error execute query. Error  = ", tx.Error)
		return _order.Service{}, tx.Error
	}

	helper.LogDebug("Order - query - GetServiceID | Row Affected query get additional data : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return _order.Service{}, tx.Error
	}

	helper.LogDebug("Order - query - GetServiceID | serviceData = ", serviceData)
	data = serviceData.toCoreGetById()

	return data, tx.Error
}

func (do *DetailOrder) BeforeCreate(tx *gorm.DB) (err error) {
	// Check service_additional_id is connected to service_id in order
	var serviceAdditionalData ServiceAdditional
	txCheckService := tx.Raw("SELECT service_id FROM service_additionals WHERE id = ? AND service_id IN (SELECT service_id FROM orders WHERE id = ?);", do.ServiceAdditionalID, do.OrderID).Find(&serviceAdditionalData)

	if txCheckService.Error != nil {
		helper.LogDebug("Order - query - BeforeCreate Order Detail | Error execute check service connected. Error  = ", txCheckService.Error)
		return txCheckService.Error
	}

	helper.LogDebug("Order - query - BeforeCreate Order Detail | Row Affected querycheck service connected: ", txCheckService.RowsAffected)
	if txCheckService.RowsAffected == 0 {
		return errors.New("Service at Detail Order didn't match with service at Order. Please check input again.")
	}

	var additionalData Additional
	txBeforeCreate := tx.Raw("SELECT `additionals`.`additional_name`, `additionals`.`additional_price`  FROM `additionals` JOIN `service_additionals` ON `additionals`.`id` = `service_additionals`.`additional_id` WHERE `service_additionals`.`id` = ?;", do.ServiceAdditionalID).Find(&additionalData)

	if txBeforeCreate.Error != nil {
		helper.LogDebug("Order - query - BeforeCreate Order Detail | Error execute query. Error  = ", txBeforeCreate.Error)
		return txBeforeCreate.Error
	}

	helper.LogDebug("Order - query - BeforeCreate Order Detail | Row Affected query get additional data : ", txBeforeCreate.RowsAffected)
	if txBeforeCreate.RowsAffected == 0 {
		return txBeforeCreate.Error
	}

	helper.LogDebug("Order - query - BeforeCreate Order Detail| additionalData = ", additionalData)

	do.AdditionalName = additionalData.AdditionalName
	do.AdditionalPrice = additionalData.AdditionalPrice
	do.DetailOrderTotal = do.Qty * do.AdditionalPrice

	helper.LogDebug("Order - query - BeforeCreate | additionalData = ", additionalData)
	return txBeforeCreate.Error
}

func (repo *orderRepository) GetAll(query string) (data []_order.OrderJoinPartner, err error) {
	var results []OrderJoinPartner

	if query == "" {
		tx := repo.db.Raw("SELECT `orders`.*, `partners`.`id` AS `partner_id`, `partners`.`company_name`, `partners`.`bank_name`, `partners`.`bank_account_number`, `partners`.`bank_account_name` FROM `partners` JOIN `services` ON `partners`.`id` = `services`.`partner_id` JOIN `orders` ON `orders`.`service_id` = `services`.`id`;").Scan(&results)
		if tx.Error != nil {
			helper.LogDebug("Order - query - GetAll | Error Query Find = ", tx.Error)
			return nil, tx.Error
		}


		helper.LogDebug("Order - query - GetAll | Resukt data : ", results)
		helper.LogDebug("Order - query - GetAll | Row Affected query get additional data : ", tx.RowsAffected)
		if tx.RowsAffected == 0 {
			return nil, tx.Error
		}

	} else {
		repo.GetAllWithSearch(query)
	}

	var dataCore = toOrderJoinPartnerCoreList(results)
	return dataCore, nil
}

func (repo *orderRepository) GetAllWithSearch(eventName string) (data []_order.OrderJoinPartner, err error) {
	var results []OrderJoinPartner

	tx := repo.db.Raw("SELECT `orders`.*, `partners`.`id` AS `partner_id`, `partners`.`company_name`, `partners`.`bank_name`, `partners`.`bank_account_number`, `partners`.`bank_account_name` FROM `partners` JOIN `services` ON `partners`.`id` = `services`.`partner_id` JOIN `orders` ON `orders`.`service_id` = `services`.`id`;").Where("event_name LIKE ?", "%"+eventName+"%").Scan(&results)
	if tx.Error != nil {
		helper.LogDebug("Order - query - GetAll | Error Query Find = ", tx.Error)
		return nil, tx.Error
	}

	helper.LogDebug("Order - query - GetAll | Result data : ", results)
	helper.LogDebug("Order - query - GetAll | Row Affected query get additional data : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	var dataCore = toOrderJoinPartnerCoreList(results)
	return dataCore, nil
}

func (repo *orderRepository) GetById(id uint) (data _order.Core, dataDetail []_order.DetailOrder, err error) {
	var order Order
	var detail []DetailOrder

	tx := repo.db.First(&order, id)
	yx := repo.db.Where("order_id = ?", id).Find(&detail)

	if tx.Error != nil && yx.Error != nil {
		return data, dataDetail, tx.Error
	}

	if tx.RowsAffected == 0 && yx.RowsAffected == 0 {
		return data, dataDetail, tx.Error
	}

	var dataCore = order.toCoreOrder()
	var dataCoreDetail = toCoreDetailOrderList(detail)
	return dataCore, dataCoreDetail, nil
}

func (repo *orderRepository) UpdateStatusCancel(input _order.Core, id uint) error {
	resultGorm := fromCoreStatus(input)
	var result Order
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm.OrderStatus) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *orderRepository) UpdateStatusPayout(input _order.Core, id uint) error {
	resultGorm := fromCorePayout(input)
	var result Order
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		helper.LogDebug("Order - query - UpdateStatusPayout | Error execute query. Error  = ", tx.Error)
		return tx.Error
	}

	helper.LogDebug("Order - query - BeforeCreate Order Detail | Row Affected query get additional data : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}

	helper.LogDebug("Order - query -  UpdateStatusPayout | Order data : ", result)

	return nil
}

// UPDATE STATUS ORDER AFTER PAYMENT MIDTRANS
func (rq *orderRepository) UpdateMidtrans(input _order.Core) error {
	orderGorm := fromCore(input)
	if err := rq.db.Where("midtrans_transaction_id = ?", orderGorm.MidtransTransactionID).Updates(&orderGorm).Error; err != nil {
		return err
	}

	return nil
}
