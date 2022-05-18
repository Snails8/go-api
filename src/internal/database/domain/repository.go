package domain

type Repository interface {
	GetUsers() []User
}