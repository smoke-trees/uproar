package post

type PostUser struct {
	UserId string `json:"user_id"`
	Cred float64 `json:"cred"`
}

type Post struct {
	PostId string `json:"post_id",bson:"_id"`
	PostContent string `json:"content",bson:"_content"`
	PostUp float64 `json:"post_up_score",bson:"up_score"`
	PostDown float64 `json:"post_down_score",bson:"down_score"`
	Rel float64 `json:"post_rel",bson:"_rel"`
}
