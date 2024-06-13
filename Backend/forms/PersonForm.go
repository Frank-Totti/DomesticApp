package forms

type PersonForm struct {
	Address  string `json:"Address"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	TNumber  string `json:"TNumber"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type SearchEmail struct {
	Email string `json:"Email"`
}

type SearchName struct {
	Name string `json:"Name"`
}

type SearchLastName struct {
	LastName string `json:"LastName"`
}
