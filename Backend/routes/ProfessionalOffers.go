package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/forms"
	"github.com/Frank-totti/DomesticApp/models"
)

func CreateOffert(w http.ResponseWriter, r *http.Request) {

	var offer models.ProfessionalOffer

	var service models.Service

	var professional models.Professional

	err := json.NewDecoder(r.Body).Decode(&offer)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return

	}

	transaction := config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Table("professional_offer").Where("sid = ? AND pid = ?", offer.SID, offer.PID).First(&offer).Error; err != nil {

		if err_service := transaction.Table("service").Where("sid = ?", offer.SID).First(&service).Error; err_service != nil {
			transaction.Rollback()
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"Error": "service doesn't exists"})
			return
		}

		if err_professional := transaction.Preload("Person").Table("professional").Where("id = ?", offer.PID).First(&professional).Error; err_professional != nil {
			transaction.Rollback()
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"Error": "professioanl doesn't exists"})
			return
		}

		if err2 := transaction.Create(&offer).Error; err2 != nil {
			transaction.Rollback()
			w.WriteHeader(http.StatusConflict)
			return
		}

		offer.Service = service
		offer.Professional = professional

	} else {
		transaction.Rollback()
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]string{"Error": "offer already exist"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&offer)

}

func GetOffertsByServiceType(w http.ResponseWriter, r *http.Request) {

	var offerts []models.ProfessionalOffer
	var serviceType forms.SearchServiceName

	err := json.NewDecoder(r.Body).Decode(&serviceType)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return

	}

	transaction := config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Professional.Person").Preload("Service").Table("professional_offer").
		Joins("JOIN service ON professional_offer.sid = service.sid").Where("service.type LIKE ? AND service.state = ?", serviceType.Type+"%", true).Find(&offerts).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Failed to find service type"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&offerts)

}

func GetOfferts(w http.ResponseWriter, r *http.Request) {

	var offerts []models.ProfessionalOffer

	transaction := config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Professional.Person").Preload("Service").Find(&offerts).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Failed to find service type"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&offerts)

}

func GetOffertsByProfessionalName(w http.ResponseWriter, r *http.Request) {

	var offerts []models.ProfessionalOffer
	var request forms.SearchName

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return

	}

	transaction := config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Professional.Person").Preload("Service").Table("professional_offer").
		Joins("JOIN professional ON professional_offer.sid = professional.sid").
		Where("professional LIKE ? AND service.state = ?", serviceType.Type+"%", true).Find(&offerts).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Failed to find service type"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&offerts)
}

func GetOffertsByProfessionalLastName(w http.ResponseWriter, r *http.Request) {

}

/*
CREATE
{
  "SID":15,
  "PID":45,
  "Major":"Ingenierio Industrial",
  "RelationalExperience":"Nothing",
  "RelationalExperiencePhoto":"",
  "MajorPhoto":"",
  "UnitPrice":30000,
  "PricePerHour":10000
}

*/
