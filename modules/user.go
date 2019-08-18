package modules

import (
	"github.com/spadesk1991/easChat-serve/lib/mongo"
	"gopkg.in/mgo.v2/bson"
)

// User 用户
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	UserName string        `json:"user_name" bson:"user_name"`
	Password string        `json:"password" bson:"password"`
	CreateAt string        `json:"create_at" bson:"create_at"`
	UpdateAt string        `json:"update_at" bson:"update_at"`
}

func (user *User) Create() error {
	mongo.GetSession()
	return nil
}
