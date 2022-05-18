package domain

import "context"

type Repository interface {
	GetUsers(ctx context.Context) []User
}