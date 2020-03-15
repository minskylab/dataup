package dataup


type PatientProfile struct {
	ID      string          `json:"id"`
	DNI     string          `json:"dni"`
	Phone   string          `json:"phone"`
	Email   string          `json:"email"`
	Records []PatientRecord `json:"records"`
}
