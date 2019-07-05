package models

import (
	"github.com/stephenrh/gosvc1/internal/app/models/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Password  string        `json:"password" bson:"password"`
	Email     string        `json:"email" bson:"email" binding:"required"`
}

func newUserCollection() *mongo.Collection {
	return mongo.NewCollectionSession("users")
}

func CreateUser(m User) (user User, err error) {
	c := newUserCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	return m, c.Session.Insert(m)
}

func GetUsers() ([]User, error) {
	c := newUserCollection()
	defer c.Close()
	var users []User
	err := c.Session.Find(nil).All(&users)
	return users, err
}

func GetUser(id bson.ObjectId) (User, error) {
	c := newUserCollection()
	defer c.Close()
	var (
		user User
		err  error
	)
	err = c.Session.Find(bson.M{"_id": id}).One(&user)
	return user, err
}
