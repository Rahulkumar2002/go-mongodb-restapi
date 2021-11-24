package models

import (
	"gopkg.in/mgo.v2/bson" //We are using it to convert bson into json and vice versa .
)

type User struct {
	//Id is of type bson.ObjectId and in json we will have it like id and in bson we will have it like _id .
	Id bson.ObjectId `json:"id" bson:"_id"`
	//Name of type string and in both json and bson it will look like name .
	Name string `json:"name" bson:"name"`
	//Gender of type string and in both json and bson it will look same as gender.
	Gender string `json:"gender" bson:"gender"`
	//Age of type integer and in both json and bson it will look same as age.
	Age int `json:"age" bson:"age"`
}
