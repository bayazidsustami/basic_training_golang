package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"grpc-proto/common/config"
	model "grpc-proto/common/model"
	"log"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Couldn't connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Couldn't connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	user := serviceUser()
	fmt.Println("\n", "===========> user test")
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	garage := serviceGarage()
	fmt.Println("\n", "===========> garage test A")

	//using to add garage by user
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	res2, err := garage.List(context.Background(), &model.GarageUserId{
		UserId: user1.Id,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	log.Fatal(user.Register(context.Background(), &user1))
}
