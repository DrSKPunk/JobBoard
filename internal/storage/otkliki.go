package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/database"
	"github.com/DrSKPunk/JobBoard/internal/models"
)

func GetOtklikiByVacancyID(ctx context.Context, vacID int) ([]models.Otkliki, error) {
	rows, err := database.DB.Query(ctx, "SELECT id_otklika, id_rezume, id_vakansii, data_otklika::text, status FROM otkliki WHERE id_vakansii = $1", vacID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var otkliki []models.Otkliki
	for rows.Next() {
		var o models.Otkliki
		err := rows.Scan(&o.ID, &o.IDRezume, &o.IDVakansii, &o.DataOtk, &o.Status)
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
		"INSERT INTO otkliki (id_rezume, id_vakansii, data_otklika, status) VALUES ($1, $2, $3, $4) RETURNING id_otklika",
		o.IDRezume, o.IDVakansii, o.DataOtk, o.Status).Scan(&id)
	return id, err
}
