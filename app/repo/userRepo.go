package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	db "pracSpace/restHandler_Gin/app/config"
	"pracSpace/restHandler_Gin/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type userRepository struct {
	db   string
	coll *mgo.Collection
}

//NewUserRepository new object
func NewUserRepository() *userRepository {
	coll := db.GetCollection(dbName, "users")
	if coll == nil {
		return nil
	}
	return &userRepository{
		dbName, coll,
	}
}

func (repo *userRepository) closeSession() {
	repo.coll.Database.Session.Close()
}

func (repo *userRepository) FindAllAggregate() (model.Users, error) {
	defer repo.closeSession()

	/*	lookupQ := bson.M{
			"$lookup": bson.M{
				"from":         "addresses",
				"localField":   "address",
				"foreignField": "_id",
				"as":           "address",
			},
		}
	*/
	matchQ := bson.M{
		"$match": bson.M{
			"gender": "Male",
		},
	}

	projetcQ := bson.M{
		"$project": bson.M{
			"_id":     0,
			"name":    1,
			"gender":  1,
			"cuisine": 1,
			"address": 1,
		},
	}

	pipeline := []bson.M{matchQ, projetcQ}

	pipe := repo.coll.Pipe(pipeline)
	result := []bson.M{}
	var users model.Users
	err := pipe.All(&users)
	//	err := repo.coll.Find(nil).All(&users)

	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)

	jsonResp, merr := json.Marshal(result)
	if merr != nil {
		fmt.Println(err)
	}

	// We can get the data just fine as JSON
	fmt.Printf("\nJSON DATA\n")
	jsonRes := string(jsonResp)
	fmt.Println(jsonRes)

	return users, nil
}

func (repo *userRepository) FindAll() (model.Users, error) {
	defer repo.closeSession()
	var users model.Users
	err := repo.coll.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	return users, nil
}

func (repo *userRepository) FindWithRegex() (model.Users, error) {
	defer repo.closeSession()
	var users model.Users
	regex := bson.M{"$regex": bson.RegEx{Pattern: "Chinese", Options: "i"}}
	err := repo.coll.Find(bson.M{"cuisine": regex}).All(&users)
	if err != nil {
		panic(err)
	}
	return users, nil
}

func (repo *userRepository) FindAllWithSelect() (model.Users, error) {
	defer repo.closeSession()
	var users model.Users

	filter := bson.M{"name": "Raghav"}
	cols := bson.M{"name": 1, "cuisine": 1, "_id": 0}

	err := repo.coll.Find(filter).Select(cols).All(&users)
	if err != nil {
		panic(err)
	}
	return users, nil
}

func (repo *userRepository) FindByNameAndMobile(user *model.User) (*model.User, error) {
	defer repo.closeSession()
	regexName := bson.M{"$regex": bson.RegEx{Pattern: user.Name, Options: "i"}}
	andFilter := bson.M{"$and": []bson.M{
		{
			"name": regexName,
		},
		{
			"mobile": user.Mobile,
		},
	},
	}
	var loggedUser model.User
	err := repo.coll.Find(andFilter).One(&loggedUser)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	return &loggedUser, nil
}

func (repo *userRepository) SearchByIndex(term string) (model.Users, error) {
	defer repo.closeSession()
	var users model.Users
	query := bson.M{
		"$text": bson.M{
			"$search": term,
		},
	}
	err := repo.coll.Find(query).All(&users)
	if err != nil {
		panic(err)
	}
	return users, nil
}
