package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ID = pgtype.UUID

var NilID = &ID{Valid: false}

func StringToID(str string) (ID, error) {
	id := pgtype.UUID{}
	err := id.Scan(str)
	return id, err
}

func StringsToIDs(strs []string) ([]ID, error) {
	ids := make([]ID, len(strs))
	for i, str := range strs {
		id, err := StringToID(str)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
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
