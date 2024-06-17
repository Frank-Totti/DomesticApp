package forms

import (
	"time"

	"github.com/Frank-totti/DomesticApp/models"
)

type GetGeneralPunctuationPro struct {
	Name      string
	Last_name string
}

type Professional_id struct {
	ID uint `json:"id"`
}

type ProfesionalForm struct {
	ID           uint             `json:"id"`
	Professional InnerProfesional `json:"Professional"`
}

type InnerProfesional struct {
	Person           PersonForm `json:"Person"`
	ProfilePicture   string     `json:"ProfilePicture"`
	Birth            time.Time  `json:"Birth"`
	IdentifyDocument string     `json:"IdentifyDocument"`
	PhotoDocument    string     `json:"PhotoDocument"`
}

type ParticularPunctuation struct {
	Punctuation string `json:"punctuation"`
}

type ProfessionalWriterHistory struct {
	ProfessionalID uint             `json:"id"`
	Total          int              `json:"total"`
	RequestHistory []models.Payment `json:"history"`
}

type ProfessionalRequestHistory struct {
	ID    uint `json:"id"`
	State bool `json:"State"`
}
