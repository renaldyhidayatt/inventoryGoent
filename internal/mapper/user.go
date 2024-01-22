package mapper

import (
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent"
)

type userMapper struct {
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserResponse(user *ent.User) *response.User {
	return &response.User{
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (m *userMapper) ToUserResponses(users []*ent.User) *[]response.User {
	var userResponses []response.User

	for _, user := range users {
		userResponse := m.ToUserResponse(user)
		userResponses = append(userResponses, *userResponse)
	}

	return &userResponses
}
