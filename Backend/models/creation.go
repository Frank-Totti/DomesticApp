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
	PID       uint   `gorm:"primaryKey;autoIncrement"`
	Address   string `gorm:"not null;varchar(50)"`
	Name      string `gorm:"not null;varchar(100)"`
	LastName  string `gorm:"not null;varchar(100)"`
	TNumber   string `gorm:"not null;unique;varchar(12)"`
	OwnerID   uint
	OwnerType string
}

type User struct {
	gorm.Model
	Person        Person `gorm:"polymorphic:Owner;"`
	Email         string `gorm:"not null;unique;varchar(30)"`
	PublicService []byte
}

type Professional struct {
	gorm.Model
	Person           Person    `gorm:"polymorphic:Owner;"`
	ProfilePicture   []byte    `gorm:"not null"`
	Birth            time.Time `gorm:"not null"`
	IdentifyDocument string    `gorm:"not null;unique;varchar(15)"`
	PhotoDocument    []byte    `gorm:"not null;unique"`
}

type Service struct {
	gorm.Model
	SID         uint   `gorm:"primaryKey;autoIncrement"`
	Type        string `gorm:"not null;varchar(100);unique"`
	Description string
	State       bool `gorm:"varchar(30);default:true"`
}

/*
	type PriceType struct {
		gorm.Model
		PTID  uint   `gorm:"primaryKey;autoIncrement"`
		Type  string `gorm:"not null;varchar(30)"`
		Value uint   `gorm:"not null"`

		ProfessionalOffer ProfessionalOffer `gorm:"foreignKey:PTID;references:PTID"`
	}
*/
type ProfessionalOffer struct {
	gorm.Model
	SID uint `gorm:"primaryKey"`
	PID uint `gorm:"primaryKey"`
	//PTID                      uint   `gorm:"primaryKey"`
	Major                     string `gorm:"varchar(30)"`
	RelationalExperience      string `gorm:"varchar(30)"`
	RelationalExperiencePhoto []byte
	MajorPhoto                []byte
	UnitPrice                 float32
	PricePerHour              float32

	Service Service `gorm:"foreignKey:SID;references:SID"`
	Person  Person  `gorm:"foreignKey:PID;references:PID"` // cambiar por professional
	//PriceType PriceType `gorm:"foreignKey:PTID;references:PTID"`
}

type Request struct {
	gorm.Model
	RID            uint `gorm:"primaryKey;autoIncrement"`
	UserID         uint `gorm:"not null"`
	ProfessionalID uint `gorm:"not null"`
	SID            uint `gorm:"not null"`
	TravelHour     int
	State          string `gorm:"varchar(50)"`

	User         User         `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET DEFAULT"`
	Professional Professional `gorm:"foreignKey:ProfessionalID;references:ID;constraint:OnDelete:SET DEFAULT"`
	Service      Service      `gorm:"foreignKey:SID;references:SID"`
}

type Bill struct {
	gorm.Model
	BID              uint `gorm:"primaryKey;autoIncrement"`
	RID              uint `gorm:"not null"`
	InitWorkHour     time.Time
	FinalWorkHour    time.Time
	FinalTravelHour  time.Time
	DiscountsApplied float64
	PartialPayment   float64

	Request Request `gorm:"foreignKey:RID;references:RID"`
}

type Payment struct {
	gorm.Model
	PYID          uint `gorm:"primaryKey;autoIncrement"`
	BID           uint `gorm:"not null"`
	TotalPayment  float64
	Nequi         bool
	Transferencia bool
	Efectivo      bool

	Bill Bill `gorm:"foreignKey:BID;references:BID"`
}

type Punctuation struct {
	gorm.Model
	SPID         uint `gorm:"primaryKey;autoIncrement"`
	RID          uint `gorm:"not null"`
	GeneralScore int

	Request Request `gorm:"foreignKey:RID;references:RID"`
}

type PunctuationType struct {
	gorm.Model
	SPTID           uint `gorm:"primaryKey;autoIncrement"`
	SPID            uint `gorm:"not null"`
	TimeTravelPoint int
	KindnessPoint   int
	TimeWorkPoint   int
	QualityPoint    int

	Punctuation Punctuation `gorm:"foreignKey:SPID;references:SPID"`
}
