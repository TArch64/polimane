package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ID = pgtype.UUID

func StringToID(str string) (ID, error) {
	id := pgtype.UUID{}
	err := id.Scan(str)
	return id, err
}

func MustStringToID(str string) ID {
	id, err := StringToID(str)
	if err != nil {
		panic(err)
	}
	return id
}

type Identifiable struct {
	ID ID `gorm:"primaryKey;default:gen_random_uuid()" json:"id"`
}
