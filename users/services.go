package users

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (Users, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (Users, error) {
	user := Users{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
