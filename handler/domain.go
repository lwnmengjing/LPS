package handler

import (
	"context"
	"encoding/json"
	"github.com/lwnmengjing/LPS/models"
	proto "github.com/lwnmengjing/LPS/proto/domain"
	"github.com/lwnmengjing/LPS/service"
	"github.com/micro/go-micro/errors"
	"net/http"
	"time"
)

const (
	IdName = "go.micro.srv.LPS.domain"
)

type Domain struct {}

func (d *Domain) Create(ctx context.Context, req *proto.CreateModel, rsp *proto.Empty) error {
	id := IdName + ".create"
	var metadata []byte
	var err error
	if req.Metadata == nil {
		metadata = []byte(`[]`)
	} else {
		if metadata, err = json.Marshal(req.Metadata); err != nil {
			return errors.New(id, err.Error(), http.StatusBadRequest)
		}
	}

	expiredAt, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
	if err != nil {
		return errors.New(id, err.Error(), http.StatusBadRequest)
	}
	base := &service.Base{
		Model: &models.Domain{
			Name:      req.Name,
			Metadata:  metadata,
			Service:   req.Service,
			ExpiredAt: expiredAt,
		},
	}
	if err := base.Create(); err != nil {
		return errors.New(id, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (d *Domain) Update(ctx context.Context, req *proto.Model, rsp *proto.Empty) error {
	id := IdName + ".update"
	model := &models.Domain{
		ID: req.Id,
	}
	base := &service.Base{
		Model: model,
	}
	err := base.View(models.Condition{})
	if err != nil {
		return errors.New(id, err.Error(), http.StatusNotFound)
	}
	if req.Name != "" {
		model.Name = req.Name
	}
	if len(req.Metadata) > 0 {
		model.Metadata, _  = json.Marshal(req.Metadata)
	}
	if len(req.Service) > 0 {
		model.Service = req.Service
	}
	if err = base.Update(); err != nil {
		return errors.New(id, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (d *Domain) View(ctx context.Context, req *proto.ModelId, rsp *proto.Model) error {
	id := IdName + ".view"
	object := &models.Domain{
		ID: req.Id,
	}
	base := &service.Base{
		Model: object,
	}
	err := base.View(models.Condition{})
	if err != nil {
		return errors.New(id, err.Error(), http.StatusNotFound)
	}
	rsp.Id = req.Id
	rsp.Name = object.Name
	if err = json.Unmarshal(object.Metadata, &rsp.Metadata); err != nil {
		return errors.New(id, "metadata " + err.Error(), http.StatusInternalServerError)
	}
	rsp.ExpiredAt = object.ExpiredAt.String()
	return nil
}


func (d *Domain) List(ctx context.Context, req *proto.Condition, rsp *proto.Models) error {
	id := IdName + ".list"
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	domain := &models.Domain{}
	list, totalItem, err := domain.List(models.DB, models.Condition{
		Limit: int(req.PageSize),
		Offset: int(req.Offset),
	}, nil)
	if err != nil {
		return errors.New(id, err.Error(), http.StatusInternalServerError)
	}
	for _, o := range list {
		model := &proto.Model{
			Id: o.ID,
			Name: o.Name,
			Service: o.Service,
			ExpiredAt: o.ExpiredAt.String(),
		}
		_ = json.Unmarshal(o.Metadata, &model.Metadata)
		rsp.Data = append(rsp.Data, model)
	}
	pages := int64(totalItem / req.PageSize)
	if totalItem % req.PageSize > 0 {
		pages++
	}
	rsp.Page = &proto.Page{
		PageSize: req.PageSize,
		TotalItem: totalItem,
		Page: req.Offset,
		Pages: pages,
	}
	return nil
}

func (d *Domain) Delete(ctx context.Context, req *proto.Condition, rsp *proto.Empty) error {
	id := IdName + ".delete"
	base := &service.Base{
		Model: &models.Domain{
			ID: req.Id,
		},
	}
	err := base.Delete(nil, false)
	if err != nil {
		return errors.New(id, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
