package users

import "gorm.io/gorm"

type Repository interface {
	Save(users Users) (Users, error)
	FindByEmail(email string) (Users, error)
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

func (r repository) FindByEmail(email string) (Users, error) {
	var users Users
	err := r.db.Where("email = ?", email).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
