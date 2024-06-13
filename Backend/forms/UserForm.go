package forms

import (
	"time"

	"github.com/Frank-totti/DomesticApp/models"
)

type UserResponseGetRequest struct {
	Payment_id             uint      `json:"Payment_id"`
	Bill_id                uint      `json:"Bill_id"`
	Request_id             uint      `json:"Request_id"`
	Duser_id               uint      `json:"Duser_id"`
	Payment_totalPayment   int       `json:"Payment_totalPayment"`
	Payment_transferencia  bool      `json:"Payment_transferencia"`
	Payment_efectivo       bool      `json:"Payment_efectivo"`
	Payment_nequi          bool      `json:"Payment_nequi"`
	Bill_init_work_hour    time.Time `json:"Bill_init_work_hour"`
	Bill_final_work_hour   time.Time `json:"Bill_final_work_hour"`
	Bill_final_travel_hour time.Time `json:"Bill_final_travel_hour"`
	Bill_discounts_applied int       `json:"Bill_discounts_applied"`
	Bill_partial_payment   int       `json:"Bill_partial_payment"`
	Request_travel_hour    time.Time `json:"Request_travel_hour"`
	Request_state          string    `json:"Request_state"`
}

type UserWriterHistory struct {
	UserID         uint             `json:"id"`
	Total          int              `json:"total"`
<<<<<<< HEAD
<<<<<<< HEAD
	RequestHistory []models.Payment `json:"history"`
=======
	RequestHistory []models.Request `json:"history"`
>>>>>>> 76553de (repair of users/request route)
=======
	RequestHistory []models.Payment `json:"history"`
>>>>>>> 6115b9b (Creation of search email users function)
}

type UserRequestHistory struct {
	ID uint `json:"id"`
}

type UserDelete struct {
	ID uint `json:"id"`
}

type UserRequest struct {
	ID         uint      `json:"id"`
	UpdateUser InnerUser `json:"UpdateUser"`
}

type InnerUser struct {
	Person PersonForm `json:"Person"`
	//NewEmail      string     `json:"NewEmail"`
	PublicService []byte `json:"PublicService"`
}
