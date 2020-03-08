package user

type Post struct {
	PostId   string  `json:"post_id",bson:"_id"`
	PostUp   float64 `json:"post_up",bson:"up_score"`
	PostDown float64 `json:"post_down",bson:"down_score"`
	Rel      float64 `json:"post_rel",bson:"_rel"`
}
