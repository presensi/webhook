package route

import (
	"log"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
)

// URL mengarahkan permintaan HTTP masuk ke fungsi controller yang tepat.
func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // Jika ini adalah permintaan preflight, kembalikan segera.
	}

	if config.ErrorMongoconn != nil {
		log.Println(config.ErrorMongoconn.Error())
	}

	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.GetHome(w, r)
	case method == "GET" && path == "/refresh/token":
		controller.GetNewToken(w, r)
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	
	//Register
	case method == "POST" && path == "/data/adminregister":
		controller.RegisterHandler(w, r)
	
	//Login
	case method == "POST" && path == "/data/user":
		controller.GetUser(w, r)
	
	//Data Siswa
	case method == "POST" && path == "/data/siswa":
		controller.AddSiswa(w, r)
	case method == "PUT" && path == "/data/siswa":
		controller.UpdateSiswa(w, r)
	case method == "DELETE" && path == "/data/siswa":
		controller.DeleteSiswa(w, r)
	case method == "GET" && path == "/data/siswa":
		controller.GetAllSiswa(w, r)
	
	//Presensi WA
	case method == "GET" && path == "/data/presensi":
		controller.GetAllPresensi(w, r)
		
	//Presensi Manual
	case method == "POST" && path == "/data/kehadiran":
		controller.AddKehadiran(w, r)
	case method == "GET" && path == "/data/kehadiran":
		controller.GetAllKehadiran(w, r)
	case method == "PUT" && path == "/data/kehadiran":
		controller.UpdateKehadiran(w, r)
	case method == "DELETE" && path == "/data/kehadiran":
		controller.DeleteKehadiran(w, r)
	
	default:
		controller.NotFound(w, r)
	}
}
