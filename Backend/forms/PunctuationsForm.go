package forms

type UpdateGeneralScore struct{
	ID uint `json:"id"`

	GeneralScore int `json:"GeneralScore"`
}

type PunctuationRequest struct{
	ID uint `json:"id"`
}

