package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/database"
	"github.com/DrSKPunk/JobBoard/internal/models"
)

func GetOtklikiByVacancyID(ctx context.Context, vacID int) ([]models.Otkliki, error) {
	rows, err := database.DB.Query(ctx, "SELECT id, resume_id, vacancy_id, created_at::text, status FROM applications WHERE vacancy_id = $1", vacID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var otkliki []models.Otkliki
	for rows.Next() {
		var o models.Otkliki
		err := rows.Scan(&o.ID, &o.ResumeID, &o.VacancyID, &o.CreatedAt, &o.Status)
		if err != nil {
			return nil, err
		}
		otkliki = append(otkliki, o)
	}
	return otkliki, nil
}

func CreateOtklik(ctx context.Context, o models.Otkliki) (int, error) {
	var id int
	err := database.DB.QueryRow(ctx,
		"INSERT INTO applications (resume_id, vacancy_id, created_at, status) VALUES ($1, $2, $3, $4) RETURNING id",
		o.ResumeID, o.VacancyID, o.CreatedAt, o.Status).Scan(&id)
	return id, err
}
