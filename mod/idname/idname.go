package idname

import (
	"fmt"
	"time"

	"github.com/gocroot/helper/atdb"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IDNameMasuk(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	if !Pesan.LiveLoc {
		return "Minimal share live location dulu lah kak " + Pesan.Alias_name
	}
	longitude := fmt.Sprintf("%f", Pesan.Longitude)
	latitude := fmt.Sprintf("%f", Pesan.Latitude)
	lokasiuser, err := GetLokasi(db, Pesan.Longitude, Pesan.Latitude)
	if err != nil {
		return "Mohon maaf kak, kakak " + Pesan.Alias_name + " belum berada di lokasi presensi, silahkan menuju lokasi presensi dahulu baru bisa presensi masuk kak."
	}
	if lokasiuser.Nama == "" {
		return "Nama nya kok kosong kak?! " + Pesan.Alias_name
	}
	dtuser := &PresensiLokasi{
		PhoneNumber: Pesan.Phone_number,
		Lokasi:      lokasiuser,
		IsMasuk:     true,
		CreatedAt:   time.Now(),
	}
	_, err = atdb.InsertOneDoc(db, "presensi", dtuser)
	if err != nil {
		return "Gagal insert ke database nih kak :( " + Pesan.Alias_name
	}

	return "Hai.. hai.. kakak atas nama:\n" + Pesan.Alias_name + "\nLongitude: " + longitude + "\nLatitude: " + latitude + "\nLokasi: " + lokasiuser.Nama + "\nTelah berhasil melakukan absensi masuk\nTerimakasih"
}

func IDNamePulang(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	if !Pesan.LiveLoc {
		return "Minimal share live location dulu lah kak " + Pesan.Alias_name
	}
	longitude := fmt.Sprintf("%f", Pesan.Longitude) 	
	latitude := fmt.Sprintf("%f", Pesan.Latitude)
	lokasiuser, err := GetLokasi(db, Pesan.Longitude, Pesan.Latitude)
	if err != nil {
		return "Mohon maaf kak " + Pesan.Alias_name + ", kakak belum berada di lokasi presensi, silahkan menuju lokasi presensi dahulu baru bisa presensi pulang."
	}
	if lokasiuser.Nama == "" {
		return "Nama nya kok kosong kak?! " + Pesan.Alias_name
	}
	dtuser := &PresensiLokasi{
		PhoneNumber: Pesan.Phone_number,
		Lokasi:      lokasiuser,
		IsMasuk:     false,
		CreatedAt:   time.Now(),
	}
	filter := bson.M{"_id": atdb.TodayFilter(), "cekinlokasi.phonenumber": Pesan.Phone_number, "ismasuk": true}
	docselfie, err := atdb.GetOneLatestDoc[PresensiSelfie](db, "selfie", filter)
	if err != nil {
		return "Kakak " + Pesan.Alias_name + " belum selfie masuk ini :( " + err.Error()
	}
	if docselfie.CekInLokasi.Lokasi.ID != lokasiuser.ID {
		return "Lokasi pulang nya harus sama dengan lokasi masuknya dong kak " + Pesan.Alias_name + ".\nLokasi : " + lokasiuser.Nama
	}
	_, err = atdb.InsertOneDoc(db, "presensi", dtuser)
	if err != nil {
		return "Gagal insert ke database nih kak :( " + Pesan.Alias_name
	}

	return "Hai.. hai.. kakak atas nama:\n" + Pesan.Alias_name + "\nLongitude: " + longitude + "\nLatitude: " + latitude + "\nLokasi: " + lokasiuser.Nama + "\nTelah berhasil melakukan absensi keluar\nTerimakasih"
}

func GetLokasi(mongoconn *mongo.Database, long float64, lat float64) (lokasi Lokasi, err error) {
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

	lokasi, err = atdb.GetOneDoc[Lokasi](mongoconn, "lokasi", filter)
	if err != nil {
		return
	}
	return
}
