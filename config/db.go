package config

import (
	"github.com/gocroot/helper/atdb"
	"os"

	"github.com/gocroot/model"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = model.DBInfo{
	DBString: MongoString,
	DBName:   "webhook",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(atdb.DBInfo(mongoinfo))
