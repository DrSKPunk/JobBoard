package models

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

type Vakansii struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      int    `json:"salary"`
	City        string `json:"city"`
	PublishedAt string `json:"published_at"`
}

type Otkliki struct {
	ID        int    `json:"id"`
	ResumeID  int    `json:"resume_id"`
	VacancyID int    `json:"vacancy_id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}
