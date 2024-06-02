package forms

type UpdateTotalPayment struct{
	ID uint `json:"id"`

	TotalPayment float64 `json:"TotalPayment"`
}

type SetPaymentMethod struct{
	ID uint `json:"id"`
}
