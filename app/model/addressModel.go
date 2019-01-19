package model

import (
	"github.com/globalsign/mgo/bson"
)

//Address address model
type Address struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	City  string        `json:"city" bson:"city"`
	State string        `json:"state" bson:"state"`
}

//Addresses array
type Addresses []Address

//ToString address model
func (addr *Address) ToString() string {
	return "ID: " + addr.ID.String() + " , City: " + addr.City + " , State: " + addr.State
}
