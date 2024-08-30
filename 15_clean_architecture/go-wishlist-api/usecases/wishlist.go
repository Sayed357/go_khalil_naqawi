package usecases

import "go-wishlist-api/models"

type WishlistRepository interface {
	GetAll() ([]models.Wishlist, error)
	Create(wishlist models.Wishlist) error
}

type WishlistUseCase struct {
	Repo WishlistRepository
}

func (uc *WishlistUseCase) GetAllWishlists() ([]models.Wishlist, error) {
	return uc.Repo.GetAll()
}

func (uc *WishlistUseCase) CreateWishlist(wishlist models.Wishlist) error {
	return uc.Repo.Create(wishlist)
}
