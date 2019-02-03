package repo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"

	"SampleIntelliq/app/config"
	db "SampleIntelliq/app/config"
	"SampleIntelliq/app/model"
)

//AddressRepository type
type AddressRepository struct {
	db   string
	coll *mongo.Collection
}

//NewAddressRepository new object
func NewAddressRepository() *AddressRepository {
	coll := db.GetCollection(config.DbName, "addresses")
	if coll == nil {
		return nil
	}
	return &AddressRepository{
		config.DbName, coll,
	}
}

//Save it inserts one record to collection
func (repo *AddressRepository) Save(address *model.Address) error {
	res, err := repo.coll.InsertOne(config.GetContext(), address)
	if err != nil {
		log.Printf("Insert failed")
	}
	log.Print(res)
	return err
}

//Update updates the document
func (repo *AddressRepository) Update(address *model.Address) error {
	filter := bson.M{"_id": address.ID}
	value := bson.M{"$set": address}
	res, err := repo.coll.UpdateOne(config.GetContext(), filter, value)
	if err != nil {
		log.Printf("Update failed")
	}
	log.Print(res)
	return err
}

//Delete removes record from collection
func (repo *AddressRepository) Delete(addressID primitive.ObjectID) error {
	res := repo.coll.FindOneAndDelete(config.GetContext(), bson.M{"_id": addressID})
	if res == nil {
		return errors.New("Remove failed")
	}
	return nil
}

//FindByID searches by id
func (repo *AddressRepository) FindByID(addressID primitive.ObjectID) (*model.Address, error) {
	var address model.Address
	err := repo.coll.FindOne(context.Background(), addressID).Decode(&address)
	if err != nil {
		return nil, err
	}
	return &address, err
}

//FindByCity searches collection by City
func (repo *AddressRepository) FindByCity(city string) (*model.Address, error) {
	var address model.Address

	err := repo.coll.FindOne(context.Background(), bson.M{"city": city}).Decode(&address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

//FindAll search all documents
func (repo *AddressRepository) FindAll() (model.Addresses, error) {
	var addresses model.Addresses
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cur, err := repo.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var result model.Address
		err := cur.Decode(&result)
		if err != nil {
			log.Print("Unable to decode the results!!")
			return nil, err
		}
		addresses = append(addresses, result)
	}
	return addresses, nil
}

//FindByState search filter by state
func (repo *AddressRepository) FindByState(state string) (model.Addresses, error) {
	var addresses model.Addresses
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.coll.Find(ctx, bson.M{"state": state})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var result model.Address
		err := cur.Decode(&result)
		if err != nil {
			log.Print("Unable to decode the results!!")
			return addresses, err
		}
		addresses = append(addresses, result)
	}
	return addresses, nil
}

/*
func (repo *addressRepository) samplePipeGroupBy() {
	defer repo.closeSession()

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":    "$state",
				"cities": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"cities": -1,
			},
		}, //1: Ascending, -1: Descending
		{
			"$project": bson.M{
				"_id":    1,
				"cities": 1,
			},
		},
	}
	pipe := repo.coll.Pipe(pipeline)

	result := []bson.M{}
	err := pipe.All(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)

	for index, val := range result {
		fmt.Println(index+1, " State= ", val["_id"], " Cities= ", val["cities"])
	}
}*/
