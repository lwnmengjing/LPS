package models

import (
	"github.com/google/uuid"
	"time"
)

//角色
type Role struct {
	Base
	ID        uint32	`gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name 			string
	Alias			string
	Metadata 		[]byte			`gorm:"type:jsonb"`
}

func (r *Role) BeforeCreate() (err error) {
	r.ID = uuid.New().ID()
	return
}
