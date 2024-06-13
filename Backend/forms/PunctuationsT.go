package forms

type UpdatePunctuationType struct {
	ID uint `json:"id"`

	TimeTravelPoint int `json:"TimeTravelPoint"`

	KindnessPoint int `json:"KindnessPoint"`

	TimeWorkPoint int `json:"TimeWorkPoint"`

	QualityPoint int `json:"QualityPoint"`
}

type PunctuationTRequest struct {
	ID uint `json:"id"`
}
