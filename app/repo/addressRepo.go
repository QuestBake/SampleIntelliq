package repo

import (
	db "pracSpace/restHandler_Gin/app/config"
	"pracSpace/restHandler_Gin/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type addressRepository struct {
	db   string
	coll *mgo.Collection
}

//NewAddressRepository new object
func NewAddressRepository() *addressRepository {
	coll := db.GetCollection(dbName, "addresses")
	if coll == nil {
		return nil
	}
	return &addressRepository{
		dbName, coll,
	}
}

func (repo *addressRepository) closeSession() {
	repo.coll.Database.Session.Close()
}

func (repo *addressRepository) Save(address *model.Address) error {
	defer repo.closeSession()
	err := repo.coll.Insert(address)
	return err
}

func (repo *addressRepository) Update(address *model.Address) error {
	defer repo.closeSession()
	err := repo.coll.Update(bson.M{"_id": address.ID}, address)
	return err
}

func (repo *addressRepository) Delete(addressID bson.ObjectId) error {
	defer repo.closeSession()
	err := repo.coll.Remove(bson.M{"_id": addressID})
	return err
}

func (repo *addressRepository) FindByID(addressID bson.ObjectId) (*model.Address, error) {
	defer repo.closeSession()
	var address model.Address
	err := repo.coll.FindId(addressID).One(&address)
	if err != nil {
		return nil, err
	}
	return &address, err
}

func (repo *addressRepository) FindByCity(city string) (*model.Address, error) {
	defer repo.closeSession()
	var address model.Address

	err := repo.coll.Find(bson.M{"city": city}).One(&address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repo *addressRepository) FindAll() (model.Addresses, error) {
	defer repo.closeSession()
	var addresses model.Addresses
	err := repo.coll.Find(nil).All(&addresses)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (repo *addressRepository) FindByState(state string) (model.Addresses, error) {
	defer repo.closeSession()
	var addresses model.Addresses
	err := repo.coll.Find(bson.M{"state": state}).All(&addresses)
	if err != nil {
		return nil, err
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
