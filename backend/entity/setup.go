package entity

import (
	//   "fmt"
	//   "time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&User{}, &Patient{}, &Tenderness{}, &Department{}, &Symptom{},
		&Doctor{}, &Location{}, &Authoritie{}, &Room{}, &Schedule{},
	)

	db = database
	/////////////////////////////////////////
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14) //password mail
	password1, err := bcrypt.GenerateFromPassword([]byte("555555"), 14)

	db.Model(&User{}).Create(&User{
		Name:  "Panadda Srisawat",
		Email: "panadsada@gmail.com",

		Password: string(password),
		Role:     "doctor",
	})

	db.Model(&User{}).Create(&User{
		Name:  "abc",
		Email: "abc@gmail.com",

		Password: string(password1),
		Role:     "user",
	})

	db.Model(&User{}).Create(&User{
		Name:     "Promporn Phinitphong",
		Email:    "promporn@gmail.com",
		Password: string(password),
		Role:     "admin",
	})

	var promporn User
	db.Raw("SELECT * FROM users WHERE email = ?", "promporn@gmail.com").Scan(&promporn)

	var panadda User
	db.Raw("SELECT * FROM users WHERE email = ?", "panadsada@gmail.com").Scan(&panadda)

	var abc User
	db.Raw("SELECT * FROM users WHERE email = ?", "abc@gmail.com").Scan(&abc)

	//---Department Data
	General := Department{
		Name: "แพทย์ทั่วไป",
	}
	db.Model(&Department{}).Create(&General)

	Orthopedics := Department{
		Name: "แพทย์กระดูก",
	}
	db.Model(&Department{}).Create(&Orthopedics)

	Cardiac := Department{
		Name: "แพทย์หัวใจ",
	}
	db.Model(&Department{}).Create(&Cardiac)

	Gynecologist := Department{
		Name: "สูตินารีแพทย์(ตรวจภายใน)",
	}
	db.Model(&Department{}).Create(&Gynecologist)

	Otolaryngology := Department{
		Name: "แพทย์เฉพาะทางด้าน ตา หู คอ จมูก",
	}
	db.Model(&Department{}).Create(&Otolaryngology)

	Psychology := Department{
		Name: "จิตเวช",
	}
	db.Model(&Department{}).Create(&Psychology)

	Phamarceutical := Department{
		Name: "เภสัชกรรม",
	}
	db.Model(&Department{}).Create(&Phamarceutical)

	Skin := Department{
		Name: "แพทย์ผิวหนัง",
	}
	db.Model(&Department{}).Create(&Skin)

	//---Tenderness
	Head := Tenderness{
		Name: "ศีรษะ",
	}
	db.Model(&Tenderness{}).Create(&Head)

	Neck := Tenderness{
		Name: "คอ",
	}
	db.Model(&Tenderness{}).Create(&Neck)

	Ears := Tenderness{
		Name: "หู",
	}
	db.Model(&Tenderness{}).Create(&Ears)

	Eyes := Tenderness{
		Name: "ตา",
	}
	db.Model(&Tenderness{}).Create(&Eyes)

	Nose := Tenderness{
		Name: "จมูก",
	}
	db.Model(&Tenderness{}).Create(&Nose)

	Mouth := Tenderness{
		Name: "ปาก",
	}
	db.Model(&Tenderness{}).Create(&Mouth)

	Back := Tenderness{
		Name: "หลัง",
	}
	db.Model(&Tenderness{}).Create(&Back)

	Thorax := Tenderness{
		Name: "อก",
	}
	db.Model(&Tenderness{}).Create(&Thorax)

	Abdomen := Tenderness{
		Name: "ท้อง",
	}
	db.Model(&Tenderness{}).Create(&Abdomen)

	Pelvis := Tenderness{
		Name: "เชิงกร้าน",
	}
	db.Model(&Tenderness{}).Create(&Pelvis)

	hand := Tenderness{
		Name: "มือ",
	}
	db.Model(&Tenderness{}).Create(&hand)

	muscle := Tenderness{
		Name: "กล้ามเนื้อ",
	}
	db.Model(&Tenderness{}).Create(&muscle)

	skeleton := Tenderness{
		Name: "กระดูก",
	}
	db.Model(&Tenderness{}).Create(&skeleton)

	hip := Tenderness{
		Name: "สะโพก",
	}
	db.Model(&Tenderness{}).Create(&hip)

	knee := Tenderness{
		Name: "เข่า",
	}
	db.Model(&Tenderness{}).Create(&knee)

	ankle := Tenderness{
		Name: "ข้อเท้า",
	}
	db.Model(&Tenderness{}).Create(&ankle)

	//---patient
	db.Model(&Patient{}).Create(&Patient{
		Name:    "Thanawat Nitikarun",
		Email:   "thanawat@gmail.com",
		Blood:   "A",
		Gender:  "Male",
		Disease: "-",
	})

	var Thanawat Patient
	db.Raw("SELECT * FROM doctors WHERE email = ?", "thanawat@gmail.com").Scan(&Thanawat)

	//
	db.Model(&Patient{}).Create(&Patient{
		Name:    "Promporn   Phinitphong",
		Email:   "promporn@gmail.com",
		Blood:   "B",
		Gender:  "Female",
		Disease: "-",
	})

	var Promporn Patient
	db.Raw("SELECT * FROM doctors WHERE email = ?", "promporn@gmail.com").Scan(&Promporn)

	//
	db.Model(&Patient{}).Create(&Patient{
		Name:    "Jakkrit   Chaiwan",
		Email:   "jakkrit@gmail.com",
		Blood:   "O",
		Gender:  "Male",
		Disease: "-",
	})

	var Jakkrit Patient
	db.Raw("SELECT * FROM doctors WHERE email = ?", "jakkrit@gmail.com").Scan(&Jakkrit)

	//
	db.Model(&Patient{}).Create(&Patient{
		Name:    "Wanllaya  Patisang",
		Email:   "wanllaya@gmail.com",
		Blood:   "AB",
		Gender:  "Female",
		Disease: "-",
	})

	var Wanllaya Patient
	db.Raw("SELECT * FROM doctors WHERE email = ?", "wanllaya@gmail.com").Scan(&Wanllaya)

	//---Location Data
	A := Location{
		Name: "Building A",
	}
	db.Model(&Location{}).Create(&A)

	B := Location{
		Name: "Building B",
	}
	db.Model(&Location{}).Create(&B)

	//---Room A Data
	A1 := Room{
		Name: "A101",
	}
	db.Model(&Room{}).Create(&A1)

	A2 := Room{
		Name: "A102",
	}
	db.Model(&Room{}).Create(&A2)

	A3 := Room{
		Name: "A103",
	}
	db.Model(&Room{}).Create(&A3)

	B1 := Room{
		Name: "B101",
	}
	db.Model(&Room{}).Create(&B1)

	B2 := Room{
		Name: "B102",
	}
	db.Model(&Room{}).Create(&B2)

	B3 := Room{
		Name: "B103",
	}
	db.Model(&Room{}).Create(&B3)

	//---Doctor
	db.Model(&Doctor{}).Create(&Doctor{
		Name: "Thanawat Nitikarun",
	})

	DoctorNo1 := Doctor{
		Name: "Panadda Srisawat",
	}
	db.Model(&Doctor{}).Create(&DoctorNo1)

	DoctorNo2 := Doctor{
		Name: "Jakkrit   Chaiwan",
	}
	db.Model(&Doctor{}).Create(&DoctorNo2)

	DoctorNo3 := Doctor{
		Name: "Wallaya  Patisang",
	}
	db.Model(&Doctor{}).Create(&DoctorNo3)
	DoctorNo4 := Doctor{
		Name: "Ratchapol Piyaman",
	}
	db.Model(&Doctor{}).Create(&DoctorNo4)

	DoctorNo5 := Doctor{
		Name: "Thanawat Nitikarun",
	}
	db.Model(&Doctor{}).Create(&DoctorNo5)

}
