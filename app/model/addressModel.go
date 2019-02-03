package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

//Address address model
type Address struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	City  string             `json:"city" bson:"city"`
	State string             `json:"state" bson:"state"`
}

//Addresses array
type Addresses []Address

//ToString address model
func (addr *Address) ToString() string {
	return "ID: " + addr.ID.String() + " , City: " + addr.City + " , State: " + addr.State
}
