package models

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

type Vakansii struct {
	ID       int    `json:"id_vakansii"`
	Nazvanie string `json:"nazvanie"`
	Opisanie string `json:"opisanie"`
	Zarplata int    `json:"zarplata"`
	Gorod    string `json:"gorod"`
	DataPub  string `json:"data_publikacii"`
}

type Otkliki struct {
	ID         int    `json:"id_otklika"`
	IDRezume   int    `json:"id_rezume"`
	IDVakansii int    `json:"id_vakansii"`
	DataOtk    string `json:"data_otklika"`
	Status     string `json:"status"`
}
