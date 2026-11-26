package dto

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type VacancyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      int    `json:"salary"`
	City        string `json:"city"`
	PublishedAt string `json:"published_at"`
}

type VacancyResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      int    `json:"salary"`
	City        string `json:"city"`
	PublishedAt string `json:"published_at"`
}

type ApplicationRequest struct {
	ResumeID int `json:"resume_id"`
}

type ApplicationResponse struct {
	ID        int    `json:"id"`
	ResumeID  int    `json:"resume_id"`
	VacancyID int    `json:"vacancy_id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}
