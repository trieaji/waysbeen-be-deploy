package models

// User model struct
type User struct {
	ID       int             `json:"id" gorm:"primary_key:auto_increment"`
	Name     string          `json:"name" gorm:"type: varchar(255)"`
	Email    string          `json:"email" gorm:"type: varchar(255)"`
	Password string          `json:"password" gorm:"type: varchar(255)"`
	Status   string          `json:"status"`
	Profile  ProfileResponse `json:"profile"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserProfile struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

func (UserProfile) TableName() string {
	return "users"
}

// menyiapkan proses respons relasi
type UsersProfileResponse struct { // membuat respon untuk relasi
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// fungsi agar memberitahu gorm usersprofileresponse bukan sebuah struck yang akan di migrasi ke database
func (UsersProfileResponse) TableName() string {
	return "users"
}
