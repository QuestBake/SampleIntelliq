package model

import "github.com/globalsign/mgo/bson"

//User user model
type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Age       int16         `json:"age" bson:"age"`
	Mobile    string        `json:"mobile" bson:"mobile"`
	Gender    string        `json:"gender" bson:"gender"`
	Cuisine   []string      `json:"cuisine" bson:"cuisine"`
	Education []education   `json:"education" bson:"education"`
	Address   Address       `json:"address" bson:"address"`
}

type education struct {
	Name string `json:"name" bson:"name"`
	City string `json:"city" bson:"city"`
	Type string `json:"type" bson:"type"`
}

//Users array
type Users []User
