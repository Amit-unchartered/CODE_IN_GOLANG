package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Amit-unchartered/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://learncodeonline:amit@cluster0.zwm6vi7.mongodb.net/?retryWrites=true&w=majority" //where your database is located

const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB

// init is a specialised method in golang that is going to run only for the very first time when the entire application is going to execute
func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString) //this has not actually fired up the connection request
	//it is done below

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection sucess")

	collection = client.Database(dbName).Collection(colName)

	//if the collection reference is ready for me
	fmt.Println("Collection instance is ready")

}

// MONGODB helpers - file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie in db with id: ", inserted.InsertedID)

}

// update 1 record
func updateOneMovie(movieID string) {
	//we need to convert this id to be able to read by mongodb(i.e. _id)
	id, _ := primitive.ObjectIDFromHex(movieID)

	//everything inside mongodb is not json, it is bson(binary JS object notation)
	//As per the documentation bson.D should be used if the order of elements matters and bson.M should be used otherwise.
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}} //we are not checking whether it is true, everytime just we are making it true.

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one record
func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllmovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	//we need to look into the documentation for the type of value the methods of mogodb return

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount //if we donot want to print we will return the value.
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode") //what types of content we are allowing
	w.Header().Set("Allow-Control-Allow-Methods", "POST")              //i will allow only this method

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode") //what types of content we are allowing
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")               //i will allow only this method

	params := mux.Vars(r) //now this params have key value pairs, i just need to accept it
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode") //what types of content we are allowing
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")            //i will allow only this method

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode") //what types of content we are allowing
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")            //i will allow only this method

	count := deleteAllmovies()
	json.NewEncoder(w).Encode(count)
}
