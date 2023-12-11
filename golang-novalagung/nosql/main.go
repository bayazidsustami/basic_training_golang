package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func main() {
	//insert()
	//update()
	delete()
	find()
	aggragateData()
}

func connect() (*mongo.Database, error) {
	clientOption := options.Client()
	clientOption.ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert Success")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "wick"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name	:", result[0].Name)
		fmt.Println("Grade	:", result[0].Grade)
	}

}

func update() {
	db, err := connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "wick"}
	var changes = student{"Jhon Wick", 3}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Update Success!")
}

func delete() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Ethan"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("remove success")
}

func aggragateData() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": null,
				"Total": { "$sum": 1 }
			} },
			{ "$project": {
				"Total": 1,
				"_id": 0
			} }
		]
		`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}

}
