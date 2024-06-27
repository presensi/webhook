package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// RegisterHandler menghandle permintaan registrasi admin.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var registrationData model.AdminRegistration

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registrationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lakukan validasi dan pemrosesan data di sini
	if registrationData.Password != registrationData.ConfirmPassword {
		http.Error(w, "Password tidak sesuai", http.StatusBadRequest)
		return
	}

	// Simpan data ke database atau lakukan tindakan lain yang diperlukan
	_, err = atdb.InsertOneDoc(config.Mongoconn, "user", registrationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Registrasi berhasil"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUser mengambil informasi user dari database berdasarkan email dan password.
func GetUser(respw http.ResponseWriter, req *http.Request) {
	var loginDetails model.User
	if err := json.NewDecoder(req.Body).Decode(&loginDetails); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	var user model.User
	filter := bson.M{"email": loginDetails.Email, "password": loginDetails.Password}
	user, err := atdb.GetOneDoc[model.User](config.Mongoconn, "user", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	helper.WriteJSON(respw, http.StatusOK, user)
}

// AddSiswa inserts a new student record in the database
func AddSiswa(siswa Siswa) error {
	_, err := siswaCollection.InsertOne(context.Background(), siswa)
	return err
}

// UpdateSiswa updates an existing student record in the database
func UpdateSiswa(nama string, siswa Siswa) error {
	filter := bson.M{"nama": nama}
	update := bson.M{
		"$set": siswa,
	}
	_, err := siswaCollection.UpdateOne(context.Background(), filter, update)
	return err
}

// DeleteSiswa deletes a student record from the database
func DeleteSiswa(nama string) error {
	filter := bson.M{"nama": nama}
	_, err := siswaCollection.DeleteOne(context.Background(), filter)
	return err
}

// GetAllSiswa retrieves all student records from the database
func GetAllSiswa() ([]Siswa, error) {
	var siswaList []Siswa
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := siswaCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var siswa Siswa
		err := cursor.Decode(&siswa)
		if err != nil {
			return nil, err
		}
		siswaList = append(siswaList, siswa)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return siswaList, nil
}
