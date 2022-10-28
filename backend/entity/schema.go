package entity

import (
	"time"

	"gorm.io/gorm"
)

type Authoritie struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqeIndex"`
	//Password  string
	Schedules []Schedule `gorm:"foreignKey:AuthoritieID"`
}




type Location struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:LocationID"`
}

type Room struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:RoomID"`
}



type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	Role     string
	Symptoms []Symptom `gorm:"foreignKey:DoctorID"`
	Schedules []Schedule `gorm:"foreignKey:AuthoritieID"`
}

type Doctor struct {
	gorm.Model
	Name  string
	Email string `gorm:"unigueIndex"`
	//  Password    string

	Symptoms []Symptom `gorm:"foreignKey:DoctorID"`
}

type Patient struct {
	gorm.Model
	Name    string
	Email   string `gorm:"unigueIndex"`
	Blood   string
	Gender  string
	Disease string

	Symptoms []Symptom `gorm:"foreignKey:PatientID"`
}

type Department struct {
	gorm.Model
	Name string

	Symptoms []Symptom `gorm:"foreignKey:DepartmentID"`
}

type Tenderness struct {
	gorm.Model
	Name string

	Symptoms []Symptom `gorm:"foreignKey:TendernessID"`
}
type Schedule struct {
	gorm.Model
	ScheduleTime time.Time


	// DoctorID เป็น FK
	UserID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	User User `gorm:"references:id"`



	// DoctorID เป็น FK
	DoctorID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	Doctor Doctor `gorm:"references:id"`

	// LocationID  เป็น FK
	LocationID *uint
	// ข้อมูลของ Location เมื่อ join ตาราง
	Location Location `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	AuthoritieID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Authoritie Authoritie `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	RoomID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Room Room `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	DepartmentID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Department Department `gorm:"references:id"`
}


type Symptom struct {
	gorm.Model
	Explain     string
	SymptomTime time.Time

	// DoctorID เป็น FK
	UserID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	User User

	// DoctorID เป็น FK
	DoctorID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	Doctor Doctor

	// PatientID เป็น FK
	PatientID *uint
	// ข้อมูลของ Patient เมื่อ join ตาราง
	Patient Patient

	// TendernessID เป็น FK
	TendernessID *uint
	// ข้อมูลของ Tenderness เมื่อ join ตาราง
	Tenderness Tenderness

	// DepartmentID เป็น FK
	DepartmentID *uint
	// ข้อมูลของ Department เมื่อ join ตาราง
	Department Department
}
