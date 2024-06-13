package forms

type ServiceUpdateTDRequest struct {
	ID uint `json:"id"`

	Type string `json:"Type"`

	Description string `json:"Description"`
}

type ServiceUpdateSetTrueFalseState struct {
	ID uint `json:"id"`
}

type SearchServiceName struct {
	Type string `json:"Type"`
}
