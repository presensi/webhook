package config

import (
	"log"
	"os"
	"strings"
)

var IPPort, Net = GetAddress()

var PhoneNumber string

func SetEnv() {
	if ErrorMongoconn != nil {
		log.Println(ErrorMongoconn.Error())
	}
	// 	profile, err := atdb.GetOneDoc[model.Profile](Mongoconn, "profile", primitive.M{})
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	PublicKeyWhatsAuth = profile.PublicKey
	// 	WAAPIToken = profile.Token
}

func GetAddress() (ipport string, network string) {
	port := os.Getenv("PORT")
	network = "tcp4"
	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		ip := os.Getenv("IP")
		if ip == "" {
			ipport = ":" + port
		} else {
			if strings.Contains(ip, ".") {
				ipport = ip + ":" + port
			} else {
				ipport = "[" + ip + "]" + ":" + port
				network = "tcp6"
			}
		}
	}
	return
}
