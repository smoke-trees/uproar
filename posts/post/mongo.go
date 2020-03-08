package post

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostMongoClient struct {
	Client   *mongo.Client
	database string
}

func (p PostMongoClient) GetPostFromPostId(id string) (Post, error) {
	database := p.Client.Database(p.database)
	collection := database.Collection("post_data")

	filter := bson.D{{Key:"_id", Value:id}}

	var post Post

	val := collection.FindOne(context.Background(), filter)

	if val.Err() != nil {
		return post, val.Err()
	}

	err := val.Decode(&post)
	if err != nil{
		return post, err
	}

	return post, nil
}

func (p PostMongoClient) UpdatePostAfterAction(Post) (error) {
	panic("implement me")
}

func (p PostMongoClient) NewPost(Post) (error) {
	panic("implement me")
}

func (p PostMongoClient) Disconnect() {
	panic("implement me")
}

func NewPostDB(connectionString string, database string) (*PostMongoClient, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &PostMongoClient{
		Client:   client,
		database: database,
	}, nil
}

