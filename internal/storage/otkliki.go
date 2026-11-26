package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OtklikiStorage struct {
	db *pgxpool.Pool
}

func NewOtklikiStorage(db *pgxpool.Pool) *OtklikiStorage {
	return &OtklikiStorage{db: db}
}

func (r *OtklikiStorage) GetOtklikiByVacancyID(ctx context.Context, vacID int) ([]models.Application, error) {
	rows, err := r.db.Query(ctx, "SELECT id, resume_id, vacancy_id, created_at::text, status FROM applications WHERE vacancy_id = $1", vacID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var otkliki []models.Application
	for rows.Next() {
		var o models.Application
		err := rows.Scan(&o.ID, &o.ResumeID, &o.VacancyID, &o.CreatedAt, &o.Status)
		if err != nil {
			return nil, err
		}
		otkliki = append(otkliki, o)
	}
	return otkliki, nil
}

func (r *OtklikiStorage) CreateOtklik(ctx context.Context, o models.Application) (int, error) {
	var id int
	err := r.db.QueryRow(ctx,
		"INSERT INTO applications (resume_id, vacancy_id, created_at, status) VALUES ($1, $2, $3, $4) RETURNING id",
		o.ResumeID, o.VacancyID, o.CreatedAt, o.Status).Scan(&id)
	return id, err
}
