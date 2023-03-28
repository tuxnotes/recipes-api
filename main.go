// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//		Schemes: http
//	 Host: localhost:8080
//		BasePath: /
//		Version: 1.0.0
//		Contact: Mohamed Labouardy <mohamed@labouardy.com> https://labouardy.com
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tuxnotes/recipes-api/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	// recipes = make([]Recipe, 0)
	// recipes = make([]Recipe, 0)
	// file, _ := ioutil.ReadFile("recipes.json")
	// _ = json.Unmarshal([]byte(file), &recipes)
	// ctx = context.Background()
	// client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	// if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// log.Fatal(err)
	// }
	// log.Println("Connected to MongoDB")
	// var listOfRecipes []interface{}
	// for _, recipe := range recipes {
	//     listOfRecipes = append(listOfRecipes, recipe)
	// }
	// collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	//     log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))
	// }
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	// router.GET("/recipes/search", SearchRecipeHandler)
	router.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	router.Run()
}
