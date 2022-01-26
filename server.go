package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type test struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

func readTest() (result []test, err error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	mongoclient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = mongoclient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := mongoclient.Database("test").Collection("test")
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	err = cur.All(ctx, result)

	return
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"msg":"hello"}`))
	})

	mux.HandleFunc("/test/all", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		result, err := readTest()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ret := map[string]interface{}{
				"err": err.Error(),
			}
			j, _ := json.Marshal(ret)
			w.Write(j)
			return
		}
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(result)
		w.Write(j)
	})

	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	log.Printf("Starting server on %s", ":3000")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
