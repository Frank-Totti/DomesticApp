package forms

type UpdateBill struct{
	ID uint `json:"id"`

	InitWorkHour time.Time `json:"InitWorkHour"`

	FinalWorkHour time.Time `json:"FinalWorkHour"`

	FinalTravelHour time.Time `json:"FinalTravelHour"`

	DiscountsApplied float64 `json:"DiscountsApplied"`

	PartialPayment float64 `json:"PartialPayment"`
}