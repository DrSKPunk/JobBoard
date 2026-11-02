package storage

import "github.com/DrSKPunk/JobBoard/internal/models"

var Vakansii = []models.Vakansii{
	{ID: 1, Nazvanie: "Программист", Opisanie: "Нужен программист для программирования", Zarplata: 150000, Gorod: "Барнаул", DataPub: "2026-10-28"},
	{ID: 2, Nazvanie: "Тракторист", Opisanie: "Нужно чистить дорогу", Zarplata: 100000, Gorod: "Сургут", DataPub: "2026-10-27"},
}

var Otkliki = []models.Otkliki{}

var NextVakansiiID = 3
var NextOtklikID = 1
