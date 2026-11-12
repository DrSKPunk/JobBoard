package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/database"
	"github.com/DrSKPunk/JobBoard/internal/models"
)

func GetAllVacancies(ctx context.Context) ([]models.Vakansii, error) {
	rows, err := database.DB.Query(ctx, "SELECT id, title, description, salary, city, published_at::text FROM vacancies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []models.Vakansii
	for rows.Next() {
		var v models.Vakansii
		err := rows.Scan(&v.ID, &v.Title, &v.Description, &v.Salary, &v.City, &v.PublishedAt)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}
	return vacancies, nil
}

func GetVacancyByID(ctx context.Context, id int) (models.Vakansii, error) {
	var v models.Vakansii
	err := database.DB.QueryRow(ctx, "SELECT id, title, description, salary, city, published_at::text FROM vacancies WHERE id = $1", id).
		Scan(&v.ID, &v.Title, &v.Description, &v.Salary, &v.City, &v.PublishedAt)
	return v, err
}

func CreateVacancy(ctx context.Context, v models.Vakansii) (int, error) {
	var companyID int
	err := database.DB.QueryRow(ctx, "SELECT id FROM companies LIMIT 1").Scan(&companyID)
	if err != nil {
		err = database.DB.QueryRow(ctx,
			"INSERT INTO companies (name) VALUES ('Test Company') RETURNING id").Scan(&companyID)
		if err != nil {
			return 0, err
		}
	}

	var id int
	err = database.DB.QueryRow(ctx,
		"INSERT INTO vacancies (company_id, title, description, salary, city, published_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		companyID, v.Title, v.Description, v.Salary, v.City, v.PublishedAt).Scan(&id)
	return id, err
}

func UpdateVacancy(ctx context.Context, id int, v models.Vakansii) error {
	_, err := database.DB.Exec(ctx,
		"UPDATE vacancies SET title = $1, description = $2, salary = $3, city = $4, published_at = $5 WHERE id = $6",
		v.Title, v.Description, v.Salary, v.City, v.PublishedAt, id)
	return err
}

func DeleteVacancy(ctx context.Context, id int) error {
	_, err := database.DB.Exec(ctx, "DELETE FROM vacancies WHERE id = $1", id)
	return err
}
