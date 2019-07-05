package models

import (
	"fmt"
	"selfscale/users/app/models/mongo"

	"gopkg.in/mgo.v2/bson"
)

// User - Uset struct
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

// CreateUser - Creates a user
func CreateUser(m User) (user User, err error) {
	c := newUserCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	return m, c.Session.Insert(m)
}

// GetUsers - Lists all Users
func GetUsers() ([]User, error) {
	c := newUserCollection()
	defer c.Close()
	var users []User
	err := c.Session.Find(nil).All(&users)
	return users, err
}

// GetUser - Gets a single user
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

// UpdateUser - updates a single user's information
func UpdateUser(user User) (User, error) {
	c := newUserCollection()
	defer c.Close()
	err := c.Session.Update(bson.M{"_id": user.ID}, bson.M{"$set": bson.M{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	}})
	if err != nil {
		fmt.Print("User doesnt exist")
		return User{}, err
	}
	newUser, err := GetUser(user.ID)
	if err != nil {
		fmt.Print("Couldn't return updated user")
		return User{}, err
	}
	return newUser, nil
}
