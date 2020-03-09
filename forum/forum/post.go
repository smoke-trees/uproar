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
	PostUp      float64 `json:"postUp_score",bson:"up_score"`
	PostDown    float64 `json:"postDown_score",bson:"down_score"`
	Rel         float64 `json:"postRel",bson:"_rel"`
}

type Posts []Post

func (j Posts) Len() int {
	return len([]Post(j))
}

func (p Posts) Less(i, j int) bool {
	return p[i].Rel > p[j].Rel
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
