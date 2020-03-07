package user

type user struct {
	userId string
	password string
	relUp []post
	relDown []post
	cred float64
	posts []post
}

type post struct {
	postId string
	postUp float64
	postDown float64
	rel float64
}

func updateRelAndUserOnUp(user user,post post) (user, post) {

	post.rel = post.postUp - post.postDown

	user.relUp = append(user.relUp, post)

	return user, post
}

func updateRelAndUserOnDown(user user,post post) (user, post) {

	post.rel = post.postUp - post.postDown

	user.relDown = append(user.relDown, post)

	return user, post
}
