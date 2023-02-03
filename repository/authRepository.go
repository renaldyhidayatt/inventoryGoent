package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/ent/user"
	"github.com/renaldyhidayatt/inventorygoent/interfaces"
	"github.com/renaldyhidayatt/inventorygoent/security"
)

type AuthRepository = interfaces.IAuthRepository

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, context context.Context) *authRepository {
	return &authRepository{db: db, context: context}

}

func (r *authRepository) Register(input request.RegisterRequest) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("email already exitst")
	}

	newUser, err := r.db.User.Create().
		SetFirstname(input.Firstname).
		SetLastname(input.LastName).
		SetEmail(input.Email).
		SetPassword(security.HashPassword(input.Password)).
		Save(r.context)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

func (r *authRepository) Login(input request.AuthRequest) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	checkPassword := security.VerifyPassword(user.Password, input.Password)

	if checkPassword != nil {
		return nil, fmt.Errorf("failed checkhash password: %w", err)

	}

	return user, nil
}
