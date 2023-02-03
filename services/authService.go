package services

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/repository"
)

type AuthService = interfaces.IAuthService

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) Register(input request.RegisterRequest) (*ent.User, error) {
	var userRequest request.RegisterRequest

	userRequest.Firstname = input.Firstname
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.Register(userRequest)

	return res, err
}

func (s *authService) Login(input request.AuthRequest) (*ent.User, error) {
	var authRequest request.AuthRequest

	authRequest.Email = input.Email
	authRequest.Password = input.Password

	res, err := s.repository.Login(authRequest)

	return res, err
}
