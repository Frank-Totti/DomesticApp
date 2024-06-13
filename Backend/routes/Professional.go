package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/forms"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/gorilla/mux"
)

func CreateProffesionalHandler(w http.ResponseWriter, r *http.Request) {

	var professional models.Professional // Create the struct to save them

	err := json.NewDecoder(r.Body).Decode(&professional) // get the request data in the client

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if transation.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	// Use the library crypto/bcrypt to Encrypt the password
	crypt_password, err := HashPassword(professional.Person.Password)

	if err != nil { // If something wrong with crypt the password
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save the password"})
		return
	}

	professional.Person.Password = crypt_password

	if err := transation.Create(&professional.Person).Error; err != nil {

		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create person"})
		return

	}

	if err := transation.Create(&professional).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create professional"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&professional)
}

func GetProffesionalHandlerById(w http.ResponseWriter, r *http.Request) {
	var professional models.Professional

	params := mux.Vars(r)

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").First(&professional, params["id"]).Error; err != nil {
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
	json.NewEncoder(w).Encode(&professional)

}

func GetProfessionalHandlerByEmail(w http.ResponseWriter, r *http.Request) {
	var user models.Professional

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

	if err := transation.Preload("Person").Table("professional").
		Joins("JOIN person ON person.owner_id = professional.ID AND person.owner_type = 'professional'").
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

func GetProfessionalHandlerByName(w http.ResponseWriter, r *http.Request) {
	var users []models.Professional

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

	if err := transation.Preload("Person").Table("professional").
		Joins("JOIN person ON professional.ID = person.owner_id AND person.owner_type = 'professional'").
		Where("TRIM(person.name) LIKE ? AND person.email <> 'defaultProfessional@example.com'", request.Name+"%").
		Find(&users).Error; err != nil {
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

func GetProfessionalHandlerByLastName(w http.ResponseWriter, r *http.Request) {
	var users []models.Professional

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
		Table("professional").
		Joins("JOIN person ON professional.ID = person.owner_id AND person.owner_type = 'professional'").
		Where("person.email <> 'defaultProfessional@example.com' AND TRIM(person.last_name) LIKE ?", request.LastName+"%").
		Find(&users).Error; err != nil {
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

func GetProffesionalsHandler(w http.ResponseWriter, r *http.Request) {

	var professionals []models.Professional

	transation := config.Db.Begin()

	if transation.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").Find(&professionals).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed Find Person data"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&professionals)

}

func UpdateProffesionalHandler(w http.ResponseWriter, r *http.Request) {

	var professional models.Professional

	var professionalRequest forms.ProfesionalForm

	if err := json.NewDecoder(r.Body).Decode(&professionalRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if transation.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	if err := transation.Preload("Person").First(&professional, professionalRequest.ID).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}

	profilePictureData, err := base64.StdEncoding.DecodeString(professionalRequest.Professional.ProfilePicture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid base64 image data"})
		return
	}

	photoDocumentData, err := base64.StdEncoding.DecodeString(professionalRequest.Professional.PhotoDocument)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid base64 image data"})
		return
	}

	/////////////////////////////////////////////////////////////////////////////// Professional Updates

	if len(profilePictureData) > 0 {
		professional.ProfilePicture = profilePictureData
	}

	if len(photoDocumentData) > 0 {
		professional.PhotoDocument = photoDocumentData
	}

	if professionalRequest.Professional.IdentifyDocument != "" {
		professional.IdentifyDocument = professionalRequest.Professional.IdentifyDocument
	}

	if !professionalRequest.Professional.Birth.IsZero() {
		professional.Birth = professionalRequest.Professional.Birth
	}

	//////////////////////////////////////////////////////////////////////////// Professional.Person updates

	if professionalRequest.Professional.Person.Email != "" {
		professional.Person.Email = professionalRequest.Professional.Person.Email
	}

	if professionalRequest.Professional.Person.Address != "" {
		professional.Person.Address = professionalRequest.Professional.Person.Address
	}

	if professionalRequest.Professional.Person.Name != "" {
		professional.Person.Name = professionalRequest.Professional.Person.Name
	}

	if professionalRequest.Professional.Person.LastName != "" {
		professional.Person.LastName = professionalRequest.Professional.Person.LastName
	}

	if professionalRequest.Professional.Person.TNumber != "" {
		professional.Person.TNumber = professionalRequest.Professional.Person.TNumber
	}

	// Password verifications

	if professionalRequest.Professional.Person.Password != "" {
		if CheckPasswordHash(professionalRequest.Professional.Person.Password, professional.Person.Password) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "The new password can not be your actual password, write another one"})
			transation.Rollback()
			return
		} else {
			professional.Person.Password = professionalRequest.Professional.Person.Password
		}
	}

	//err := transation.Model(&OldUser.Person).Updates(newUser.Person).Error; err != nil

	if err := transation.Save(&professional.Person).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update person data"})
		return
	}

	// err := transation.Model(&OldUser).Updates(newUser).Error; err != nil
	if err := transation.Save(&professional).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update user data"})
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&professional)

}

