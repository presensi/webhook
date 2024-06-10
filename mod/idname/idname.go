package idname

import (
	"context"
	"fmt"
	"time"

	"github.com/gocroot/helper/atdb"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func IDName(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
// 	if !Pesan.LiveLoc {
// 		return "Minimal share live location dulu lah kak."
// 	}
// 	longitude := fmt.Sprintf("%f", Pesan.Longitude)
// 	latitude := fmt.Sprintf("%f", Pesan.Latitude)

// 	return "Hai.. hai.. kakak atas nama:\n" + Pesan.Alias_name + "\nLongitude: " + longitude + "\nLatitude: " + latitude + "\nLokasi:" + GetLokasi(db, Pesan.Longitude, Pesan.Latitude) + "\nTelah berhasil melakukan absen\nTerimakasih"
// }

func IDNameMasuk(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	if !Pesan.LiveLoc {
		return "Minimal share live location dulu lah kak " + Pesan.Alias_name
	}
	longitude := fmt.Sprintf("%f", Pesan.Longitude)
	latitude := fmt.Sprintf("%f", Pesan.Latitude)
	lokasiuser, err := GetLokasi(db, Pesan.Longitude, Pesan.Latitude)
	if err != nil {
		return "Mohon maaf kak, kakak " + Pesan.Alias_name + " belum berada di lokasi presensi, silahkan menuju lokasi presensi dahulu baru presensi masuk."
	}
	if lokasiuser.Nama == "" {
		return "Nama nya kosong kak " + Pesan.Alias_name
	}
	dtuser := &PresensiLokasi{
		PhoneNumber: Pesan.Phone_number,
		Lokasi:      lokasiuser,
		IsMasuk:     true,
		CreatedAt:   time.Now(),
	}
	_, err = atdb.InsertOneDoc(db, "presensi", dtuser)
	if err != nil {
		return "Gagal insert ke database kak " + Pesan.Alias_name
	}

	return "Hai.. hai.. kakak atas nama:\n" + Pesan.Alias_name + "\nLongitude: " + longitude + "\nLatitude: " + latitude + "\nLokasi:" + lokasiuser.Nama + "\nTelah berhasil melakukan absensi masuk\nTerimakasih"
}

func IDNamePulang(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	if !Pesan.LiveLoc {
		return "Minimal share live location dulu lah kak " + Pesan.Alias_name
	}
	longitude := fmt.Sprintf("%f", Pesan.Longitude)
	latitude := fmt.Sprintf("%f", Pesan.Latitude)
	lokasiuser, err := GetLokasi(db, Pesan.Longitude, Pesan.Latitude)
	if err != nil {
		return "Mohon maaf kak, kakak " + Pesan.Alias_name + " belum berada di lokasi presensi, silahkan menuju lokasi presensi dahulu baru presensi masuk."
	}
	if lokasiuser.Nama == "" {
		return "Nama nya kosong kak " + Pesan.Alias_name
	}
	dtuser := &PresensiLokasi{
		PhoneNumber: Pesan.Phone_number,
		Lokasi:      lokasiuser,
		IsMasuk:     true,
		CreatedAt:   time.Now(),
	}
	_, err = atdb.InsertOneDoc(db, "presensi", dtuser)
	if err != nil {
		return "Gagal insert ke database kak " + Pesan.Alias_name
	}

	return "Hai.. hai.. kakak atas nama:\n" + Pesan.Alias_name + "\nLongitude: " + longitude + "\nLatitude: " + latitude + "\nLokasi:" + lokasiuser.Nama + "\nTelah berhasil melakukan absensi masuk\nTerimakasih"
}

func GetLokasi(mongoconn *mongo.Database, long float64, lat float64) (Lokasi, error) {
	lokasicollection := mongoconn.Collection("lokasi")
	filter := bson.M{
		"batas": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
			},
		},
	}
	var lokasi Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		fmt.Printf("GetLokasi: %v\n", err)
		return Lokasi{}, err
	}
	return lokasi, nil
}
