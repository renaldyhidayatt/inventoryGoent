package service

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/pkg/auth"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/hash"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"go.uber.org/zap"
)

type authService struct {
	hash           hash.Hashing
	token          auth.TokenManager
	authRepository repository.AuthRepository
	logger         logger.Logger
}

func NewAuthService(hash hash.Hashing, token auth.TokenManager, authRepository repository.AuthRepository, logger logger.Logger) *authService {
	return &authService{
		hash:           hash,
		token:          token,
		authRepository: authRepository,
		logger:         logger,
	}
}

func (s *authService) Register(requests request.RegisterRequest) (*ent.User, error) {
	var registerRequest request.RegisterRequest

	hashing, err := s.hash.HashPassword(registerRequest.Password)

	if err != nil {
		s.logger.Error("failed to hash password:", zap.Error(err))
		return nil, err
	}

	registerRequest.Firstname = requests.Firstname
	registerRequest.LastName = requests.LastName
	registerRequest.Email = requests.Email
	registerRequest.Password = hashing

	user, err := s.authRepository.Register(requests)

	if err != nil {
		s.logger.Error("failed to register user:", zap.Error(err))
		return nil, err
	}

	return user, nil

}

func (s *authService) Login(request request.AuthRequest) (string, error) {
	user, err := s.authRepository.Login(request)

	if err != nil {
		return "", err
	}

	err = s.hash.ComparePassword(user.Password, request.Password)

	if err != nil {
		s.logger.Error("failed to compare password:", zap.Error(err))

		return "", err
	}

	jwtToken, err := s.createJwt(user.ID)

	if err != nil {
		s.logger.Error("failed to create jwt token:", zap.Error(err))
		return "", err
	}

	return jwtToken, nil
}

func (s *authService) createJwt(id int) (string, error) {
	token, err := s.token.NewJwtToken(id)

	if err != nil {
		s.logger.Error("failed to create jwt token:", zap.Error(err))

		return "", err
	}

	return token, nil
}
