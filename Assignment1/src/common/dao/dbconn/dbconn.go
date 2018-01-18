package dbconn

import (
	m "github.com/jabong/floRest/src/common/mongodb"
	"os"
)

var DbObj m.MongodbInterface
var DbObjSlave m.MongodbInterface

var conErr *m.MongodbError

func InitializeDB() {

	InitializeMongoDB()
	InitializeMongoDBSlave()
}

func InitializeMongoDB() {

	conf := new(m.Config)

	conf.Url = os.Getenv("MONGO_HOST")
	conf.DbName = os.Getenv("MONGO_DB_NAME")
	DbObj, conErr = m.Get(conf)

	if conErr != nil {
		panic("Mongo Connection Error.")
	}

}

func InitializeMongoDBSlave() {

	conf := new(m.Config)

	conf.Url = os.Getenv("MONGO_HOST_SLAVE")
	conf.DbName = os.Getenv("MONGO_DB_NAME")
	DbObjSlave, conErr = m.Get(conf)

	if conErr != nil {
		panic("Mongo Connection Error.")
	}

}
