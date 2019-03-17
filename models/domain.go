package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

//域
type Domain struct {
	Base
	ID        uint32	`gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name 			string
	Service 		pq.StringArray	`gorm:"type:varchar(64)[]"`
	Metadata 		[]byte 			`gorm:"type:jsonb"`
	ExpiredAt		time.Time
}

func (d *Domain) BeforeCreate() (err error) {
	d.ID = uuid.New().ID()
	return
}

func (d *Domain) List(db *gorm.DB, condition Condition, after func(d *Domain) (*Domain, error)) ([]*Domain, int64, error) {
	var list []*Domain
	d.SetAction(ActionList)
	query := db.Model(d)
	var totalItem int64
	if err := query.Count(&totalItem).Error; err != nil {
		return nil, 0, err
	}
	if condition.Limit == 0 {
		condition.Limit = LimitDefault
	}
	err := findCondition(query, condition).Find(&list).Error
	if f := d.GetEventFunc(EventAfter); f != nil && err == nil {
		for _, object := range list {
			 object, _ = after(d)
			 _ = object.Service
		}
	}
	return list, totalItem, err

}
//func (b *Domain) SetAction(v Action) {
//	b.action = v
//}
//
//func (b *Domain) GetAction() Action {
//	return b.action
//}
//
//func (b *Domain) BeforeFunc() func(base *Baser) (*Baser, error) {
//	return b.Before
//}
//
//func (b *Domain) AfterFunc() func(base *Baser) (*Baser, error) {
//	return b.Before
//}
