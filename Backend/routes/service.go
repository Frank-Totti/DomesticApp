package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/models"
)

func GetActiveServices(w http.ResponseWriter, r *http.Request) {

	var services []models.Service
	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Services").Where("state = ?", true).Find(&services).Error; err != nil {
		transaction.Rollback()

	}

	json.NewEncoder(w).Encode(&services)

}
