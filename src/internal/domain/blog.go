package domain

import "time"

// book のドメインモデル(ドメイン貧血症)
type Book struct {
	Id       int64
	Title    string
	Author   string
	IssuedAt time.Time
}
