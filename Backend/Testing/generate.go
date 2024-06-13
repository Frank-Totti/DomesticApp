package Testing

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Frank-totti/DomesticApp/config"
	"github.com/Frank-totti/DomesticApp/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomBytes(length int) []byte {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return b
}

// Funciones de inserción para cada tabla
func InsertUsers(db *gorm.DB) {
	for i := 0; i < 50; i++ {
		crypt_password, err := HashPassword(fmt.Sprintf("Contraseña %d", i+1))
		if err != nil {
			log.Fatal("Imposible to Hash Password")
			return
		}
		person := models.Person{
			Address:  fmt.Sprintf("Dirección %d", i+1),
			Name:     fmt.Sprintf("Nombre %d", i+1),
			LastName: fmt.Sprintf("Apellido %d", i+1),
			TNumber:  fmt.Sprintf("Número %012d", i+1),
			Password: crypt_password,
			Email:    fmt.Sprintf("email%d@example.com", i+1),
		}
		user := models.User{PublicService: GenerateRandomBytes(16), Person: person}
		db.Create(&user.Person)
		db.Create(&user)
	}
}

func InsertProfessionals(db *gorm.DB) {
	for i := 50; i < 100; i++ {
		person := models.Person{
			Address:  fmt.Sprintf("Dirección %d", i+1),
			Name:     fmt.Sprintf("Nombre %d", i+1),
			LastName: fmt.Sprintf("Apellido %d", i+1),
			TNumber:  fmt.Sprintf("Número %012d", i+1),
			Password: fmt.Sprintf("Contraseña %d", i+1),
			Email:    fmt.Sprintf("email%d@example.com", i+1),
		}
		profesional := models.Professional{Person: person, Birth: time.Now(), IdentifyDocument: fmt.Sprintf("Documento %d", i), PhotoDocument: GenerateRandomBytes(16), ProfilePicture: GenerateRandomBytes(16)}
		db.Create(&profesional.Person)
		db.Create(&profesional)
	}
}

func InsertServices(db *gorm.DB) {
	for i := 0; i < 50; i++ {
		service := models.Service{
			Type:        fmt.Sprintf("Tipo %d", i+1),
			Description: fmt.Sprintf("Descripción %d", i+1),
		}
		db.Create(&service)
	}
}

// Insertar registros en la tabla ProfessionalOffer y relacionar con servicios existentes y profesionales aleatorios
func InsertProfessionalOffers(db *gorm.DB) {
	// Obtener servicios existentes
	services := []models.Service{}
	db.Find(&services)

	// Obtener profesionales existentes
	professionals := []models.Professional{}
	db.Find(&professionals)

	for i := 0; i < 50; i++ {
		service := services[rand.Intn(len(services))]
		professional := professionals[rand.Intn(len(professionals))]

		professionalOffer := models.ProfessionalOffer{
			SID:                       service.SID,
			PID:                       professional.ID,
			Major:                     fmt.Sprintf("Especialidad %d", i+1),
			RelationalExperience:      fmt.Sprintf("Experiencia relacional %d", i+1),
			RelationalExperiencePhoto: GenerateRandomBytes(16),
			MajorPhoto:                GenerateRandomBytes(16),
			UnitPrice:                 float64(i + 1),
			PricePerHour:              float64(i + 1),
		}
		db.Create(&professionalOffer)
	}
}

// Insertar registros en la tabla Request y relacionar con usuarios, profesionales y servicios aleatorios
func InsertRequests(db *gorm.DB) {
	// Obtener usuarios existentes
	users := []models.User{}
	db.Find(&users)

	// Obtener profesionales existentes
	professionals := []models.Professional{}
	db.Find(&professionals)

	// Obtener servicios existentes
	services := []models.Service{}
	db.Find(&services)

	for i := 0; i < 50; i++ {
		user := users[rand.Intn(len(users))]
		professional := professionals[rand.Intn(len(professionals))]
		service := services[rand.Intn(len(services))]

		request := models.Request{
			UserID:         user.ID,
			ProfessionalID: professional.ID,
			SID:            service.SID,
			TravelHour:     time.Now(), // Hora de viaje aleatoria
			State:          fmt.Sprintf("Estado %d", i+1),
		}
		db.Create(&request)
	}
}

// Insertar registros en la tabla Bill y relacionar con solicitudes existentes
func InsertBills(db *gorm.DB) {
	// Obtener solicitudes existentes
	requests := []models.Request{}
	db.Find(&requests)

	for _, request := range requests {
		bill := models.Bill{
			RID:              request.RID,
			InitWorkHour:     time.Now(), // Hora de inicio del trabajo aleatoria
			FinalWorkHour:    time.Now(), // Hora de finalización del trabajo aleatoria
			FinalTravelHour:  time.Now(), // Hora de finalización del viaje aleatoria
			DiscountsApplied: float64(rand.Intn(100)),
			PartialPayment:   float64(rand.Intn(100)),
		}
		db.Create(&bill)
	}
}

// Insertar registros en la tabla Payment y relacionar con facturas existentes
func InsertPayments(db *gorm.DB) {
	// Obtener facturas existentes
	bills := []models.Bill{}
	db.Find(&bills)

	for _, bill := range bills {
		payment := models.Payment{
			BID:           bill.BID,
			TotalPayment:  float64(rand.Intn(100)),
			Nequi:         rand.Float32() < 0.5, // Generar aleatoriamente un valor booleano
			Transferencia: rand.Float32() < 0.5, // Generar aleatoriamente un valor booleano
			Efectivo:      rand.Float32() < 0.5, // Generar aleatoriamente un valor booleano
		}
		db.Create(&payment)
	}
}

func InsertPunctuations(db *gorm.DB) {
	// Obtener solicitudes existentes
	requests := []models.Request{}
	db.Find(&requests)

	for _, request := range requests {
		punctuation := models.Punctuation{
			RID:          request.RID,
			GeneralScore: rand.Intn(5) + 1, // Generar una calificación aleatoria entre 1 y 5
		}
		db.Create(&punctuation)
	}
}

func InsertPunctuationTypes(db *gorm.DB) {
	// Obtener puntuaciones existentes
	punctuations := []models.Punctuation{}
	db.Find(&punctuations)

	for _, punctuation := range punctuations {
		punctuationType := models.PunctuationType{
			SPID:            punctuation.SPID,
			TimeTravelPoint: rand.Intn(5) + 1, // Generar una puntuación aleatoria entre 1 y 5
			KindnessPoint:   rand.Intn(5) + 1,
			TimeWorkPoint:   rand.Intn(5) + 1,
			QualityPoint:    rand.Intn(5) + 1,
		}
		db.Create(&punctuationType)
	}
}

func ExecuteTestingData() {
	InsertUsers(config.Db)
	InsertProfessionals(config.Db)
	InsertServices(config.Db)
	InsertProfessionalOffers(config.Db)
	InsertRequests(config.Db)
	InsertBills(config.Db)
	InsertPayments(config.Db)
	InsertPunctuations(config.Db)
	InsertPunctuationTypes(config.Db)
}
