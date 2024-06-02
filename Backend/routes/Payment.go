package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/forms"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/gorilla/mux"
)

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request){

	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Create(&payment).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create payment, maybe type exists already"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&payment)
}
 
func GetPaymentHandler(w http.ResponseWriter, r *http Request){
	var payment models.Payment

	params := mux.Vars(r)

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.First(&payment, params["id"]).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find payment"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&payment)
}

func UpdateTotalPayment(w http.ResponseWriter, r *http.Request){
	var request forms.UpdateTotalPayment
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.First(&payment, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find payment"})
		return
	}

	if err := transaction.Model(&payment).Select("TotalPayment").Updates(map[string]interface{}{"TotalPayment": request.TotalPayment}).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update payment"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&payment)
}

func SetPaymentMethodNequi(w http.ResponseWriter, r *http.Request){
	var request forms.SetPaymentMethod
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.First(&payment, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find payment"})
		return
	}

	if err := transaction.Model(&payment).Select("Nequi","Transferencia","Efectivo").Updates(map[string]interface{}{
		"Nequi": true,
		"Transferencia": false,
		"Efectivo": false}).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update payment"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&payment)
}

func SetPaymentMethodTransferencia(w http.ResponseWriter, r *http.Request){
	var request forms.SetPaymentMethod
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.First(&payment, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find payment"})
		return
	}

	if err := transaction.Model(&payment).Select("Nequi","Transferencia","Efectivo").Updates(map[string]interface{}{
		"Nequi": false,
		"Transferencia": true,
		"Efectivo": false}).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update payment"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&payment)
}

func SetPaymentMethodEfectivo(w http.ResponseWriter, r *http.Request){
	var request forms.SetNequiPayment
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.First(&payment, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find payment"})
		return
	}

	if err := transaction.Model(&payment).Select("Nequi","Transferencia","Efectivo").Updates(map[string]interface{}{
		"Nequi": false,
		"Transferencia": false,
		"Efectivo": true}).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update payment"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&payment)
}

