package modules

import (
	"github.com/spadesk1991/easChat-server/lib/mongo"
	"github.com/spadesk1991/easChat-server/lib/utils"
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
	s, db := mongo.GetSession()
	defer s.Close()
	c := db.C("user")
	user.ID = bson.NewObjectId()
	user.CreateAt = utils.GetCurrTime()
	return c.Insert(user)
}
