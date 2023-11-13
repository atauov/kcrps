package service

import (
	"dashboard"
	"dashboard/pkg/repository"
)

type Authorization interface {
	CreateUser(user dashboard.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Invoice interface {
}

type Service struct {
	Authorization
	Invoice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
