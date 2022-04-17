package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByID(id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	SaveDetails(details []TransactionDetail) ([]TransactionDetail, error)
	Update(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
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

func (r *repository) SaveDetails(details []TransactionDetail) ([]TransactionDetail, error) {
	err := r.db.Create(&details).Error
	return details, err
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}
