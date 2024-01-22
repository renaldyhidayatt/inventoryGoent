package repository

import (
	"context"
	"errors"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/user"
)

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, ctx context.Context) *authRepository {
	return &authRepository{
		db:      db,
		context: ctx,
	}
}

func (r *authRepository) Register(requests request.RegisterRequest) (*ent.User, error) {
	_, err := r.db.User.Query().Where(user.EmailEQ(requests.Email)).First(r.context)
	if err == nil {
		return nil, errors.New("email already exists")
	} else if !ent.IsNotFound(err) {
		return nil, errors.New("failed to check email existence")
	}

	newUser, err := r.db.User.
		Create().
		SetFirstname(requests.Firstname).
		SetLastname(requests.LastName).
		SetEmail(requests.Email).
		SetPassword(requests.Password).
		Save(r.context)

	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return newUser, nil
}

func (r *authRepository) Login(requests request.AuthRequest) (*ent.User, error) {
	user, err := r.db.User.
		Query().
		Where(user.EmailEQ(requests.Email)).
		Where(user.PasswordEQ(requests.Password)).
		Only(r.context)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
