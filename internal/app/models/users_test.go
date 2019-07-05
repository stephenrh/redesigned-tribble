package models

import (
	"testing"

	"github.com/stephenrh/gosvc1/internal/app/models/mongo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func _connect() {
	mongo.MaxPool = 8
	mongo.PATH = "localhost"
	mongo.DBNAME = "users"
	mongo.CheckAndInitServiceConnection()
}

var createdUserID bson.ObjectId

func TestCreateUser(t *testing.T) {
	_connect()
	user := User{
		FirstName: "Testing",
		LastName:  "User",
		Email:     "testing@test.com",
		Password:  "fdsafdsa",
		ID:        bson.NewObjectId(),
	}
	user, err := CreateUser(user)
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, user.ID, 12)
	createdUserID = user.ID
}

func TestListUsers(t *testing.T) {
	_connect()
	users, err := GetUsers()
	if err != nil {
		t.FailNow()
	}
	assert.Len(t, users, 1)
	assert.Equal(t, users[0].ID, createdUserID)
}

func TestFindUser(t *testing.T) {
	_connect()
	user, err := GetUser(createdUserID)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, user.ID, createdUserID)
}
