package dataup

import "time"

type PatientRecord struct {
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"created_at"`
	PatientID        string          `json:"patient_id"`
	Input            DiseasesPayload `json:"input"`
	EvaluatedWeight  DiseasesWeight  `json:"evaluated_weight"`
	EvaluationResult float64         `json:"evaluation_result"`
}