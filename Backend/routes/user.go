package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/forms"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User // Create the struct to save them

	err := json.NewDecoder(r.Body).Decode(&user) // get the request data in the client

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

	if err := transaction.Create(&user.Person).Error; err != nil {

		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create person"})
		return

	}

	if err := transaction.Create(&user).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create user"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Person").First(&user, params["id"]).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	transaction := config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transaction.Preload("Person").Find(&users).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed Find Person data"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	var userRequest forms.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
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

	if err := transaction.Preload("Person").First(&user, userRequest.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if userRequest.UpdateUser.NewEmail != "" {
		user.Email = userRequest.UpdateUser.NewEmail
	}

	if len(userRequest.UpdateUser.PublicService) > 0 {
		user.PublicService = userRequest.UpdateUser.PublicService
	}

	if userRequest.UpdateUser.Person.Address != "" {
		user.Person.Address = userRequest.UpdateUser.Person.Address
	}

	if userRequest.UpdateUser.Person.Name != "" {
		user.Person.Name = userRequest.UpdateUser.Person.Name
	}

	if userRequest.UpdateUser.Person.LastName != "" {
		user.Person.LastName = userRequest.UpdateUser.Person.LastName
	}

	if userRequest.UpdateUser.Person.TNumber != "" {
		user.Person.TNumber = userRequest.UpdateUser.Person.TNumber
	}

	//err := transaction.Model(&OldUser.Person).Updates(newUser.Person).Error; err != nil

	if err := transaction.Save(&user.Person).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update person data"})
		return
	}

	// err := transaction.Model(&OldUser).Updates(newUser).Error; err != nil
	if err := transaction.Save(&user).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update user data"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	var request forms.UserRequest

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
	//err := transaction.Preload("Person").First(&user, userRequest.ID).Error; err != ni
	if err := transaction.Preload("Person").First(&user, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed Find Person data"})
		return
	}

	var dummyUser models.User
	if err := transaction.Where("email = ?", "default@example.com").First(&dummyUser).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find dummy user"})
		return
	}

	if err := transaction.Model(&models.Request{}).Where("user_id = ?", user.ID).Update("user_id", dummyUser.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update requests"})
		return
	}

	if err := transaction.Delete(&models.User{}, request.ID).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

/* CREATE
{
  "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Francesco",
    "LastName":"García Vargas",
    "TNumber":"3008473008"
  },
  "Email":"jcesc.g@hotmail.com",
  "PublicService":""
}
*/

/* UPDATE
{
  "id":8,
  "UpdateUser":{
    "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Alberto",
    "LastName":"Valencia García",
    "TNumber":"3008341273"
  },
  "NewEmail":"JeanAlbert@hotmail.com",
  "PublicService":""
  }
}
*/
