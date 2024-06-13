package models

import (
	"time"

	"gorm.io/gorm"
)

/*
At the right of each attribute struct will be the configuration of data types for the database
If you find just an int or string element without specification, thesse attributes take the
int, text value respectly by default.
*/

type Person struct {
	gorm.Model
	PID       uint   `gorm:"column:pid;primaryKey;autoIncrement"`
	Address   string `gorm:"column:address;not null;varchar(50)"`
	Name      string `gorm:"column:name;not null;varchar(100)"`
	LastName  string `gorm:"column:last_name;not null;varchar(100)"`
	TNumber   string `gorm:"column:t_number;not null;unique;varchar(12)"`
	Password  string `gorm:"column:password;not null;varchar(25)"`
	Email     string `gorm:"column:email;not null;varchar(100);unique"`
	OwnerID   uint   `gorm:"not null"`
	OwnerType string `gorm:"not null"`
}

// Declare the name in database
func (Person) TableName() string {
	return "person"
}

type User struct {
	gorm.Model
	Person        Person `gorm:"polymorphic:Owner;"`
	PublicService []byte `gorm:"column:public_service;not null"`
}

func (User) TableName() string {
	return "duser"
}

type Professional struct {
	gorm.Model
	Person           Person    `gorm:"polymorphic:Owner;"`
	ProfilePicture   []byte    `gorm:"column:profile_picture;not null"`
	Birth            time.Time `gorm:"column:birth"`
	IdentifyDocument string    `gorm:"column:identify_document;not null;unique;varchar(15)"`
	PhotoDocument    []byte    `gorm:"column:photo_document;not null;unique"`
}

func (Professional) TableName() string {
	return "professional"
}

type Service struct {
	gorm.Model
	SID         uint   `gorm:"column:sid;primaryKey;autoIncrement"`
	Type        string `gorm:"column:type;varchar(100);unique"`
	Description string `gorm:"column:description"`
	// true State means that is active now the service, false means the service is unuse
	State bool `gorm:"column:state;varchar(30);default:true"`
}

func (Service) TableName() string {
	return "service"
}

type ProfessionalOffer struct {
	gorm.Model
	SID                       uint    `gorm:"column:sid;primaryKey"`
	PID                       uint    `gorm:"column:pid;primaryKey"`
	Major                     string  `gorm:"column:major;varchar(30)"`
	RelationalExperience      string  `gorm:"column:relational_experience;varchar(30)"`
	RelationalExperiencePhoto []byte  `gorm:"column:relational_experience_photo"`
	MajorPhoto                []byte  `gorm:"column:major_photo"`
	UnitPrice                 float64 `gorm:"column:unit_price;not null"`
	PricePerHour              float64 `gorm:"column:price_per_hour;not null"`

	Service      Service      `gorm:"foreignKey:sid;references:sid"`
	Proffesional Professional `gorm:"foreignKey:pid;references:ID"` // cambiar por professional
}

func (ProfessionalOffer) TableName() string {
	return "professional_offer"
}

type Request struct {
	gorm.Model
	RID            uint      `gorm:"column:rid;primaryKey;autoIncrement"`
	UserID         uint      `gorm:"column:user_id;not null"`
	ProfessionalID uint      `gorm:"column:professional_id;not null"`
	SID            uint      `gorm:"column:sid;not null"`
	TravelHour     time.Time `gorm:"column:travel_hour;not null"`
	State          string    `gorm:"column:state;varchar(50);default:'Travel'"`

	User         User         `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET DEFAULT"`
	Professional Professional `gorm:"foreignKey:professional_id;references:ID;constraint:OnDelete:SET DEFAULT"`
	Service      Service      `gorm:"foreignKey:sid;references:sid"`
}

func (Request) TableName() string {
	return "request"
}

type Bill struct {
	gorm.Model
	BID              uint      `gorm:"column:bid;primaryKey;autoIncrement"`
	RID              uint      `gorm:"column:rid;not null"`
	InitWorkHour     time.Time `gorm:"column:init_work_hour;not null"`
	FinalWorkHour    time.Time `gorm:"column:final_work_hour;not null"`
	FinalTravelHour  time.Time `gorm:"column:final_travel_hour;not null"`
	DiscountsApplied float64   `gorm:"column:discounts_applied"`
	PartialPayment   float64   `gorm:"column:partial_payment;not null"`

	Request Request `gorm:"foreignKey:rid;references:rid"`
}

func (Bill) TableName() string {
	return "bill"
}

type Payment struct {
	gorm.Model
	PYID          uint    `gorm:"column:pyid;primaryKey;autoIncrement"`
	BID           uint    `gorm:"column:bid;not null"`
	TotalPayment  float64 `gorm:"column:total_payment;not null"`
	Nequi         bool    `gorm:"column:nequi"`
	Transferencia bool    `gorm:"column:transferencia"`
	Efectivo      bool    `gorm:"column:efectivo"`

	Bill Bill `gorm:"foreignKey:bid;references:bid"`
}

func (Payment) TableName() string {
	return "payment"
}

type Punctuation struct {
	gorm.Model
	SPID         uint `gorm:"column:spid;primaryKey;autoIncrement"`
	RID          uint `gorm:"column:rid;not null"`
	GeneralScore int  `gorm:"column:general_score"`

	Request Request `gorm:"foreignKey:rid;references:rid"`
}

func (Punctuation) TableName() string {
	return "punctuation"
}

type PunctuationType struct {
	gorm.Model
	SPTID           uint `gorm:"column:sptid;primaryKey;autoIncrement"`
	SPID            uint `gorm:"column:spid;not null"`
	TimeTravelPoint int  `gorm:"column:time_travel_point"`
	KindnessPoint   int  `gorm:"column:kindness_point"`
	TimeWorkPoint   int  `gorm:"column:time_work_point"`
	QualityPoint    int  `gorm:"column:quality_point"`

	Punctuation Punctuation `gorm:"foreignKey:spid;references:spid"`
}

func (PunctuationType) TableName() string {
	return "punctuation_type"
}
