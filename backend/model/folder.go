package model

type Folder struct {
	*Identifiable
	*Timestamps
	Name   string `json:"name"`
	UserID ID     `json:"userId"`

	// Relations
	User    *User     `json:"-"`
	Schemas []*Schema `gorm:"many2many:folder_schemas" json:"-"`
}
