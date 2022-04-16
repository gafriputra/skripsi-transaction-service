package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByUserID(userID int) ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	SaveDetail(detail TransactionDetail) (TransactionDetail, error)
	Update(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("created_at desc").Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetByID(id int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("id = ?", id).Find(&transaction).Error
	return transaction, err
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) SaveDetail(detail TransactionDetail) (TransactionDetail, error) {
	err := r.db.Create(&detail).Error
	return detail, err
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}
