package domain

import "demo/src/products/domain/entities"

type IProduct interface {
    GetAll() ([]entities.Product, error)
	Save(product entities.Product)error
	Update(product entities.Product) error
	Delete(productID int32) error
}


