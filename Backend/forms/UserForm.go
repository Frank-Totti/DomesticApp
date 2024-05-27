package forms

type UserRequest struct {
	ID         uint      `json:"id"`
	UpdateUser InnerUser `json:"UpdateUser"`
}

type InnerUser struct {
	Person        PersonForm `json:"Person"`
	NewEmail      string     `json:"NewEmail"`
	PublicService []byte     `json:"PublicService"`
}

type PersonForm struct {
	Address  string `json:"Address"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	TNumber  string `json:"TNumber"`
}
