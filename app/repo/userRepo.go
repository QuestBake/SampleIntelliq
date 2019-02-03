package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	"SampleIntelliq/app/config"
	db "SampleIntelliq/app/config"
	"SampleIntelliq/app/model"
)

type userRepository struct {
	db   string
	coll *mongo.Collection
}

//NewUserRepository new object
func NewUserRepository() *userRepository {
	coll := db.GetCollection(config.DbName, "users")
	if coll == nil {
		return nil
	}
	return &userRepository{
		config.DbName, coll,
	}
}

func (repo *userRepository) FindAllAggregate() (model.Users, error) {

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
	var users model.Users
	cur, err := repo.coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return users, err
	}
	results := []bson.M{}
	for cur.Next(context.Background()) {
		var result bson.M
		cur.Decode(&result)
		results = append(results, result)
	}
	//	err := repo.coll.Find(nil).All(&users)

	if err != nil {
		panic(err)
	}
	fmt.Println("result:", results)

	jsonResp, merr := json.Marshal(results)
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
	var users model.Users
	cur, err := repo.coll.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		var user model.User
		cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) FindWithRegex() (model.Users, error) {

	var users model.Users
	regex := bson.M{"$regex": primitive.Regex{Pattern: "Chinese", Options: "i"}}
	cur, err := repo.coll.Find(context.Background(), bson.M{"cuisine": regex})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		var user model.User
		cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) FindAllWithSelect() (model.Users, error) {
	var users model.Users

	filter := bson.M{"name": "Raghav"}
	cols := bson.M{"name": 1, "cuisine": 1, "_id": 0}
	options := options.Find().SetProjection(cols)
	cur, err := repo.coll.Find(context.Background(), filter, options)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		var user model.User
		cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) FindByNameAndMobile(user *model.User) (*model.User, error) {
	regexName := bson.M{"$regex": primitive.Regex{Pattern: user.Name, Options: "i"}}
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
	err := repo.coll.FindOne(context.Background(), andFilter).Decode(&loggedUser)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	return &loggedUser, nil
}

func (repo *userRepository) SearchByIndex(term string) (model.Users, error) {
	var users model.Users
	query := bson.M{
		"$text": bson.M{
			"$search": term,
		},
	}
	ctx := context.Background()
	cur, err := repo.coll.Find(ctx, query)
	if err != nil {
		panic(err)
	}
	for cur.Next(ctx) {
		var user model.User
		err := cur.Decode(&cur)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}
