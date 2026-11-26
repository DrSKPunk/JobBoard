package services

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/DrSKPunk/JobBoard/internal/storage"
)

type VacancyService struct {
	storage *storage.VacancyStorage
}

func NewVacancyService(s *storage.VacancyStorage) *VacancyService {
	return &VacancyService{storage: s}
}

func (s *VacancyService) GetAll(ctx context.Context) ([]dto.VacancyResponse, error) {
	vacancies, err := s.storage.GetAllVacancies(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.VacancyResponse
	for _, v := range vacancies {
		res = append(res, dto.VacancyResponse{
			ID: v.ID, Title: v.Title, Description: v.Description,
			Salary: v.Salary, City: v.City, PublishedAt: v.PublishedAt,
		})
	}
	return res, nil
}

func (s *VacancyService) GetByID(ctx context.Context, id int) (dto.VacancyResponse, error) {
	v, err := s.storage.GetVacancyByID(ctx, id)
	if err != nil {
		return dto.VacancyResponse{}, err
	}
	return dto.VacancyResponse{
		ID: v.ID, Title: v.Title, Description: v.Description,
		Salary: v.Salary, City: v.City, PublishedAt: v.PublishedAt,
	}, nil
}

func (s *VacancyService) Create(ctx context.Context, req dto.VacancyRequest) (int, error) {
	v := models.Vacancy{
		Title: req.Title, Description: req.Description,
		Salary: req.Salary, City: req.City, PublishedAt: req.PublishedAt,
	}
	return s.storage.CreateVacancy(ctx, v)
}

func (s *VacancyService) Update(ctx context.Context, id int, req dto.VacancyRequest) error {
	v := models.Vacancy{
		Title: req.Title, Description: req.Description,
		Salary: req.Salary, City: req.City, PublishedAt: req.PublishedAt,
	}
	return s.storage.UpdateVacancy(ctx, id, v)
}

func (s *VacancyService) Delete(ctx context.Context, id int) error {
	return s.storage.DeleteVacancy(ctx, id)
}
