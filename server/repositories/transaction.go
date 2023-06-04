package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	FindTransactionId(Id int) (models.Transaction, error)
	GetTripId(Id int) (models.TripResponse, error)
	GetUserId(Id int) (models.UsersProfileResponse, error)
	GetTransByUser(Id int) ([]models.Transaction, error)
	DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var Tansactions []models.Transaction
	err := r.db.Preload("Trip.Country").Find(&Tansactions).Error

	return Tansactions, err
}
func (r *repository) GetTransByUser(Id int) ([]models.Transaction, error) {

	var Transactions []models.Transaction
	err := r.db.Where("id_user = ?", Id).Preload("User").Preload("Trip.Country").Find(&Transactions).Error

	return Transactions, err
}

func (r *repository) FindTransactionId(Id int) (models.Transaction, error) {
	var Tansactions models.Transaction
	err := r.db.Preload("Trip.Country").First(&Tansactions, Id).Error

	return Tansactions, err
}
func (r *repository) GetTripId(Id int) (models.TripResponse, error) {
	var Tansactions models.TripResponse
	err := r.db.Preload("Country").First(&Tansactions, Id).Error

	return Tansactions, err
}

func (r *repository) GetUserId(Id int) (models.UsersProfileResponse, error) {
	var Users models.UsersProfileResponse
	err := r.db.First(&Users, Id).Error

	return Users, err
}

func (r *repository) DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&Transaction).Error

	return Transaction, err
}

func (r *repository) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&Transaction).Error

	return Transaction, err
}

func (r *repository) UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&Transaction).Error

	return Transaction, err
}