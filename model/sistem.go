package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

// Siswa represents the data structure for a student
type Siswa struct {
    Nama  string `json:"nama" bson:"nama"`
    Kelas string `json:"kelas" bson:"kelas"`
    Umur  int    `json:"umur" bson:"umur"`
}