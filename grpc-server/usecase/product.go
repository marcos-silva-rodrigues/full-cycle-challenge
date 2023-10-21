package usecase

import (
	"github.com/full-cycle-challenge/model"
	"github.com/jinzhu/gorm"
)

type ProductUseCase struct {
	DB *gorm.DB
}

func (p *ProductUseCase) Create(name string, description string, price float32) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	result := p.DB.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (p *ProductUseCase) FindAll() ([]model.Product, error) {
	var products []model.Product

	p.DB.Find(&products)
	return products, nil
}
