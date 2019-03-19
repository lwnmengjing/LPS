package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

//域
type Domain struct {
	Base
	ID        uint32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string
	Service   pq.StringArray `gorm:"type:varchar(64)[]"`
	Metadata  []byte         `gorm:"type:jsonb"`
	ExpiredAt time.Time
}

func (*Domain) TableName() string {
	return "domains"
}

func (d *Domain) BeforeCreate() (err error) {
	d.ID = uuid.New().ID()
	return
}
