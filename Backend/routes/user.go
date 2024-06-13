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

	// Use the library crypto/bcrypt to Encrypt the password
	crypt_password, err := HashPassword(user.Person.Password)

	if err != nil { // If something wrong with crypt the password
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save the password"})
		return
	}

	user.Person.Password = crypt_password

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

func GetUserHandlerById(w http.ResponseWriter, r *http.Request) {
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

func GetUserHandlerByEmail(w http.ResponseWriter, r *http.Request) {
	var user models.User

	var request forms.SearchEmail

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").Table("duser").
		Joins("JOIN person ON person.owner_id = duser.ID AND person.owner_type = 'duser'").
		Where("person.email = ?", request.Email).First(&user).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}

<<<<<<< HEAD
func GetUserHandlerByName(w http.ResponseWriter, r *http.Request) {
=======
func GetUserHandlerByEmail(w http.ResponseWriter, r *http.Request) {
	var user models.User

	var request forms.UserSearchEmail

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").Table("duser").
		Joins("JOIN person ON person.owner_id = duser.ID AND person.owner_type = 'duser'").
		Where("person.email = ?", request.Email).First(&user).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}

func GetUserHandlerByName(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	var request forms.UserSearchName

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").Table("duser").Joins("JOIN person ON duser.ID = person.owner_id AND person.owner_type = 'duser'").Where("TRIM(person.name) LIKE ? AND person.email <> 'defaultUser@example.com'", request.Name+"%").Find(&users).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandlerByLastName(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	//var personP []models.Person

	var request forms.UserSearchLastName

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Debug().
		Preload("Person").
		Table("duser").
		Joins("JOIN person ON duser.ID = person.owner_id AND person.owner_type = 'duser'").
		Where("person.email <> 'defaultUser@example.com' AND TRIM(person.last_name) LIKE ?", request.LastName+"%").
		Find(&users).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}
	//fmt.Println(users)

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

>>>>>>> 6115b9b (Creation of search email users function)
	var users []models.User

	var request forms.SearchName

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").Table("duser").Joins("JOIN person ON duser.ID = person.owner_id AND person.owner_type = 'duser'").Where("TRIM(person.name) LIKE ? AND person.email <> 'defaultUser@example.com'", request.Name+"%").Find(&users).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandlerByLastName(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	//var personP []models.Person

	var request forms.SearchLastName

	err := json.NewDecoder(r.Body).Decode(&request) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Debug().
		Preload("Person").
		Table("duser").
		Joins("JOIN person ON duser.ID = person.owner_id AND person.owner_type = 'duser'").
		Where("person.email <> 'defaultUser@example.com' AND TRIM(person.last_name) LIKE ?", request.LastName+"%").
		Find(&users).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}
	//fmt.Println(users)

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

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

	if userRequest.UpdateUser.Person.Email != "" {
		user.Person.Email = userRequest.UpdateUser.Person.Email
	}

	if userRequest.UpdateUser.Person.Password != "" {
		if CheckPasswordHash(userRequest.UpdateUser.Person.Password, user.Person.Password) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "The new password can not be your actual password, write another one"})
			transaction.Rollback()
			return
		} else {
			new_password, err := HashPassword(userRequest.UpdateUser.Person.Password)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"Error": "Can't Hashing the password"})
			}
			user.Person.Password = new_password
		}
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

	if err := transaction.Save(&user.Person).Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update person data"})
		return
	}

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
	var request forms.UserDelete

	// Decode the request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Start a transaction
	tx := config.Db.Begin()
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	// Find the user and preload the Person relation
	if err := tx.Preload("Person").First(&user, request.ID).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound) // Use 404 Not Found for missing records
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	// Find the dummy user for reassignment
	var dummyUser models.User
	if err := tx.Table("person").Where("email = ?", "defaultUser@example.com").First(&dummyUser).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find dummy user"})
		return
	}

	// Update requests associated with the user to use the dummy user
	if err := tx.Model(&models.Request{}).Where("user_id = ?", user.ID).Update("user_id", dummyUser.ID).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update requests"})
		return
	}

	// Delete the user
	if err := tx.Unscoped().Delete(&user.Person).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
		return
	}

	if err := tx.Unscoped().Delete(&user).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"success": "User deleted successfully"})
}

