package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/database"
	"github.com/DrSKPunk/JobBoard/internal/models"
)

func GetAllVacancies(ctx context.Context) ([]models.Vakansii, error) {
	rows, err := database.DB.Query(ctx, "SELECT id_vakansii, nazvanie, opisanie, zarplata, gorod, data_publikacii::text FROM vakansii")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []models.Vakansii
	for rows.Next() {
		var v models.Vakansii
		err := rows.Scan(&v.ID, &v.Nazvanie, &v.Opisanie, &v.Zarplata, &v.Gorod, &v.DataPub)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}
	return vacancies, nil
}

func GetVacancyByID(ctx context.Context, id int) (models.Vakansii, error) {
	var v models.Vakansii
	err := database.DB.QueryRow(ctx, "SELECT id_vakansii, nazvanie, opisanie, zarplata, gorod, data_publikacii::text FROM vakansii WHERE id_vakansii = $1", id).
		Scan(&v.ID, &v.Nazvanie, &v.Opisanie, &v.Zarplata, &v.Gorod, &v.DataPub)
	return v, err
}

func CreateVacancy(ctx context.Context, v models.Vakansii) (int, error) {
	var companyID int
	err := database.DB.QueryRow(ctx, "SELECT id_kompanii FROM kompanii LIMIT 1").Scan(&companyID)
	if err != nil {
		err = database.DB.QueryRow(ctx,
			"INSERT INTO kompanii (nazvanie, opisanie, gorod, telefon, email) VALUES ('Тестовая компания', 'Описание', 'Москва', '+70000000000', 'test@test.ru') RETURNING id_kompanii").Scan(&companyID)
		if err != nil {
			return 0, err
		}
	}
	var id int
	err = database.DB.QueryRow(ctx,
		"INSERT INTO vakansii (id_kompanii, nazvanie, opisanie, zarplata, gorod, data_publikacii) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_vakansii",
		companyID, v.Nazvanie, v.Opisanie, v.Zarplata, v.Gorod, v.DataPub).Scan(&id)

	return id, err
}

func UpdateVacancy(ctx context.Context, id int, v models.Vakansii) error {
	_, err := database.DB.Exec(ctx,
		"UPDATE vakansii SET nazvanie = $1, opisanie = $2, zarplata = $3, gorod = $4, data_publikacii = $5 WHERE id_vakansii = $6",
		v.Nazvanie, v.Opisanie, v.Zarplata, v.Gorod, v.DataPub, id)
	return err
}

func DeleteVacancy(ctx context.Context, id int) error {
	_, err := database.DB.Exec(ctx, "DELETE FROM vakansii WHERE id_vakansii = $1", id)
	return err
}
