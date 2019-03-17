package models

import (
	"github.com/google/uuid"
	"time"
)

//用户
type User struct {
	Base
	ID        uint32 	`gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name 			string
	Username		string		`gorm:"not null;unique"`
	Scope 			string		`gorm:"-"`
	PasswordHash 	string
	salt			string
	Metadata 		[]byte		`gorm:"type:jsonb"`
	RoleId			int
	Role 			*Role		`gorm:"ForeignKey:RoleId"`
}


func (r *User) BeforeCreate() (err error) {
	r.ID = uuid.New().ID()
	return
}
