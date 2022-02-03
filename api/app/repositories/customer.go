package repositories

import (
	"api/app/database"
	"api/app/models"

	"gorm.io/gorm"
)

var CustomerRepo CustomerRepoInterface = &customerRepo{}

type CustomerRepoInterface interface {
	GetAll() []models.Customer
	Init()
}

type customerRepo struct {
	db *gorm.DB
}

func (customerRepo *customerRepo) Init() {
	customerRepo.db = database.Database
}

func NewCustomerRepository(db *gorm.DB) CustomerRepoInterface {
	return &customerRepo{db: db}
}

func (customerRepo *customerRepo) GetAll() []models.Customer {
	var customers []models.Customer
	customerRepo.db.Table("customer").Find(&customers)
	return customers
}
