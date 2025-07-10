package modelbase

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ID = pgtype.UUID

func StringToID(str string) (ID, error) {
	id := pgtype.UUID{}
	err := id.Scan(str)
	return id, err
}

type Identifiable struct {
	ID ID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
}
