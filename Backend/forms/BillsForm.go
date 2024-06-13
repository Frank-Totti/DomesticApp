package forms

type UpdateBill struct {
	ID uint `json:"id"`

	DiscountsApplied float64 `json:"DiscountsApplied"`

	PartialPayment float64 `json:"PartialPayment"`
}
