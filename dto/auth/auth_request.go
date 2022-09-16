package authdto

type RegisterRequest struct { //Sebenarnya isi dari RegisterRequest itu sama dengan isi dari user_request
	Name     string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}

// Proses autentifikasi ialah bagaimana cara mengenali orang yang masuk ke dalam system
