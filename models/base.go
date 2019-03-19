package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/micro/go-log"
)

var (
	DB *gorm.DB
)

type (
	Event string
)

const (
	EventBefore Event = "before"
	EventAfter        = "after"
)

func DBInit() {
	var err error
	DB, err = gorm.Open("postgres", "postgresql://root@linwenxiangdeMacBook-Air.local:26257/m_sso?sslmode=disable")
	//DB, err = gorm.Open("mysql", "mysql://test:$j1mepaag^@101.132.122.150:3306/m_lps")
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&Domain{}, &Role{}, &User{})
}

type Base struct {
	action                  Action
	eventBefore, eventAfter func(baser Baser) (Baser, error)
}

func (b *Base) SetEventFunc(event Event, f func(baser Baser) (Baser, error)) {
	switch event {
	case EventBefore:
		b.eventBefore = f
	case EventAfter:
		b.eventAfter = f
	}
}

func (b *Base) GetEventFunc(event Event) (f func(baser Baser) (Baser, error)) {
	switch event {
	case EventBefore:
		f = b.eventBefore
	case EventAfter:
		f = b.eventAfter
	}
	return f
}

func (b Base) SetAction(v Action) {
	b.action = v
}

func (b Base) GetAction() Action {
	return b.action
}

func Create(base Baser, db *gorm.DB) (Baser, error) {
	base.SetAction(ActionCreate)
	b, _ := eventFunc(base, base.GetEventFunc(EventBefore), nil)

	err := db.Create(base).Error

	return eventFunc(b, base.GetEventFunc(EventAfter), err)
}

func Update(base Baser, db *gorm.DB, cols ...string) (Baser, error) {
	base.SetAction(ActionUpdate)
	base, _ = eventFunc(base, base.GetEventFunc(EventBefore), nil)

	var err error
	query := db.Model(base)
	if len(cols) == 0 {
		query = query.Save(base)
	} else {
		for _, col := range cols {
			query = query.Update(col)
		}
	}
	if err = query.Error; err != nil {
		return base, err
	}
	return eventFunc(base, base.GetEventFunc(EventAfter), err)
}

func View(base Baser, db *gorm.DB, condition Condition) (Baser, error) {
	base.SetAction(ActionView)
	base, _ = eventFunc(base, base.GetEventFunc(EventBefore), nil)

	query := db.Model(base)
	err := findCondition(query, condition).First(base).Error

	return eventFunc(base, base.GetEventFunc(EventAfter), err)
}

func List(base Baser, db *gorm.DB, list interface{}, condition Condition) error {
	base.SetAction(ActionList)
	query := db.Table(base.TableName())
	if condition.Limit == 0 {
		condition.Limit = LimitDefault
	}
	findCondition(query, condition).Find(list)
	return nil

}

func Delete(base Baser, db *gorm.DB, condition map[string][]interface{}, force bool) (Baser, error) {
	base.SetAction(ActionDelete)
	base, _ = eventFunc(base, base.GetEventFunc(EventBefore), nil)
	var err error
	if force {
		//硬删除
		db = db.Unscoped()
	}
	for k, v := range condition {
		db = db.Where(k, v...)
	}
	err = db.Delete(base).Error
	return eventFunc(base, base.GetEventFunc(EventAfter), err)
}

func eventFunc(base Baser, f func(base Baser) (Baser, error), err error) (Baser, error) {
	if f != nil && err == nil {
		return f(base)
	}
	return base, err
}

//组装查询条件
func findCondition(query *gorm.DB, condition Condition) *gorm.DB {
	for k, v := range condition.Where {
		query = query.Where(k, v...)
	}

	for k, v := range condition.Not {
		query = query.Not(k, v...)
	}

	for k, v := range condition.Or {
		query = query.Or(k, v...)
	}

	for k, v := range condition.Having {
		query = query.Having(k, v...)
	}

	if condition.Group != "" {
		query = query.Group(condition.Group)
	}

	if condition.Assign != nil {
		query = query.Assign(condition.Assign)
	}

	for _, o := range condition.Order {
		query = query.Order(o)
	}

	if condition.Attrs != nil {
		query = query.Attrs(condition.Attrs)
	}

	if condition.Limit != 0 {
		query = query.Limit(condition.Limit)
	}

	if condition.Offset > 0 {
		query = query.Offset(condition.Offset)
	}

	return query
}
