package models

type Vakansii struct {
	ID       int
	Nazvanie string
	Opisanie string
	Zarplata int
	Gorod    string
	DataPub  string
}

type Otkliki struct {
	ID         int
	IDRezume   int
	IDVakansii int
	DataOtk    string
	Status     string
}
