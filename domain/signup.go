package domain

import (
	"context"
)

type SignupRequest struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required""`
	UserName  string `form:"last_name"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
