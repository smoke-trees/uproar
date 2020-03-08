package database

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type AuthMongoClient struct {
	Client   *mongo.Client
	database string
}

func NewAuthDB(connectionString string, database string) (*AuthMongoClient, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &AuthMongoClient{
		Client:   client,
		database: database,
	}, nil
}

// AddUser adds a user to the database
func (mc *AuthMongoClient) AddUser(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$or",
		bson.A{bson.D{{"username", user.Username}}, bson.D{{"email", user.Email}}}}}

	ctx := context.Background()

	one := collection.FindOne(ctx, filter)
	if one.Err() == nil {
		return errors.New("user already exists")
	}

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return errors.New("user account can't be created")
	}
	return nil
}

// RemoveUser removes a user from the database
func (mc *AuthMongoClient) RemoveUser(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$or",
		bson.A{bson.D{{"username", user.Username}}, bson.D{{"email", user.Email}}}}}

	ctx := context.Background()

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("user account can't be deleted")
	}

	return nil
}

// Authenticate authenticates the user
func (mc *AuthMongoClient) Authenticate(user User) (AuthResponse, error) {
	response := AuthResponse{
		Status: Success,
		JWT:    "",
	}

	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$or",
		bson.A{bson.D{{"username", user.Username}}, bson.D{{"email", user.Email}}}}}

	ctx := context.Background()

	one := collection.FindOne(ctx, filter)
	if one.Err() != nil {
		response.Status = WrongUsername
		return response, errors.New("can not find the username or email")
	}

	var readUser User
	err := one.Decode(&readUser)
	if err != nil {
		response.Status = DBError
		log.Errorf("Couldn't parse data for following user:%v", user)
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(user.Password))
	if err != nil {
		response.Status = WrongPassword
		return response, errors.New("Wrong Password")
	}

	return response, nil
}

func (mc *AuthMongoClient) Disconnect() error {
	err := mc.Client.Disconnect(context.Background())
	if err != nil {
		log.Errorf("Error in disconnecting to database")
		return err
	}
	return nil
}
