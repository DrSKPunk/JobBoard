package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VacancyStorage struct {
	db *pgxpool.Pool
}

func NewVacancyStorage(db *pgxpool.Pool) *VacancyStorage {
	return &VacancyStorage{db: db}
}

func (r *VacancyStorage) GetAllVacancies(ctx context.Context) ([]models.Vacancy, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, salary, city, published_at::text FROM vacancies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []models.Vacancy
	for rows.Next() {
		var v models.Vacancy
		err := rows.Scan(&v.ID, &v.Title, &v.Description, &v.Salary, &v.City, &v.PublishedAt)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}
	return vacancies, nil
}

func (r *VacancyStorage) GetVacancyByID(ctx context.Context, id int) (models.Vacancy, error) {
	var v models.Vacancy
	err := r.db.QueryRow(ctx, "SELECT id, title, description, salary, city, published_at::text FROM vacancies WHERE id = $1", id).
		Scan(&v.ID, &v.Title, &v.Description, &v.Salary, &v.City, &v.PublishedAt)
	return v, err
}

func (r *VacancyStorage) CreateVacancy(ctx context.Context, v models.Vacancy) (int, error) {
	var companyID int
	err := r.db.QueryRow(ctx, "SELECT id FROM companies LIMIT 1").Scan(&companyID)
	if err != nil {
		err = r.db.QueryRow(ctx,
			"INSERT INTO companies (name) VALUES ('Test Company') RETURNING id").Scan(&companyID)
		if err != nil {
			return 0, err
		}
	}

	var id int
	err = r.db.QueryRow(ctx,
		"INSERT INTO vacancies (company_id, title, description, salary, city, published_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		companyID, v.Title, v.Description, v.Salary, v.City, v.PublishedAt).Scan(&id)
	return id, err
}

func (r *VacancyStorage) UpdateVacancy(ctx context.Context, id int, v models.Vacancy) error {
	_, err := r.db.Exec(ctx,
		"UPDATE vacancies SET title = $1, description = $2, salary = $3, city = $4, published_at = $5 WHERE id = $6",
		v.Title, v.Description, v.Salary, v.City, v.PublishedAt, id)
	return err
}

func (r *VacancyStorage) DeleteVacancy(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM vacancies WHERE id = $1", id)
	return err
}
