package helper

import (
	"github.com/gocroot/config"
	"github.com/gocroot/helper/atdb"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDokumen() *string {
	return &config.DOKUMENPANDUAN
}

func GetMongo() *mongo.Database {
	return config.Mongoconn
}

func GetAppProfile(phonenumber string, db *mongo.Database) (apitoken itmodel.Profile, err error) {
	filter := bson.M{"phonenumber": phonenumber}
	apitoken, err = atdb.GetOneDoc[itmodel.Profile](db, "profile", filter)

	return
}
