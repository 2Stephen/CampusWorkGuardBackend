package dto

type StudentAuthParams struct {
	ID     string `json:"studentId"`
	School string `json:"school"`
	Vcode  string `json:"vCode"`
	Email  string `json:"email"`
}
