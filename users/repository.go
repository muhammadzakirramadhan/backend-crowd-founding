package users

import "gorm.io/gorm"

type Repository interface {
	Save(users Users) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Save(users Users) (Users, error) {
	err := r.db.Create(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
