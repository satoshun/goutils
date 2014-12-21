package mgodb

import (
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	dbname  string
	url     string
	db      *mgo.Database
)

type Config struct {
	DBName string `json:"dbname"`
	Debug  bool   `json:"debug"`
	Url    string `json:"url"`
}

func GetCollection(colName string) *mgo.Collection {
	return GetDB().C(colName)
}

func GetDB() *mgo.Database {
	if db == nil {
		db = session.DB(dbname)
	}

	return db
}

func setSession(debug bool) {
	mgo.SetDebug(debug)
	_session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	_session.SetSafe(&mgo.Safe{})
	session = _session
}

func InitApp(ur, dbn string, debug bool) {
	url = ur
	dbname = dbn

	setSession(debug)
}

func InitAppObject(config Config) {
	InitApp(config.Url, config.DBName, config.Debug)
}
