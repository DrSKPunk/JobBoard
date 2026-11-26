package services

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/DrSKPunk/JobBoard/internal/storage"
)

type ApplicationService struct {
	storage *storage.OtklikiStorage
}

func NewApplicationService(s *storage.OtklikiStorage) *ApplicationService {
	return &ApplicationService{storage: s}
}

func (s *ApplicationService) GetByVacancyID(ctx context.Context, vacID int) ([]dto.ApplicationResponse, error) {
	apps, err := s.storage.GetOtklikiByVacancyID(ctx, vacID)
	if err != nil {
		return nil, err
	}

	var res []dto.ApplicationResponse
	for _, a := range apps {
		res = append(res, dto.ApplicationResponse{
			ID: a.ID, ResumeID: a.ResumeID, VacancyID: a.VacancyID,
			CreatedAt: a.CreatedAt, Status: a.Status,
		})
	}
	return res, nil
}

func (s *ApplicationService) Create(ctx context.Context, vacID int, req dto.ApplicationRequest) (int, error) {
	o := models.Application{
		ResumeID: req.ResumeID, VacancyID: vacID,
	}
	return s.storage.CreateOtklik(ctx, o)
}
