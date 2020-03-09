package forum

import (
	"context"
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

func (mc *ForumMongoClient) IsUserAction(u User, p Post) bool {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"$and",
		bson.A{bson.D{{"$or", bson.A{bson.D{{"relUp.postid", p.PostId}},
			bson.D{{"relDown.postid", p.PostId}}}}},
			bson.D{{"userid", u.UserId}}}}}
	one := collection.FindOne(context.Background(), filter)
	if one.Err() != nil {
		return false
	}
	return true
}

func (mc *ForumMongoClient) GetAllPosts() ([]Post, error) {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("post_data")

	res, err := collection.Find(context.Background(), bson.D{{}}, )
	if err != nil {
		log.Error(err)
	}
	var posts []Post
	err = res.All(context.Background(), &posts)
	return posts, nil
}

func (mc *ForumMongoClient) GetUserFromUserName(username string) (User, error) {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"username", username}}

	var user User

	one := collection.FindOne(context.Background(), filter)
	if one.Err() != nil {
		log.Warn("No forum found for username:", username)
		return user, one.Err()
	}

	err := one.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (mc *ForumMongoClient) GetPostFromPostId(p string) (Post, error) {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("post_data")

	filter := bson.D{{"postid", p}}
	one := collection.FindOne(context.Background(), filter)
	if one.Err() != nil {
		return Post{}, one.Err()
	}

	var post Post

	err := one.Decode(&post)
	if err != nil {
		return Post{}, one.Err()
	}
	return post, nil
}

func (mc *ForumMongoClient) UpdatePostAfterAction(p Post) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("post_data")

	filter := bson.D{{"postid", p.PostId}}
	update := bson.D{{"$set", bson.D{{"rel", p.Rel}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (mc *ForumMongoClient) NewPost(p Post) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("post_data")

	_, err := collection.InsertOne(context.Background(), p)
	return err

}

func (mc *ForumMongoClient) AddPostUpVote(p UserPost, u User) error {
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

func (mc *ForumMongoClient) RemovePostUpVote(p UserPost, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relUp", bson.D{{"postid", p.PostId}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) AddPostDownVote(p UserPost, u User) error {
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

func (mc *ForumMongoClient) RemovePostDownVote(p UserPost, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$pull",
		bson.D{{"relDown", bson.D{{"postid", p}},
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) AddUserPost(p UserPost, u User) error {
	database := mc.Client.Database(mc.database)
	collection := database.Collection("user_data")

	filter := bson.D{{"userid", u.UserId}}
	update := bson.D{{"$push",
		bson.D{{"posts", p,
		}}}}
	collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (mc *ForumMongoClient) RemovePost(p UserPost, u User) error {
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
	user.RelUp = []UserPost{}
	user.RelDown = []UserPost{}
	user.Posts = []UserPost{}

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
	update := bson.D{{"$set", bson.D{{"cred", user.Cred}}}}

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
