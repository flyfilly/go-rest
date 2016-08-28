package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Password  string        `json:"password" bson:"password"`
	Email     string        `json:"email" bson:"email"`
}

func (this *User) PrintDetails() string {
	return this.Username + " " + this.Firstname + " " + this.Lastname + " " + this.Email
}
