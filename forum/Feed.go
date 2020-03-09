package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/smoke-trees/uproar/forum/forum"
	"net/http"
	"sort"
	"sync"
)

type FeedResponse struct {
	Status int
	data   []PostElement
}

type PostElement struct {
	PostId      string
	PostContent string
	UserAction  bool
	OpUsername  string
}

var Posts forum.Posts
var postsMutex sync.Mutex

func GetPosts() {
	postsMutex.Lock()
	defer postsMutex.Unlock()
	Posts, _ = s.Database.GetAllPosts()
	sort.Sort(Posts)
}

func GetFeedHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var p forum.Posts
	uid := request.URL.Query().Get("username")
	user, _ := s.Database.GetUserFromUserName(uid)
	if len(Posts) > 100 {
		p = Posts[0:100]
	} else {
		p = Posts
	}

	data := make([]PostElement, 100)

	for i, post := range p {
		data[i].PostId = post.PostId
		data[i].PostContent = post.PostContent
		data[i].UserAction = s.Database.IsUserAction(user, post)
		op, _ := s.Database.GetUserFromUserId(post.UserId)
		data[i].OpUsername = op.UserName
	}

	js, _ := json.Marshal(data)
	writer.Write(js)
}
