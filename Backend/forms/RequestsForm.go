package forms

type UpdateTravelHour struct{
	ID uint `json:"id"`

	TravelHour time.Time `json:"TravelHour"`
}

type SetRequestStateTrueFalse struct{
	ID uint `json:"id"`
}