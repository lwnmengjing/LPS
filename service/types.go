package service

import "github.com/lwnmengjing/LPS/models"

type Service interface {
	Create() error
	Update(cols... string) error
	View(condition models.Condition) error
	Delete(force bool) error
	List(condition models.Condition) ([]models.Baser, error)
}
