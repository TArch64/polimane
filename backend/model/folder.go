package model

type Folder struct {
	*Identifiable
	*Timestamps
	Name string `json:"name"`

	// Relations
	Schemas []*Schema `gorm:"many2many:folder_schemas" json:"-"`
}
