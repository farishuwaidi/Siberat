package service

import (
	"Siberat/dto"
	"Siberat/entity"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	// cek email pada database
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email alredy exist"}
	}
	
	// konfirmasi password
	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "Password dosn't match"}
	}

	passwrodHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User {
		Name: req.Name,
		Email: req.Email,
		Password: passwrodHash,
		Gender: req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Email not found"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Token: token,
	}
	return &data, nil
}