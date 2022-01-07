package routes

import (
	"context"
	"crm/db"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//model
type Cms struct {
	Id      primitive.ObjectID `json:_id bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name,omitempty"`
	Company string             `json:"company" bson:"company,omitempty"`
	Email   string             `json:"email" bson:"email,omitempty"`
	Phone   int                `json:"phone" bson:"phone,omitempty"`
}

func GetLists(w http.ResponseWriter, r *http.Request) {
	var lists []Cms
	cur, err := db.ConnectDB().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var list Cms
		err := cur.Decode(&list)
		if err != nil {
			log.Fatal(err)
		}

		lists = append(lists, list)
	}

	json.NewEncoder(w).Encode(lists)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var list Cms
	var Params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(Params["id"])

	filter := bson.M{"_id": id}
	err := db.ConnectDB().FindOne(context.TODO(), filter).Decode(&list)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	var list Cms

	_ = json.NewDecoder(r.Body).Decode(&list)
	result, err := db.ConnectDB().InsertOne(context.TODO(), list)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateLists(w http.ResponseWriter, r *http.Request) {
	var list Cms
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&list)

	update := bson.D{
		{"$set", bson.D{
			{"name", list.Name},
			{"company", list.Company},
			{"email", list.Email},
			{"phone", list.Phone},
		}},
	}

	err := db.ConnectDB().FindOneAndUpdate(context.TODO(), filter, update).Decode(&list)

	if err != nil {
		log.Fatal(err)
	}

	list.Id = id
	json.NewEncoder(w).Encode(list)
}

func DeletLists(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := db.ConnectDB().DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(deleteResult)
}
