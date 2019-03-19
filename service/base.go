package service

import (
	"errors"
	"github.com/lwnmengjing/LPS/models"
)

type Base struct {
	Model                   models.Baser
	EventBefore, EventAfter func(base models.Baser) (models.Baser, error)
}

func (b *Base) eventFunc() {
	b.Model.SetEventFunc(models.EventBefore, b.EventBefore)
}

func (b *Base) Create() (err error) {
	if b.Model == nil {
		return errors.New("db model is must not empty")
	}
	b.Model, err = models.Create(b.Model, models.DB)
	return err
}

func (b *Base) Update(cols ...string) (err error) {
	if b.Model == nil {
		return errors.New("db model is must not empty")
	}
	b.Model, err = models.Update(b.Model, models.DB, cols...)
	return err
}

func (b *Base) View(condition models.Condition) (err error) {
	if b.Model == nil {
		return errors.New("db model is must not empty")
	}
	b.Model, err = models.View(b.Model, models.DB, condition)
	return err
}

func (b *Base) List(list interface{}, condition models.Condition) error {
	if b.Model == nil {
		return errors.New("db model is must not empty")
	}
	return models.List(b.Model, models.DB, list, condition)
}

func (b *Base) Delete(condition map[string][]interface{}, force bool) (err error) {
	if b.Model == nil {
		return errors.New("db model is must not empty")
	}
	b.Model, err = models.Delete(b.Model, models.DB, condition, force)
	return err
}

func (b *Base) putSelf() {

}
