package idname

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lokasi struct { //lokasi yang bisa melakukan presensi
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty"`
}

type Geometry struct { //data geometry untuk lokasi presensi
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

type DataPresensi struct {
	Lokasi string `bson:"lokasi"`
}

type PresensiLokasi struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phonenumber,omitempty"`
	Lokasi      Lokasi             `bson:"lokasi,omitempty"`
	Selfie      bool               `bson:"selfie,omitempty"`
	IsMasuk     bool               `bson:"ismasuk,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt"`
}
