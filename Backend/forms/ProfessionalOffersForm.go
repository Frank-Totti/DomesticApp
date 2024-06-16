package forms

type UpdatePoRequest struct {
	Professional_id uint `json:"Professional_id"`

	Major string `json:"Major"`

	RelationalExperience string `json:"RExperience"`

	RelationalExperiencePhoto string `json:"REPhoto"` // []byte

	MajorPhoto string `json:"MPhoto"` // []byte

	UnitPrice float64 `json:"UnitPrice"`

	PricePerHour float64 `json:"PricePerHour"`
}
