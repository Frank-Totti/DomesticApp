package main

import (
	"net/http"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/Frank-totti/DomesticApp/routes"
	"github.com/gorilla/mux"
)

func createDummyUser() {

	dummy := models.User{
		Person: models.Person{
			Address:  "Nothing",
			Name:     "Default",
			LastName: "User",
			TNumber:  "Nothoing",
		},
		Email:         "default@example.com",
		PublicService: []byte{}}

	config.Db.FirstOrCreate(&dummy.Person)
	config.Db.FirstOrCreate(&dummy)

}

func migrateDatabase() {

	config.Db.AutoMigrate(models.Person{})
	config.Db.AutoMigrate(models.User{})
	config.Db.AutoMigrate(models.Professional{})
	config.Db.AutoMigrate(models.Professional{})
	config.Db.AutoMigrate(models.Service{})
	//config.Db.AutoMigrate(models.PriceType{})
	config.Db.AutoMigrate(models.ProfessionalOffer{})
	config.Db.AutoMigrate(models.Request{})
	config.Db.AutoMigrate(models.Bill{})
	config.Db.AutoMigrate(models.Payment{})
	config.Db.AutoMigrate(models.Punctuation{})
	config.Db.AutoMigrate(models.PunctuationType{})

}

func main() {

	config.Conn()

	migrateDatabase()

	createDummyUser()

	domesticApp := mux.NewRouter()

	domesticApp.HandleFunc("/", routes.HomeHandler)

	//////////////////////////////////////////////////////////////////////// User Routes
	domesticApp.HandleFunc("/users/search", routes.GetUsersHandler).Methods("GET")      // get all users
	domesticApp.HandleFunc("/users/search/{id}", routes.GetUserHandler).Methods("GET")  // get an specific user
	domesticApp.HandleFunc("/users/create", routes.CreateUserHandler).Methods("POST")   // create a user
	domesticApp.HandleFunc("/users/update", routes.UpdateUserHandler).Methods("PUT")    // update a user
	domesticApp.HandleFunc("/users/delete", routes.DeleteUserHandler).Methods("DELETE") // delete a user

	//////////////////////////////////////////////////////////////////////// Services Routes

	domesticApp.HandleFunc("/services/search/true", routes.GetActiveServices).Methods("GET")
	domesticApp.HandleFunc("/services/search/false", routes.GetNotActiveServices).Methods("GET")
	domesticApp.HandleFunc("/services/create", routes.CreateService).Methods("POST")
	domesticApp.HandleFunc("/services/update/TD", routes.UpdateTypeDescriptionService).Methods("PUT") // TD = Type or Description
	domesticApp.HandleFunc("/services/update/setTrue", routes.SetTrueServiceState)
	domesticApp.HandleFunc("/services/update/setFalse", routes.SetFalseServiceState)

	//////////////////////////////////////////////////////////////////////// Request Routes
	domesticApp.HandleFunc("/requests/create",routes.CreateRequest).Methods("POST")
	domesticApp.HandleFunc("/request/search/true", routes.GetActiveRequests).Methods("GET")
	domesticApp.HandleFunc("/request/search/false", routes.GetNotActiveRequests).Methods("GET")
	domesticApp.HandleFunc("/request/update/TH", routes.UpdateTravelHour).Methods("PUT")
	domesticApp.HandleFunc("/request/update/setTrue", routes.SetTrueRequestState)
	domesticApp.HandleFunc("/request/update/setFalse", routes.SetFalseRequestState)



	http.ListenAndServe(":3000", domesticApp)

}
