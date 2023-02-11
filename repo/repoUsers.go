package repo

import (
	"arrow_food_api/models"
	"arrow_food_api/mongodb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetUserByEmail(email string) (*models.User, error) {
	 db := mongodb.Instance()
	 coll := db.Client.Database("arrowfood").Collection("users");
	 // si no tenemos un modelo, usar el bson.M
	 //var result bson.M

	 var user *models.User

	 err :=coll.FindOne(context.TODO(),bson.D{{Key: "email", Value: email}}).Decode(&user)

	 if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the email %v\n", email)
		return nil,err
	}

	db.Disconect()

	if err != nil {
		panic("se ha producido un error: " + err.Error())
	}
	
	return user,nil

}


func RegisterUser(user *models.User) error{
	db := mongodb.Instance()
	coll := db.Client.Database("arrowfood").Collection("users");
	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	db.Disconect()
	return nil
}	

func DeleteUser(user *models.User) error{
	db := mongodb.Instance()
	coll := db.Client.Database("arrowfood").Collection("users");
	filter := bson.D{{Key: "email", Value: user.Email}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	db.Disconect()
	return nil
}

func GetUserByUsername(username string) (*models.User, error){

	db := mongodb.Instance()

	coll := db.Client.Database("arrowfood").Collection("users")
	filter := bson.D{{Key: "nombre", Value: username}}

	var user *models.User
	err := coll.FindOne(context.TODO(),filter).Decode(&user)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	db.Disconect()
	return user, nil

}