package idname

import (
	"context"
	"fmt"
	// "net/http"
	// "strconv"
	"time"

	// "github.com/gocroot/helper/atapi"
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
		return "Mohon maaf kak, kakak " + Pesan.Alias_name + " belum berada di lokasi presensi, silahkan menuju lokasi presensi dahulu baru presensi masuk."
	}
	if lokasiuser.Nama == "" {
		return "Nama nya kosong kak " + Pesan.Alias_name
	}
		// Ambil waktu saat ini
		now := time.Now()

		// Tentukan batas waktu validasi (dari jam 6.30 pagi hingga jam 7.30 pagi)
		startTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 30, 0, 0, now.Location())
		endTime := time.Date(now.Year(), now.Month(), now.Day(), 7, 30, 0, 0, now.Location())
	
		// Periksa apakah waktu saat ini berada di dalam rentang yang diizinkan
		if now.Before(startTime) || now.After(endTime) {
			// Jika di luar jam atas, maka tandai sebagai alpha atau tidak hadir
			return "Maaf kak, Anda hanya dapat melakukan presensi antara jam 6.30 pagi hingga 7.30. Diluar waktu tersebut dianggap alpha atau tidak hadir."
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

			// Ambil waktu saat ini
			now := time.Now()

			// Tentukan batas waktu validasi (dari jam 2 sore hingga jam 6 sore)
			startTime := time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, now.Location())
			endTime := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())
		
			// Periksa apakah waktu saat ini berada di dalam rentang yang diizinkan
			if now.Before(startTime) || now.After(endTime) {
				// Jika di luar jam atas, maka tandai sebagai alpha atau tidak hadir
				return "Maaf kak, Anda hanya dapat melakukan presensi antara jam 6.30 pagi hingga 7.30. Diluar waktu tersebut dianggap alpha atau tidak hadir."
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

// func CekSelfieMasuk(Profile itmodel.Profile, Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
// 	if Pesan.Filedata == "" {
// 		return "Kirim pap nya dulu dong kak.. " + Pesan.Alias_name
// 	}
// 	dt := FaceDetect{
// 		IDUser:    Pesan.Phone_number,
// 		Base64Str: Pesan.Filedata,
// 	}
// 	filter := bson.M{"_id": atdb.TodayFilter(), "phonenumber": Pesan.Phone_number, "ismasuk": true}
// 	pstoday, err := atdb.GetOneDoc[PresensiLokasi](db, "presensi", filter)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf kakak belum cekin share live location hari ini, silahkan share live loc dengan ditambah keyword\n*cekin presensi masuk*\n_" + err.Error() + "_"
// 	}
// 	conf, err := atdb.GetOneDoc[Config](db, "config", bson.M{"phonenumber": Profile.Phonenumber})
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan dalam pengambilan config di database " + err.Error()
// 	}
// 	statuscode, faceinfo, err := atapi.PostStructWithToken[FaceInfo]("secret", conf.LeaflySecret, dt, conf.LeaflyURL)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan pemanggilan API leafly :" + err.Error()
// 	}
// 	if statuscode != http.StatusOK {
// 		if statuscode == http.StatusFailedDependency {
// 			return "Wah kak " + Pesan.Alias_name + " mohon maaf, jangan kaku gitu dong. Ekspresi wajahnya ga boleh sama dengan selfie sebelumnya ya kak. Senyumnya yang lebar, giginya dilihatin, matanya pelototin, hidungnya keatasin.\n\n" + faceinfo.Error
// 		} else if statuscode == http.StatusMultipleChoices {
// 			return "IM$G#M$Gui76557u|||" + faceinfo.FileHash + "|||" + faceinfo.Error
// 		} else {
// 			return "Wah kak " + Pesan.Alias_name + " mohon maaf:\n" + faceinfo.Error + "\nCode: " + strconv.Itoa(statuscode)
// 		}

// 	}
// 	pselfie := PresensiSelfie{
// 		CekInLokasi: pstoday,
// 		IsMasuk:     true,
// 		IDUser:      faceinfo.PhoneNumber,
// 		Commit:      faceinfo.Commit,
// 		Filehash:    faceinfo.FileHash,
// 		Remaining:   faceinfo.Remaining,
// 	}
// 	_, err = atdb.InsertOneDoc(db, "selfie", pselfie)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan input ke database " + err.Error()
// 	}
// 	return "Hai kak, " + Pesan.Alias_name + "\nCekin Masuk di lokasi:" + pstoday.Lokasi.Nama + "\n*Jangan lupa _cekin presensi pulang_ ya kak biar dapat skor*"

// }


// func CekSelfiePulang(Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
// 	if Pesan.Filedata == "" {
// 		return "Kirim pap nya dulu dong kak.. " + Pesan.Alias_name
// 	}
// 	dt := FaceDetect{
// 		IDUser:    Pesan.Phone_number,
// 		Base64Str: Pesan.Filedata,
// 	}
// 	filter := bson.M{"_id": atdb.TodayFilter(), "phonenumber": Pesan.Phone_number} //, "ismasuk": false}
// 	pstoday, err := atdb.GetOneDoc[PresensiLokasi](db, "presensi", filter)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf kakak belum cekin share live location hari ini " + err.Error()
// 	}
// 	conf, err := atdb.GetOneDoc[Config](db, "config", bson.M{"phonenumber": "62895601060000"})
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan dalam pengambilan config di database " + err.Error()
// 	}
// 	statuscode, faceinfo, err := atapi.PostStructWithToken[FaceInfo]("secret", conf.LeaflySecret, dt, conf.LeaflyURL)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan pemanggilan API leafly " + err.Error()
// 	}
// 	if statuscode != http.StatusOK {
// 		if statuscode == http.StatusFailedDependency {
// 			return "Wah kak " + Pesan.Alias_name + " mohon maaf, jangan kaku gitu dong. Ekspresi wajahnya ga boleh sama dengan selfie sebelumnya ya kak. Senyumnya yang lebar, giginya dilihatin, matanya pelototin, hidungnya keatasin.\n\n" + faceinfo.Error
// 		} else if statuscode == http.StatusMultipleChoices {
// 			return "IM$G#M$Gui76557u|||" + faceinfo.FileHash + "|||" + faceinfo.Error
// 		} else {
// 			return "Wah kak " + Pesan.Alias_name + " mohon maaf:\n" + faceinfo.Error + "\nCode: " + strconv.Itoa(statuscode)
// 		}

// 	}
// 	pselfie := PresensiSelfie{
// 		CekInLokasi: pstoday,
// 		IsMasuk:     false,
// 		IDUser:      faceinfo.PhoneNumber,
// 		Commit:      faceinfo.Commit,
// 		Filehash:    faceinfo.FileHash,
// 		Remaining:   faceinfo.Remaining,
// 	}
// 	_, err = atdb.InsertOneDoc(db, "selfie", pselfie)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf ada kesalahan input ke database " + err.Error()
// 	}
// 	filter = bson.M{"_id": atdb.TodayFilter(), "cekinlokasi.phonenumber": Pesan.Phone_number, "ismasuk": true}
// 	selfiemasuk, err := atdb.GetOneLatestDoc[PresensiSelfie](db, "selfie", filter)
// 	if err != nil {
// 		return "Wah kak " + Pesan.Alias_name + " mohon maaf kakak belum selfie masuk. " + err.Error()
// 	}
// 	// Ekstrak timestamp dari ObjectID
// 	objectIDTimestamp := selfiemasuk.ID.Timestamp()
// 	// Dapatkan waktu saat ini
// 	currentTime := time.Now()
// 	// Hitung selisih waktu dalam detik
// 	diff := currentTime.Sub(objectIDTimestamp) //.Seconds()
// 	// Konversi selisih waktu ke jam, menit, dan detik
// 	hours := int(diff.Hours())
// 	minutes := int(diff.Minutes()) % 60
// 	seconds := int(diff.Seconds()) % 60
// 	KetJam := fmt.Sprintf("%d jam, %d menit, %d detik", hours, minutes, seconds)

// 	skor := diff.Seconds() / 18000 //selisih waktu dibagi 8 jam
// 	skorValue := fmt.Sprintf("%f", skor)
// 	//post ke backedn domyikado
// 	datapresensi := PresensiSMA{
// 		ID:          selfiemasuk.ID,
// 		PhoneNumber: Pesan.Phone_number,
// 		Skor:        skor,
// 		KetJam:      KetJam,
// 		LamaDetik:   diff.Seconds(),
// 		Lokasi:      pstoday.Lokasi.Nama,
// 	}
// 	statuscode, httpresp, err := atapi.PostStructWithToken[itmodel.Response]("secret", conf.DomyikadoSecret, datapresensi, conf.DomyikadoPresensiURL)
// 	if err != nil {
// 		return "Akses ke endpoint domyikado gagal: " + err.Error()
// 	}
// 	if statuscode != http.StatusOK {
// 		return "Salah posting endpoint domyikado: " + httpresp.Response + "\ninfo\n" + httpresp.Info
// 	}
// 	return "Hai kak, " + Pesan.Alias_name + "\nBerhasil Presensi Pulang di lokasi:" + pstoday.Lokasi.Nama + "\nHadir selama: " + KetJam + "\n*Skor: " + skorValue + "*"

// }

// func GetLokasi(mongoconn *mongo.Database, long float64, lat float64) (lokasi Lokasi, err error) {
// 	filter := bson.M{
// 		"batas": bson.M{
// 			"$geoIntersects": bson.M{
// 				"$geometry": bson.M{
// 					"type":        "Point",
// 					"coordinates": []float64{long, lat},
// 				},
// 			},
// 		},
// 	}

// 	lokasi, err = atdb.GetOneDoc[Lokasi](mongoconn, "lokasi", filter)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
