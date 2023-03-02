package service

import (
	"errors"
	"kstylehub/features/product"
)

type ProductService struct {
	repo product.ProductRepository_
}

// GetOneByID implements product.ProductService_
func (u *ProductService) GetOneByID(id int) (product.Product, error) {
	//validate id
	if id <= 0 {
		return product.Product{}, errors.New("invalid id")
	}
	return u.repo.GetOneByID(id)
}

func NewProductService(repo product.ProductRepository_) product.ProductService_ {
	return &ProductService{
		repo: repo,
	}
}
