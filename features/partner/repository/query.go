package repository

import (
	partner "capstone-alta1/features/partner"
	"capstone-alta1/utils/helper"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type partnerRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) partner.RepositoryInterface {
	return &partnerRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *partnerRepository) Create(input partner.Core) error {
	partnerGorm := fromCore(input)
	tx := repo.db.Create(&partnerGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements user.Repository
func (repo *partnerRepository) GetAll() (data []partner.Core, err error) {
	var partner []Partner

	tx := repo.db.Preload("User").Find(&partner)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(partner)
	return dataCore, nil
}

func (repo *partnerRepository) GetAllWithSearch(query string) (data []partner.Core, err error) {
	var partner []Partner
	tx := repo.db.Preload("User").Where("name LIKE ?", "%"+query+"%").Find(&partner)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	fmt.Println("\n\n 2 getall partner = ", partner)

	var dataCore = toCoreList(partner)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *partnerRepository) GetById(id uint) (data partner.Core, err error) {
	var partner Partner

	tx := repo.db.Preload("User").First(&partner, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = partner.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *partnerRepository) Update(input partner.Core, partnerId, userId uint) error {
	partnerGorm := fromCore(input)
	var partner Partner
	var user User

	tx := repo.db.Model(&user).Where("ID = ?", userId).Updates(&partnerGorm.User)
	yx := repo.db.Model(&partner).Where("ID = ?", partnerId).Updates(&partnerGorm)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *partnerRepository) Delete(partnerId, userId uint) error {
	var partner Partner
	var user User
	tx := repo.db.Delete(&partner, partnerId)
	yx := repo.db.Delete(&user, userId)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *partnerRepository) FindUser(email string) (result partner.Core, err error) {
	var partnerData Partner
	tx := repo.db.Where("email", email).First(&partnerData.User)
	if tx.Error != nil {
		return partner.Core{}, tx.Error
	}

	result = partnerData.toCore()

	return result, nil
}

func (repo *partnerRepository) FindPartner(partnerID uint) (result partner.Core, err error) {
	var partnerData Partner
	tx := repo.db.First(&partnerData, partnerID)
	if tx.Error != nil {
		return partner.Core{}, tx.Error
	}

	result = partnerData.toCore()

	return result, nil
}

func (repo *partnerRepository) GetServices(partnerID uint) (data []partner.ServiceCore, err error) {
	var modelData []Service
	tx := repo.db.Where("partner_id = ?", partnerID).Find(&modelData)
	// tx := repo.db.Where("service_name LIKE ?", "%"+queryServiceName+"%").Where(&Service{City: queryCity, ServiceCategory: queryPServiceCategory, ServicePrice: queryServicePrice}).Find(&modelData)

	if tx.Error != nil {
		helper.LogDebug("Partner-query-GetService | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("Partner-query-GetService | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	data = toCoreServiceList(modelData)
	return data, err
}

func (repo *partnerRepository) GetOrders(partnerID uint) (data []partner.OrderCore, err error) {
	var modelData []Order
	// tx := repo.db.Joins("JOIN partners ON services.partner_id = partners.id").Joins("JOIN orders ON orders.service_id = services.id").Find(&modelData)
	// tx := repo.db.Where("service_name LIKE ?", "%"+queryServiceName+"%").Where(&Service{City: queryCity, ServiceCategory: queryPServiceCategory, ServicePrice: queryServicePrice}).Find(&modelData)
	tx := repo.db.Raw("SELECT `orders`.`id`,`orders`.`event_name`,`orders`.`start_date`,`orders`.`end_date`,`orders`.`event_location`,`orders`.`event_address`,`orders`.`note_for_partner`,`orders`.`service_name`,`orders`.`service_price`,`orders`.`gross_ammount`,`orders`.`payment_method`,`orders`.`order_status`,`orders`.`payout_reciept_file`,`orders`.`payout_date`,`orders`.`service_id`,`orders`.`client_id` FROM services JOIN partners ON services.partner_id = partners.id JOIN orders ON orders.service_id = services.id").Where("partners.id = ?", partnerID).Scan(&modelData)

	helper.LogDebug("Partner-query-GetOrder | ModelData : ", modelData)

	if tx.Error != nil {
		helper.LogDebug("Partner-query-GetOrder | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("Partner-query-GetOrder | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	data = toOrderCoreList(modelData)
	return data, err
}
func (repo *partnerRepository) GetAdditionals(partnerID uint) (data []partner.AdditionalCore, err error) {
	var modelData []Additional
	tx := repo.db.Where("partner_id = ?", partnerID).Find(&modelData)

	if tx.Error != nil {
		helper.LogDebug("Partner-query-Get Additional | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("Partner-query-Gt Additional | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	data = toAdditionalCoreList(modelData)
	return data, err
}
func (repo *partnerRepository) GetPartnerRegisterData(queryCompanyName, queryPICName, queryPartnerStatus string) (data []partner.Core, err error) {
	var tx *gorm.DB
	var modelData []Partner
	if queryCompanyName == "" && queryPICName == "" && queryPartnerStatus == "" {
		tx = repo.db.Preload("User").Find(&modelData)
	} else {
		tx = repo.db.Preload("User").Where("company_name LIKE ?", "%"+queryCompanyName+"%").Where(&Partner{User: User{Name: queryPICName}, VerificationStatus: queryPartnerStatus}).Find(&modelData)
	}

	if tx.Error != nil {
		helper.LogDebug("Partner-query-GetPartnerRegisterData | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("Partner-query-GetPartnerRegisterData | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	data = toCoreList(modelData)
	return data, err
}
func (repo *partnerRepository) GetPartnerRegisterDataByID(partnerID uint) (data partner.Core, err error) {
	return data, err
}
func (repo *partnerRepository) UpdatePartnerVerifyStatus(partnerID uint) (data partner.Core, err error) {
	return data, err
}
func (repo *partnerRepository) UpdateOrderConfirmStatus(orderID uint) (data partner.Core, err error) {
	return data, err
}
