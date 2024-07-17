package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// LOGIN

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

// DATA SISWA

// AddSiswa inserts a new student record in the database
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

// UpdateSiswa updates an existing student record in the database
func UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	var siswa model.Siswa
	if err := json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"nama": siswa.Nama}
	update := bson.M{"$set": siswa}

	updateresult, err := atdb.UpdateOneDoc(config.Mongoconn, "siswa", filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Data siswa berhasil diperbarui", "updatedCount": updateresult.ModifiedCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteSiswa deletes a student record from the database
func DeleteSiswa(w http.ResponseWriter, r *http.Request) {
    var siswa model.Siswa
    if err := json.NewDecoder(r.Body).Decode(&siswa); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    filter := bson.M{"nama": siswa.Nama}

    // Memanggil fungsi DeleteOneDoc untuk menghapus dokumen dari koleksi 'siswa'
    updateresult, err := atdb.DeleteOneDoc(config.Mongoconn, "siswa", filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{"message": "Data siswa berhasil dihapus", "deletedCount": updateresult.DeletedCount}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetAllSiswa retrieves all student records from the database
func GetAllSiswa(w http.ResponseWriter, r *http.Request) {
    var siswaList []model.Siswa

    // Mengambil semua dokumen dari koleksi 'siswa'
    results, err := atdb.GetAllDoc[[]model.Siswa](config.Mongoconn, "siswa", bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengubah hasil BSON ke dalam struktur model.Siswa
    for _, result := range results {
        var siswa model.Siswa
        bsonBytes, _ := bson.Marshal(result)
        bson.Unmarshal(bsonBytes, &siswa)
        siswaList = append(siswaList, siswa)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(siswaList)
}

// KEHADIRAN

// AddKehadiran menambahkan catatan kehadiran baru
func AddKehadiran(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var kehadiran model.Kehadiran
	if err := json.NewDecoder(r.Body).Decode(&kehadiran); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := atdb.InsertOneDoc(config.Mongoconn, "kehadiran", kehadiran)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Catatan kehadiran berhasil ditambahkan"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllKehadiran mendapatkan semua catatan kehadiran
func GetAllKehadiran(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var kehadiranList []model.Kehadiran
	results, err := atdb.GetAllDoc[[]model.Kehadiran](config.Mongoconn, "kehadiran", bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, result := range results {
		var kehadiran model.Kehadiran
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &kehadiran)
		kehadiranList = append(kehadiranList, kehadiran)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kehadiranList)
}

// UpdateKehadiran mengupdate catatan kehadiran
func UpdateKehadiran(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var kehadiran model.Kehadiran
	if err := json.NewDecoder(r.Body).Decode(&kehadiran); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"date": kehadiran.Date, "name": kehadiran.Name, "subject": kehadiran.Subject}
	update := bson.M{"$set": kehadiran}

	updateresult, err := atdb.UpdateOneDoc(config.Mongoconn, "kehadiran", filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Catatan kehadiran berhasil diperbarui", "updatedCount": updateresult.ModifiedCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteKehadiran menghapus catatan kehadiran
func DeleteKehadiran(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	date := r.URL.Query().Get("date")
	name := r.URL.Query().Get("name")
	subject := r.URL.Query().Get("subject")

	filter := bson.M{"date": date, "name": name, "subject": subject}

	deleteresult, err := atdb.DeleteOneDoc(config.Mongoconn, "kehadiran", filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Catatan kehadiran berhasil dihapus", "deletedCount": deleteresult.DeletedCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
