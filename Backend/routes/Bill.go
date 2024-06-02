package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/forms"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/gorilla/mux"
)

func CreateBill(w http.ResponseWriter, r *http.Request){
	var bill models.Bill

	err := json.NewDecoder(r.Body).Decode(&bill)

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

	if err := transaction.Create(&bill).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create bill, maybe type exists already"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&bill)
}

func GetBillsHandler(w http.ResponseWriter, r *http.Request){
	var bills []models.Bill

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Find(&bills).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find bill data"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&bills)
}

func GetBillHandler(w http.ResponseWriter, r *http.Request){
	var bill models.Bill

	params := mux.Vars(r)

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.First(&bill, params["id"]).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find bill"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&bill)
}

func UpdateBill(w http.ResponseWriter, r *http.Request){
	
	var requestBill forms.UpdateBill
	var bill models.Bill

	err := json.NewDecoder(r.Body).Decode(&requestBill)

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

	if err := transaction.First(&bill, requestBill.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find bill"})
		return
	}

	if err := transaction.Model(&bill).Updates(map[string]interface{}{
		"InitWorkHour": request.InitWorkHour, 
		"FinalWorkHour": request.FinalWorkHour,
		"FinalTravelHour": request.FinalTravelHour,
		"DiscountsApplied": request.DiscountsApplied,
		"PartialPayment": request.PartialPayment}).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update bill"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bill)
}