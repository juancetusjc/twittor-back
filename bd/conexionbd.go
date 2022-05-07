package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ConnectBD config to BDs */
var MongoCN = ConnectBD()
var clientOptios = options.Client().ApplyURI("mongodb+srv://twittorjc:aiq8dmUkSBqXPFC@clusterjc.joxv6.mongodb.net/test")

/*ConnectBD config to BDs */
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptios)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connections sucess to BDs")
	return client
}

/*CheckConnection ping BDs */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
