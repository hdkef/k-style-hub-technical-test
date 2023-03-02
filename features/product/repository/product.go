package repository

import (
	"database/sql"
	"errors"
	"kstylehub/features/product"
	"kstylehub/features/product/models"
	"kstylehub/features/review"
)

type ProductRepository struct {
	db         *sql.DB
	reviewRepo review.ReviewRepository_
}

// GetOneByID implements product.ProductRepository_
func (u *ProductRepository) GetOneByID(id int) (product.Product, error) {
	//get product
	var data models.Product
	err := u.db.QueryRow("SELECT id_product, name_product, price FROM product where id_product = ?", id).
		Scan(&data.ID, &data.Name, &data.Price)
	if err != nil {
		return product.Product{}, errors.New("internal database error")
	}
	//get review
	review, err := u.reviewRepo.GetAllByProductID(id)
	if err != nil {
		return product.Product{}, errors.New("internal database error")
	}

	product := convertToCore(&data)
	if review != nil {
		product.Review = review
	}

	return product, nil
}

func NewProductRespository(db *sql.DB, reviewRepo review.ReviewRepository_) product.ProductRepository_ {
	return &ProductRepository{
		db:         db,
		reviewRepo: reviewRepo,
	}
}
