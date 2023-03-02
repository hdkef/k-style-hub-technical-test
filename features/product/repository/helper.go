package repository

import (
	"kstylehub/features/product"
	"kstylehub/features/product/models"
)

func convertToModels(u *product.Product) models.Product {
	return models.Product{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
	}
}

func convertToCore(u *models.Product) product.Product {
	return product.Product{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
	}
}
