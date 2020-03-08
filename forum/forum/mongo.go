package forum

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ForumMongoClient struct {
	Client   *mongo.Client
	database string
}

func (mc *ForumMongoClient) AddPostUpVote(p Post, u User) error {
	mc.RemovePostUpVote(p, u)
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"relUp", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) RemovePostUpVote(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relUp", bson.D{{"postid", p.PostId}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) AddPostDownVote(p Post, u User) error {
	mc.RemovePostDownVote(p, u)
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"relDown", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) RemovePostDownVote(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relDown", bson.D{{"postid", p}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) AddPost(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"posts", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) RemovePost(p Post, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"posts", bson.D{{"postid", p}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) GetUserFromUserId(id string) (User, error) {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", id}}

	var user User

	one := collection.FindOne(context.Background(), filter)
	if one.Err() != nil {
		log.Warn("No forum found for username:", user.UserId)
		return user, one.Err()
	}

	err := one.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (mc *ForumMongoClient) NewUserRegister(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$or",
		bson.A{bson.D{{"email", user.Email}}, bson.D{{"username", user.UserName}}},
	}}
	one := collection.FindOne(context.Background(), filter)
	if one.Err() == nil {
		log.Error("User already exist")
		return errors.New("forum already exists")
	}

	h := sha256.New()
	h.Write([]byte(user.UserName))
	user.UserId = hex.EncodeToString(h.Sum(nil))

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (mc *ForumMongoClient) UpdateUserCredibility(user User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", user.UserId}}
	update := bson.D{{"_cred", user.Cred}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Error("Error in updating reliablity")
		return err
	}
	return nil
}

func NewForumDB(connectionString string, database string) (*ForumMongoClient, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &ForumMongoClient{
		Client:   client,
		database: database,
	}, nil
}

func (mc *ForumMongoClient) Disconnect() {
	mc.Client.Disconnect(context.Background())
}
