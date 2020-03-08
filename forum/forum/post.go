package forum

type UserPost struct {
	PostId   string  `json:"postId",bson:"_id"`
	PostUp   float64 `json:"postUp",bson:"up_score"`
	PostDown float64 `json:"postDown",bson:"down_score"`
	Rel      float64 `json:"postRel"`
}

type Post struct {
	PostId      string  `json:"post_id",bson:"_id"`
	PostContent string  `json:"content",bson:"_content"`
	UserId      string  `json:"userId"`
	PostUp      float64 `json:"post_up_score",bson:"up_score"`
	PostDown    float64 `json:"post_down_score",bson:"down_score"`
	Rel         float64 `json:"post_rel",bson:"_rel"`
}
