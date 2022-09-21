package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Alex:A0076BCA@cluster0.vojlwo4.mongodb.net/test"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userColloection := client.Database("Anime").Collection("Users")

	//CREATING docuents using MongoDB

	// insert a single document into a collection
	// create a bson.D object
	user := bson.D{{"firstName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := userColloection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println(result)
	}

	fmt.Print(result.InsertedID)

	users := []interface{}{
		bson.D{{"firstName", "Alex Maina"}, {"age", 45}},
		bson.D{{"firstName", "Dennis Kyusya"}, {"age", 15}},
		bson.D{{"firstName", "Daniel Gakeri"}, {"age", 105}},
		bson.D{{"firstName", "Felster Mulei"}, {"age", 22}},
	}

	// Insert many to the DB
	results, err := userColloection.InsertMany(context.TODO(), users)

	//
	if err != nil {
		panic(err)
	}

	fmt.Println(results)

	// READING DUCUMENTS USING MONGODB ~ FINDMANY
	filter := bson.D{{Key: "$and", Value: bson.A{
		bson.D{
			{Key: "age", Value: bson.D{{Key: "$gt", Value: 25}}},
		},
	}}}

	// Retrieve all data with that filter
	cursor, err := userColloection.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	// converst the data retrieve to bson
	var filterResults []bson.M

	if err = cursor.All(context.TODO(), &filterResults); err != nil {
		panic(err)
	}

	// for _, filterResult := range filterResults {
	// 	fmt.Println(filterResult)
	// }

	// Retreiving data from the DB ~ FINDONE
	var resultFilter bson.M

	if err = userColloection.FindOne(context.TODO(), filter).Decode(&resultFilter); err != nil {
		panic(err)
	}

	//RETREIVING ALL DOCUMENTS IN THE DOCUMENT
	// cursorDocs, err := userColloection.Find(context.TODO(), bson.D{})

	// fmt.Println(cursorDocs)

	// * Replacing a single doc*//
	filterUser := bson.D{{"firstName", "Alex Maina"}}

	//Create Replacemnt Data
	replacement := bson.D{{"firstName", "Lexy Maish"}, {"lastName", "Karumaido"}, {"age", 30}, {"emailAddress", "mainahmwangi12@gmail.com"}}

	//Execute replacement function
	resultReplace, err := userColloection.ReplaceOne(context.TODO(), filterUser, replacement)

	if err != nil {
		panic(err)
	}

	fmt.Println(resultReplace)

	//Deleting Data from Mongo

	filterDelete := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "age", Value: bson.D{{Key: "$gt", Value: 10}}},
				},
			},
		},
	}

	// resultDelete, err := userColloection.DeleteOne(context.TODO(), filterDelete)
	resultDeleteMany, err := userColloection.DeleteMany(context.TODO(), filterDelete)

	fmt.Println(resultDeleteMany)

	if err != nil {
		panic(err)
	}

	// fmt.Println(resultDelete)

}
