package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	session  *mgo.Session
	database *mgo.Database
)

func init() {
	var err error
	session, err = mgo.Dial("mongodb://localhost:27017/esaChat")
	if err != nil {
		log.Fatalln(err)
	}
}

// GetSession 获得ｓｅｓｓｉｏｎ
func GetSession() (*mgo.Session, *mgo.Database) {
	s := session.Copy()
	db := s.DB("esaChat")
	return s, db
}
