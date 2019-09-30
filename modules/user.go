package modules

import (
	"github.com/spadesk1991/easChat-server/lib/mongo"
	"github.com/spadesk1991/easChat-server/lib/utils"
	"gopkg.in/mgo.v2/bson"
)

// User 用户
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id" example:"1111"`
	UserName string        `json:"user_name" bson:"user_name" example:"张三"`
	Password string        `json:"password" bson:"password" example:"密码"`
	CreateAt string        `json:"create_at" bson:"create_at" example:"创建时间"`
	UpdateAt string        `json:"update_at" bson:"update_at" example:"更新时间"`
}
type aaa struct {
	AA string `json:"aa,omitempty" bson:"aa" example:"aa"`
	BB string `json:"bb,omitempty" bson:"bb" example:"bb"`
}

func (user *User) Create() error {
	s, db := mongo.GetSession()
	defer s.Close()
	c := db.C("user")
	user.ID = bson.NewObjectId()
	user.CreateAt = utils.GetCurrTime()
	user.UpdateAt = utils.GetCurrTime()
	return c.Insert(user)
}

func GetUser(query bson.M) (User, error) {
	s, db := mongo.GetSession()
	defer s.Close()
	c := db.C("user")
	var user User
	err := c.Find(query).One(&user)
	return user, err
}

func GetUsers() ([]User, error) {
	s, db := mongo.GetSession()
	defer s.Close()
	c := db.C("user")
	users := make([]User, 0)
	err := c.Find(bson.M{}).All(&users)
	return users, err
}
