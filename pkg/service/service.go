package service

import "dashboard/pkg/repository"

type Authorization interface {
}

type Invoice interface {
}

type Service struct {
	Authorization
	Invoice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
