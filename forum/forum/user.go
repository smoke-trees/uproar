package forum

type User struct {
	UserId    string `json:"userId",bson:"_id"`
	Name      string `json:"name",bson:"name"`
	UserName  string `json:"username",bson:"username"`
	Phone     string `json:"phone",bson:"phone"`
	Email     string `json:"email",bson:"email"`
	Address1  string `json:"ad1",bson:"address_1"`
	Address2  string `json:"ad2",bson:"address_2"`
	City      string `json:"city",bson:"city"`
	State     string `json:"State",bson:"state"`
	UserLevel int8   `json:"level",bson:"user_level"`

	Cred    float64    `json:"cred",bson:"_cred"`
	RelUp   []UserPost `bson:"relUp"`
	RelDown []UserPost `bson:"relDown"`
	Posts   []UserPost `bson:"posts"`
}



