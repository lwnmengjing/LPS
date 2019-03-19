package models

import "github.com/jinzhu/gorm"

type Baser interface {
	//Create(base *Baser) error
	//Update(cols... string) error
	//View(condition Condition) error
	//Delete(force bool) error
	SetEventFunc(event Event, f func(baser Baser) (Baser, error))
	GetEventFunc(event Event) func(baser Baser) (Baser, error)
	SetAction(v Action)
	GetAction() Action
	TableName() string
}

type (
	Status  int
	ModelId string
	Action  string
)

const (
	//list默认个数
	LimitDefault = 10
	//状态
	StatusEnable  Status = 10001
	StatusDisable        = 10000
	StatusShow           = 10002
	StatusYes            = 10011
	StatusNo             = 10010
	//模型ID
	DomainID ModelId = "domain"
	RoleID           = "role"
	UserID           = "user"
	//操作场景
	ActionCreate Action = "create"
	ActionUpdate        = "update"
	ActionView          = "view"
	ActionList          = "list"
	ActionDelete        = "delete"
)

type Condition struct {
	Where   map[string][]interface{}
	Not     map[string][]interface{}
	Or      map[string][]interface{}
	Attrs   Baser
	Assign  Baser
	Order   []string
	Limit   int
	Offset  int
	Group   string
	Having  map[string][]interface{}
	Preload map[string]func(db *gorm.DB) *gorm.DB
}
