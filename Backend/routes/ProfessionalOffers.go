package routes

import (
	"encoding/base64"
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

	if err := transaction.Preload("Professional.Person").Preload("Service").Table("professional_offer").Select("Distinct professional_offer.*").
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

	//w.WriteHeader(http.StatusFound)
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
		Joins("JOIN professional ON professional_offer.pid = professional.id").
		Joins("JOIN person ON person.owner_id = professional.id AND person.owner_type = 'professional'").
		Where("person.name LIKE ? AND service.state = ?", request.Name+"%", true).Find(&offerts).Error; err != nil {
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

	var offerts []models.ProfessionalOffer
	var request forms.SearchLastName

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
		Joins("JOIN professional ON professional_offer.pid = professional.id").
		Joins("JOIN person ON person.owner_id = professional.id AND person.owner_type = 'professional'").
		Where("person.last_name LIKE ? AND service.state = ?", request.LastName+"%", true).Find(&offerts).Error; err != nil {
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

func UpdateOffer(w http.ResponseWriter, r *http.Request) {

	var request forms.UpdatePoRequest

	var temporal_offer models.ProfessionalOffer

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

	if err := transaction.Preload("Service").Preload("Professional.Person").Table("professional_offer").Where("professional_offer.pid = ?", request.Professional_id).First(&temporal_offer).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Failed to find professional offer"})
		return
	}

	relational_experience_photo_update, err := base64.StdEncoding.DecodeString(request.RelationalExperiencePhoto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid base64 image data"})
		return
	}

	major_photo_update, err := base64.StdEncoding.DecodeString(request.MajorPhoto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid base64 image data"})
		return
	}

	if request.Major != "" {
		temporal_offer.Major = request.Major
	}

	if request.RelationalExperience != "" {
		temporal_offer.RelationalExperience = request.RelationalExperience
	}

	if len(relational_experience_photo_update) > 0 {
		temporal_offer.RelationalExperiencePhoto = relational_experience_photo_update
	}

	if len(major_photo_update) > 0 {
		temporal_offer.MajorPhoto = major_photo_update
	}

	if request.UnitPrice > 0 {
		temporal_offer.UnitPrice = request.UnitPrice
	}

	if request.PricePerHour > 0 {
		temporal_offer.PricePerHour = request.PricePerHour
	}

	if err := transaction.Save(&temporal_offer).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update offer data"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&temporal_offer)
}

func DeleteOffer(w http.ResponseWriter, r *http.Request) {

	var request forms.Professional_id

	var offer models.ProfessionalOffer

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Service").Preload("Professional.Person").Table("professional_offer").Where("professional_offer.pid = ?", request.ID).First(&offer).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusNotFound) // Use 404 Not Found for missing records
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transaction.Unscoped().Delete(&offer).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

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
