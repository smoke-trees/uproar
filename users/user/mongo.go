package user

import (
	"context"
	"crypto/sha256"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoClient struct {
	Client   *mongo.Client
	database string
}

func (mc *UserMongoClient) AddPostUpVote(p Post, u User) error {
	mc.RemovePostUpVote(p, u)
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"relUp", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) RemovePostUpVote(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relUp", bson.D{{"_id", p.PostId}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) AddPostDownVote(p Post, u User) error {
	mc.RemovePostDownVote(p, u)
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"relDown", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) RemovePostDownVote(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relDown", bson.D{{"_id", p}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) AddPost(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"posts", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) RemovePost(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"posts", bson.D{{"_id", p}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *UserMongoClient) GetUserFromUserId(id string) (User, error) {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", id}}

	var user User

	one := collection.FindOne(context.Background(), filter)
	if one.Err() != nil {
		log.Warn("No user found for username:", user.UserName)
		return user, one.Err()
	}

	err := one.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (mc *UserMongoClient) NewUserRegister(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$or",
		bson.A{bson.D{{"email", user.Email}}, bson.D{{"username", user.UserName}}},
	}}
	one := collection.FindOne(context.Background(), filter)
	if one.Err() == nil {
		log.Error("User already exist")
		return errors.New("user already exists")
	}

	h := sha256.New()
	h.Write([]byte(user.UserName))
	user.UserId = string(h.Sum(nil))

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (mc *UserMongoClient) UpdateUserCredibility(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"_id", user.UserId}}
	update := bson.D{{"_cred", user.Cred}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Error("Error in updating reliablity")
		return err
	}
	return nil
}

func NewUserDB(connectionString string, database string) (*UserMongoClient, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &UserMongoClient{
		Client:   client,
		database: database,
	}, nil
}

func (mc *UserMongoClient) Disconnect() {
	mc.Client.Disconnect(context.Background())
}
