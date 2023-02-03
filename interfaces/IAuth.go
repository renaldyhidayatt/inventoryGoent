package interfaces

import (
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/ent"
)

type IAuthRepository interface {
	Register(input request.RegisterRequest) (*ent.User, error)
	Login(input request.AuthRequest) (*ent.User, error)
}

type IAuthService interface {
	Register(input request.RegisterRequest) (*ent.User, error)
	Login(input request.AuthRequest) (*ent.User, error)
}
