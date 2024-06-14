package forms

type CreatePoRequest struct {
	Service_id uint `json:"ServiceID"`

	Professional_id uint `json:"ProfessionalID"`

	Major string `json:"Major"`

	RelationalExperiencePhoto []byte `json:"REPhoto"`

	MajorPhoto []byte `json:"MPhoto"`

	UnitPrice float64 `json:"UnitPrice"`

	PricePerHour float64 `json:"PricePerHour"`
}
