package services

import (
	"context"
	"github.com/hvmidrezv/web-app/api/dto"
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/data/db"
	"github.com/hvmidrezv/web-app/data/models"
	"github.com/hvmidrezv/web-app/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.UpdateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarTypeService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarTypeService) GetById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarTypeService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
