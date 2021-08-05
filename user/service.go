package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface { //! bisnis logic
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserByID(ID int) (User, error)
}

type service struct { //! memanggil repository
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser ,err
	}

	return newUser, nil
}

func(s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func(s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, _ := s.repository.FindByEmail(email)

	if user.ID == 0 { //! email tidak di temukan atau bisa di daftarkan
		return true, nil
	}

	return false, nil //! email sudah di gunakan
}

func(s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}

	return user, nil
}