func DeleteProffesioanlHandler(w http.ResponseWriter, r *http.Request) {

	var professional models.Professional

	var request forms.Professional_id

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	transation := config.Db.Begin()

	if err := transation.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}
	//err := transation.Preload("Person").First(&user, userRequest.ID).Error; err != ni
	if err := transation.Preload("Person").First(&professional, request.ID).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed Find Person data"})
		return
	}

	var dummyProfessional models.Professional
	if err := transation.Table("person").Where("email = ?", "defaultProfessional@example.com").First(&dummyProfessional).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find dummy user"})
		return
	}

	if err := transation.Model(&models.Request{}).Where("professional_id = ?", professional.ID).Update("professional_id", dummyProfessional.ID).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update requests"})
		return
	}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 75b4821 (repair of delete functions)
	if err := transation.Unscoped().Delete(&professional.Person).Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find professional"})
		return
	}

	if err := transation.Unscoped().Delete(&professional).Error; err != nil {
<<<<<<< HEAD
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find professional"})
=======
	if err := config.Db.Delete(&models.Professional{}, request.ID).Error; err != nil {
=======
>>>>>>> 75b4821 (repair of delete functions)
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find professional"})
>>>>>>> d5680f8 (partial actualization)
		return
	}

	if err := transation.Commit().Error; err != nil {
		transation.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&professional)
}

func GetGeneralPunctuationProfessionals(w http.ResponseWriter, r *http.Request) {

	var professionals []models.Professional

	transaction := config.Db.Begin()

	if err := transaction.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

<<<<<<< HEAD
<<<<<<< HEAD
	if err := transaction.Preload("Person").Table("professional"). //Select("professional.*").
		//Joins("JOIN professional ON person.owner_id = professional.id").
		Joins("JOIN request ON professional.ID = request.Professional_ID ").
=======
	/*
		transaction.Table("people").Select("people.Name,people.Last_Name").
			Joins("JOIN professionals ON people.owner_id = professionals.id").
			Joins("JOIN requests ON people.P_ID = requests.professional_id").
			Joins("JOIN punctuations ON requests.R_ID = punctuations.R_ID").
			Order("punctuations.general_score DESC")
	*/

	if err := transaction.Table("person").Select("person.Name, person.Last_Name").
		Joins("JOIN professional ON person.owner_id = professional.id").
		Joins("JOIN request ON person.PID = request.Professional_ID ").
>>>>>>> d5680f8 (partial actualization)
=======
	if err := transaction.Preload("Person").Table("professional"). //Select("professional.*").
		//Joins("JOIN professional ON person.owner_id = professional.id").
		Joins("JOIN request ON professional.ID = request.Professional_ID ").
>>>>>>> 76553de (repair of users/request route)
		Joins("JOIN punctuation ON request.RID = punctuation.RID").
		Order("punctuation.general_score DESC").
		Find(&professionals).Error; err != nil {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&professionals)

}

func GetParticularPunctuationProfessional(w http.ResponseWriter, r *http.Request) {

	var professionals []models.Professional

	var request forms.ParticularPunctuation

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if err := transaction.Preload("person").Table("professional").
		//Joins("JOIN professional ON person.owner_id = professional.id").
=======
	/*

		transaction.Table("people").Select("people.Name, people.Last_Name").
			Joins("JOIN professionals ON people.owner_id = professionals.id").
			Joins("JOIN requests ON people.P_ID = requests.professional_id").
			Joins("JOIN punctuations ON requests.R_ID = punctuations.R_ID").
			Joins("JOIN punctuation_types ON punctuations.sp_id = punctuation_types.sp_id").
			Order("punctuation_types." + request.Punctuation + " DESC").
			Find(&professionals)
	*/

=======
>>>>>>> 75b4821 (repair of delete functions)
	if err := transaction.Table("person").Select("person.name, person.last_name").
		Joins("JOIN professional ON person.owner_id = professional.id").
>>>>>>> d5680f8 (partial actualization)
=======
	if err := transaction.Preload("person").Table("professional").
		//Joins("JOIN professional ON person.owner_id = professional.id").
>>>>>>> 76553de (repair of users/request route)
		Joins("JOIN request ON professional.id = request.professional_id ").
		Joins("JOIN punctuation ON request.rid = punctuation.rid").
		Joins("JOIN punctuation_type ON punctuation.spid = punctuation_type.spid").
		Order("punctuation_type." + request.Punctuation + " DESC").
		Find(&professionals).Error; err != nil {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&professionals)

}

