package repository

type Authorization interface {
}

type Invoice interface {
}

type Repository struct {
	Authorization
	Invoice
}

func NewRepository() *Repository {
	return &Repository{}
}
