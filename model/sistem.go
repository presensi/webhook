package model

import (
	// "time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"context"
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
    Umur  int                `bson:"umur" json:"umur"`
}

// Untuk Kehadrian
type Kehadiran struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Date      string             `bson:"date"`
	Name      string             `bson:"name"`
	Subject   string             `bson:"subject"`
	Status    string             `bson:"status"`
}

func AddKehadiran(db *mongo.Database, kehadiran Kehadiran) error {
	_, err := db.Collection("kehadiran").InsertOne(context.TODO(), kehadiran)
	return err
}

func GetAllKehadiran(db *mongo.Database) ([]Kehadiran, error) {
	var kehadiran []Kehadiran
	cursor, err := db.Collection("kehadiran").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &kehadiran)
	return kehadiran, err
}

func UpdateKehadiran(db *mongo.Database, kehadiran Kehadiran) error {
	_, err := db.Collection("kehadiran").UpdateOne(context.TODO(), bson.M{
		"date": kehadiran.Date,
		"name": kehadiran.Name,
		"subject": kehadiran.Subject,
	}, bson.M{"$set": kehadiran})
	return err
}

func DeleteKehadiran(db *mongo.Database, date string, name string, subject string) error {
	_, err := db.Collection("kehadiran").DeleteOne(context.TODO(), bson.M{
		"date": date,
		"name": name,
		"subject": subject,
	})
	return err
}
