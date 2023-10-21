package model

import "github.com/asaskevich/govalidator"

type Product struct {
	ID          int32   `json:"id" gorm:"primaryKey;autoIncrement=true" valid:"-"`
	Name        string  `json:"name" gorm:"type:varchar(60)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	Price       float32 `json:"type:float" gorm:"type:float" valid:"notnull"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