func GetProfessionalRequests(w http.ResponseWriter, r *http.Request) {

	var request forms.ProfessionalRequestHistory
	var totalRequestDone int
	var profesionalRequestDone []models.Payment
	var response forms.ProfessionalWriterHistory
<<<<<<< HEAD
	var proveProfessional models.Professional
=======
>>>>>>> 4f6aa2b (just for prove)

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
	if err := transaction.Table("professional").Where("professional.id = ?", request.ID).First(&proveProfessional).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to find user"})
		return
	}
=======
	/*
		transaction.
			Preload("Bill.Request.Professional.Person").Preload("Bill.Request.User.Person").
			Preload("Bill.Request.Professional").Preload("Bill.Request.User").
			Preload("Bill.Request.Service").Preload("Bill").
			Select("payment.*").
			Joins("JOIN bill ON bill.bid = payment.bid").
			Joins("JOIN request ON request.rid = bill.rid").
			Joins("JOIN professional ON professional.id = request.professional_id").
			Where("professional.id = ? ", request.ID).Find(&profesionalRequestDone)
	*/
>>>>>>> 4f6aa2b (just for prove)

	if err := transaction.
		Preload("Bill.Request.Professional.Person").
		Preload("Bill.Request.User.Person").
		Preload("Bill.Request.Service").
		Joins("JOIN bill ON bill.bid = payment.bid").
		Joins("JOIN request ON request.rid = bill.rid").
		Joins("JOIN professional ON professional.id = request.professional_id").
		Where("professional.id = ? ", request.ID).
		Find(&profesionalRequestDone).
		Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get professionals"})
		return
	}

	totalRequestDone = len(profesionalRequestDone)

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	response.ProfessionalID = request.ID
	response.Total = totalRequestDone
	response.RequestHistory = profesionalRequestDone

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

/* CREATE
{
  "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Alberto",
    "LastName":"Valencia García",
    "TNumber":"3008341273",
    "password":"pipotePeludo",
    "Email":"Juan_Albert@gmail.com"
  },
  "ProfilePicture":[],
  "Birth":"2000-01-01T00:00:00Z",
  "IdentifyDocument":"",
  "PhotoDocument":[]
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

/*

{
  "Person":{
    "Address":"Calle 3B #96-19",
    "Name":"Juan Alberto",
    "LastName":"Valencia García",
    "TNumber":"3008341273",
    "password":"pipotePeludo",
    "Email":"Juan_Albert@gmail.com"
  },
  "ProfilePicture":[],
  "Birth":"2000-01-01T00:00:00Z",
  "IdentifyDocument":"",
  "PhotoDocument":"iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH5QEMDhUNgJSJKAAAABl0RVh0Q29tbWVudABDcmVhdGVkIHdpdGggR0lNUFeBDhcAAAGcSURBVDjLpZNNb9pAEIb7I2YIJNaWCyUpgRJgkFKRImCBNgYNF0kDLpSEUJG0GoQSog0CoVRQlSCFQRB7oEz3Pt+58H/SwmmpKa0fv7fmJ29ufzkOBzNZzNdzXbcDXex9AJjX/dGJ5F8FwWA5HI5BIAgBYCAH8Dx+NBo8FmkUjkoRIAorjjvngcc5+PB51vtLo/8AwCKopFHnX6mlStXCFQAAIwaBQaASiKPoa3s9Fr9b6ZpOpVb9XI5/P5fN5VKhVkAgMDv9NE0ASgVkM4+e/AdtfFBBxVxCoRQtVjXAPzxxtQY9+sbHZDsx9Pv92XAfBeFwrFYDD8zVdB+ATAYhBGIq1VrMIzUeMm8DrNAINBa76TgIdip1Xpql+3U8TxHL5nFYn1+pBAEVACyyAQCKIp8MAKi6lCW2wxqtZjJ5ICmNmcaIQtDAyBQarXmGGUggOx2+0AA4/iMaK4i2XyzW5mhmHE7EBnEAw8HJbL1BIRmGxW1bkwlxPB0GoztBqnqlVu2UVHv62m2LxX6+L3n5EkbLoUor+g7Abx42HKwZQh5A0+my+m5xh7pu9kI+sA89jM9VdbUgS6Yl5+6rcZ65x/Pwk2/y7cd0V5E2S6rgKhWmSlnMyalUiuI5DhPdnkTFbAC3qlgg7Tbbr/EfY5jL7FUtQjEL18EQLW2kZ8wEZX01m00WwNn6pbDaqlXK7C1mWI2WYUnqXkmShEGYpxpTTEoV5HfFWsA6XbOZbmN1/dqDZVOS2rZZz+ag/j5M+TRJ/yH3D+ucA1+g9RaCAAAAAElFTkSuQmCC"
}

*/
