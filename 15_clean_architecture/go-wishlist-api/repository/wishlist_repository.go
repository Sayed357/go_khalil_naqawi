package repository

import (
	"go-wishlist-api/models"

	"gorm.io/gorm"
)

type GormWishlistRepository struct {
	DB *gorm.DB
}

func (r *GormWishlistRepository) GetAll() ([]models.Wishlist, error) {
	var wishlists []models.Wishlist
	if err := r.DB.Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (r *GormWishlistRepository) Create(wishlist models.Wishlist) error {
	return r.DB.Create(&wishlist).Error
}
