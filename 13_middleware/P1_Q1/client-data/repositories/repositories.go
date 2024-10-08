package repositories

import "client-data/models"

type ContentRepository interface {
	GetAll() ([]models.Content, error)
	GetByID(id string) (models.Content, error)
	Create(contentReq models.ContentRequest) (models.Content, error)
	Update(contentReq models.ContentRequest, id string) (models.Content, error)
	Delete(id string) error
}

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(categoryReq models.CategoryRequest) (models.Category, error)
	Update(categoryReq models.CategoryRequest, id string) (models.Category, error)
	Delete(id string) error
}

type UserRepository interface {
	Register(registerReq models.RegisterRequest) (models.User, error)
	GetByEmail(loginReq models.LoginRequest) (models.User, error)
	GetUserInfo(id string) (models.User, error)
}
