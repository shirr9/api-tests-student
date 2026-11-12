package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"` // #nosec G117 -- Поле запроса, значение берётся из окружения.
}
