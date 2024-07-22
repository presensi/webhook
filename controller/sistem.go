package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/mod/idname"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	if registrationData.Password != registrationData.ConfirmPassword {
		http.Error(w, "Password tidak sesuai", http.StatusBadRequest)
		return
	}

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

//DATA SISWA
// AddSiswa menambahkan data siswa baru ke database
func AddSiswa(w http.ResponseWriter, r *http.Request) {
	var siswa model.Siswa
	if err := json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedID, err := atdb.InsertOneDoc(config.Mongoconn, "siswa", siswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Data siswa berhasil disimpan", "insertedID": insertedID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateSiswa memperbarui data siswa yang ada di database
func UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	var siswa model.Siswa
	if err := json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"nama": siswa.Nama}
	update := bson.M{"$set": siswa}

	updateresult, err := atdb.ReplaceOneDoc(config.Mongoconn, "siswa", filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Data siswa berhasil diperbarui", "updatedCount": updateresult.ModifiedCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteSiswa menghapus data siswa dari database
func DeleteSiswa(w http.ResponseWriter, r *http.Request) {
	var siswa model.Siswa
	if err := json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"nama": siswa.Nama}

	deleteresult, err := atdb.DeleteOneDoc(config.Mongoconn, "siswa", filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Data siswa berhasil dihapus", "deletedCount": deleteresult.DeletedCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllSiswa mengambil semua data siswa dari database
func GetAllSiswa(w http.ResponseWriter, r *http.Request) {
	var siswaList []model.Siswa

	results, err := atdb.GetAllDoc[[]model.Siswa](config.Mongoconn, "siswa", bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, result := range results {
		var siswa model.Siswa
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &siswa)
		siswaList = append(siswaList, siswa)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(siswaList)
}

//Presemsi
// GetAllPresensi mengambil semua data presensi dari database
func GetAllPresensi(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("presensi")
	cur, err := collection.Find(r.Context(), bson.M{}, options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())

	var presensiList []idname.PresensiLokasi
	if err = cur.All(r.Context(), &presensiList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(presensiList)
}
