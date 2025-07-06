package modelbase

import "strconv"

type ID uint

func (i ID) Model() *Identifiable {
	return &Identifiable{ID: i}
}

func (i ID) String() string {
	return strconv.Itoa(int(i))
}

type Identifiable struct {
	ID ID `gorm:"type:serial;primaryKey" json:"id"`
}