func GetUserRequests(w http.ResponseWriter, r *http.Request) {

	var request forms.UserRequestHistory
	var totalRequestDone int
<<<<<<< HEAD
<<<<<<< HEAD
	var userRequestDone []models.Payment
	var response forms.UserWriterHistory
	var proveUser models.User
=======
	var userRequestDone []models.Request
=======
	var userRequestDone []models.Payment
>>>>>>> 6115b9b (Creation of search email users function)
	var response forms.UserWriterHistory
>>>>>>> 76553de (repair of users/request route)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	var transaction = config.Db.Begin()

	if transaction.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}
<<<<<<< HEAD
<<<<<<< HEAD

	if err := transaction.Table("duser").Where("duser.id = ?", request.ID).First(&proveUser).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	if err := transaction.
		Preload("Bill.Request.Professional.Person").
		Preload("Bill.Request.User.Person").
		Preload("Bill.Request.Service").
		Joins("JOIN bill ON bill.bid = payment.bid").
		Joins("JOIN request ON request.rid = bill.rid").
		Joins("JOIN duser ON duser.id = request.user_id").
		Where("duser.id = ? ", request.ID).
		Find(&userRequestDone).
=======
	if err := transaction.Preload("User.Person").Preload("Professional.Person").Preload("User").Preload("Professional").Preload("Service").Select("request.*").
		Joins("JOIN person ON person.owner_id = request.user_id").
		Where("person.owner_id = ? AND person.owner_type = 'duser'", request.ID).Find(&userRequestDone).
>>>>>>> 76553de (repair of users/request route)
=======

	/* transaction.Preload("User.Person").Preload("Professional.Person").Preload("User").Preload("Professional").Preload("Service").Select("request.*").
	Joins("JOIN person ON person.owner_id = request.user_id").
	Where("person.owner_id = ? AND person.owner_type = 'duser'", request.ID).Find(&userRequestDone)


	Select("payment.id,bill.id,request.id,duser.id,payment.total_payment,payment.transferencia,payment.efectivo,payment.nequi,bill.init_work_hour,bill.final_work_hour,bill.final_travel_hour,bill.discounts_applied, bill.partial_payment,request.travel_hour,request.state").
	*/

	if err := transaction.
		Preload("Bill.Request.User.Person").Preload("Bill.Request.Professional.Person").
		Preload("Bill.Request.User").Preload("Bill.Request.Professional").
		Preload("Bill.Request.Service").Preload("Bill").
		Select("payment.*").
		Joins("JOIN bill ON bill.bid = payment.bid").
		Joins("JOIN request ON request.rid = bill.rid").
		Joins("JOIN duser ON duser.id = request.user_id").
		Where("duser.id = ? ", request.ID).Find(&userRequestDone).
>>>>>>> 6115b9b (Creation of search email users function)
		Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get users"})
		return
	}

	totalRequestDone = len(userRequestDone)

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	response.UserID = request.ID
	response.Total = totalRequestDone
	response.RequestHistory = userRequestDone

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

	//if err :=
}

/* CREATE
{
  "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Francesco",
    "LastName":"García Vargas",
    "TNumber":"3008473008",
    "Password":"Su madre",
    "Email":"jcesc.g@hotmail.com"
  },
  "PublicService":""
}
*/

/* UPDATE
{
  "id":52,
  "UpdateUser":{
    "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Alberto",
    "LastName":"Valencia García",
    "TNumber":"3008341273",
    "Password":"JAJA_123JAJ1",
    "Email":"JeanAlbert@hotmail.com"
  },
  "PublicService":""
  }
}
*/
