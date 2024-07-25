package model

import (
	// "time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Untuk Login
type AdminRegistration struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname        string             `json:"fullname" bson:"fullname" validate:"required"`
	Email           string             `json:"email" bson:"email" validate:"required,email"`
	Password        string             `json:"password" bson:"password" validate:"required,min=8"`
	ConfirmPassword string             `json:"confirm_password" bson:"confirm_password" validate:"required,eqfield=Password"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password" bson:"password" validate:"required,min=8"`
}

// Untuk Data Siswa
type Siswa struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Nama  string             `bson:"nama" json:"nama"`
    Kelas string             `bson:"kelas" json:"kelas"`
    JenisKelamin string      `bson:"jeniskelamin" json:"jeniskelamin"`
	Phonenumber string 		 `bson:"phonenumber" json:"phonenumber"`
}

// Untuk Kehadrian
type Kehadiran struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Date      string             `json:"date" bson:"date"`
	Subject   string             `json:"subject" bson:"subject"`
	Status    string             `json:"status" bson:"status"`
}
