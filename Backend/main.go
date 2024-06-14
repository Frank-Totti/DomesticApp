package main

import (
	"net/http"
	"time"

	testing "github.com/Frank-totti/DomesticApp/Testing"
	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/models"
	"github.com/Frank-totti/DomesticApp/routes"
	"github.com/gorilla/mux"
)

func createDummyProfessional() {
	dummyProfessional := models.Professional{
		Person: models.Person{
			Address:  "Default_Professional",
			Name:     "Default_Professional",
			LastName: "Default_Professional",
			TNumber:  "0F0",
			Email:    "defaultProfessional@example.com",
			Password: "Default_Professional_Password",
		},
		ProfilePicture:   []byte{},
		Birth:            time.Now(),
		IdentifyDocument: "Default_Professional",
		PhotoDocument:    []byte{},
	}
	config.Db.Create(&dummyProfessional.Person)
	config.Db.Create(&dummyProfessional)

}

func createDummyUser() {

	dummyUser := models.User{
		Person: models.Person{
			Address:  "Default_user",
			Name:     "Default_user",
			LastName: "Default_user",
			TNumber:  "1G0",
			Email:    "defaultUser@example.com",
			Password: "Default_User_Password",
		},
		PublicService: []byte{}}

	config.Db.Create(&dummyUser.Person)
	config.Db.Create(&dummyUser)

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
	createDummyProfessional()

	testing.ExecuteTestingData()

	domesticApp := mux.NewRouter()

	domesticApp.HandleFunc("/", routes.HomeHandler)

	//////////////////////////////////////////////////////////////////////// User Routes
	// Añadir para buscar por nombre, por apellido y por correo electronico
	domesticApp.HandleFunc("/users/search", routes.GetUsersHandler).Methods("GET") // get all users
	domesticApp.HandleFunc("/users/search/id/{id}", routes.GetUserHandlerById).Methods("GET")
	domesticApp.HandleFunc("/users/search/email", routes.GetUserHandlerByEmail).Methods("GET")
	domesticApp.HandleFunc("/users/create", routes.CreateUserHandler).Methods("POST")   // create a user
	domesticApp.HandleFunc("/users/update", routes.UpdateUserHandler).Methods("PUT")    // update a user
	domesticApp.HandleFunc("/users/delete", routes.DeleteUserHandler).Methods("DELETE") // delete a user
	domesticApp.HandleFunc("/users/request", routes.GetUserRequests).Methods("GET")
	domesticApp.HandleFunc("/users/search/name", routes.GetUserHandlerByName).Methods("GET")
	domesticApp.HandleFunc("/users/search/last_name", routes.GetUserHandlerByLastName).Methods("GET")

	//////////////////////////////////////////////////////////////////////// Services Routes
	// Añadir para buscar por nombre de servicio o tipo de servicio
	domesticApp.HandleFunc("/services/search/true", routes.GetActiveServices).Methods("GET")
	domesticApp.HandleFunc("/services/search/false", routes.GetNotActiveServices).Methods("GET")
	domesticApp.HandleFunc("/services/create", routes.CreateService).Methods("POST")
	domesticApp.HandleFunc("/services/update/TD", routes.UpdateTypeDescriptionService).Methods("PUT") // TD = Type or Description
	domesticApp.HandleFunc("/services/update/setTrue", routes.SetTrueServiceState).Methods("PUT")
	domesticApp.HandleFunc("/services/update/setFalse", routes.SetFalseServiceState).Methods("PUT")
	domesticApp.HandleFunc("/services/search/type", routes.GetServiceByName).Methods("GET")

	/////////////////////////////////////////////////////////////////////// Professional Routes
	// Añadir para buscar por nombre, por apellido y por correo electronico
	domesticApp.HandleFunc("/professional/search", routes.GetProffesionalsHandler).Methods("GET")
	domesticApp.HandleFunc("/professional/search/{id}", routes.GetProffesionalHandlerById).Methods("GET")
	domesticApp.HandleFunc("/professional/create", routes.CreateProffesionalHandler).Methods("POST")
	domesticApp.HandleFunc("/professional/update", routes.UpdateProffesionalHandler).Methods("PUT")
	domesticApp.HandleFunc("/professional/delete", routes.DeleteProffesioanlHandler).Methods("DELETE")
	domesticApp.HandleFunc("/professional/punctuation/general", routes.GetGeneralPunctuationProfessionals).Methods("GET")
	domesticApp.HandleFunc("/professional/punctuation/particular", routes.GetParticularPunctuationProfessional).Methods("GET")
	domesticApp.HandleFunc("/professional/search/email", routes.GetProfessionalHandlerByEmail).Methods("GET")
	domesticApp.HandleFunc("/professional/search/name", routes.GetProfessionalHandlerByName).Methods("GET")
	domesticApp.HandleFunc("/professional/search/last_name", routes.GetProfessionalHandlerByLastName).Methods("GET")
	domesticApp.HandleFunc("/professional/request", routes.GetProfessionalRequests).Methods("GET")

	/////////////////////////////////////////////////////////////////////// Professional_offer

	domesticApp.HandleFunc("/professional_offers/create", routes.CreateOffert).Methods("POST")
	domesticApp.HandleFunc("/professional_offers/search/service/type", routes.GetOffertsByServiceType).Methods("GET")
	domesticApp.HandleFunc("/professional_offers/search", routes.GetOfferts).Methods("GET")

	//////////////////////////////////////////////////////////////////////// Request Routes
	domesticApp.HandleFunc("/requests/create", routes.CreateRequest).Methods("POST")
	domesticApp.HandleFunc("/request/search/true", routes.GetActiveRequests).Methods("GET")
	domesticApp.HandleFunc("/request/search/false", routes.GetNotActiveRequests).Methods("GET")
	domesticApp.HandleFunc("/request/update/TH", routes.UpdateTravelHour).Methods("PUT")
	domesticApp.HandleFunc("/request/update/setTrue", routes.SetTrueRequestState)
	domesticApp.HandleFunc("/request/update/setFalse", routes.SetFalseRequestState)

	//////////////////////////////////////////////////////////////////////// Bills Routes
	domesticApp.HandleFunc("/bills/create", routes.CreateBill).Methods("POST")
	domesticApp.HandleFunc("/bills/search", routes.GetBillsHandler).Methods("GET")
	domesticApp.HandleFunc("/bills/search/{id}", routes.GetBillHandler).Methods("GET")
	domesticApp.HandleFunc("/bills/update", routes.UpdateBill).Methods("PUT")

	//////////////////////////////////////////////////////////////////////// Payments Routes
	domesticApp.HandleFunc("/payments/create", routes.CreatePaymentHandler).Methods("POST")
	domesticApp.HandleFunc("/payments/search/{id}", routes.GetPaymentHandler).Methods("GET")
	domesticApp.HandleFunc("/payments/update/TP", routes.UpdateTotalPayment).Methods("PUT")
	domesticApp.HandleFunc("/payments/update/setNequi", routes.SetPaymentMethodNequi)
	domesticApp.HandleFunc("/payments/update/setTransferencia", routes.SetPaymentMethodTransferencia)
	domesticApp.HandleFunc("/payments/update/setEfectivo", routes.SetPaymentMethodEfectivo)

	//////////////////////////////////////////////////////////////////////// Punctuations Routes
	domesticApp.HandleFunc("/punctuation/create", routes.CreatePunctuationHandler).Methods("POST")
	domesticApp.HandleFunc("/punctuation/search/{id}", routes.GetPunctuationHandler).Methods("GET")
	domesticApp.HandleFunc("/punctuation/update/GS", routes.UpdateGeneralScoreHandler).Methods("PUT")
	domesticApp.HandleFunc("/punctuation/delete", routes.DeletePunctuationHandler).Methods("DELETE")

	//////////////////////////////////////////////////////////////////////// Punctuation Types Routes
	domesticApp.HandleFunc("/punctuationt/create", routes.CreatePunctuationType).Methods("POST")
	domesticApp.HandleFunc("/punctuationt/search/{id}", routes.GetPunctuationTypeHandler).Methods("GET")
	domesticApp.HandleFunc("/punctuationt/update", routes.UpdatePunctuationType).Methods("PUT")
	domesticApp.HandleFunc("/punctuationt/delete", routes.DeletePunctuationTypeHandler).Methods("DELETE")

	http.ListenAndServe(":3000", domesticApp)

}
