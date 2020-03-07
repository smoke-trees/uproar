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

func updateRelAndUserOnUp(user user,post post) user {
	user.relUp = append(user.relUp, post)

	return user
}

func updateRelAndUserOnDown(user user,post post) user {
	user.relDown = append(user.relDown, post)

	return user
}

func updateCred (user user) user{
	for _, relUp := range(user.relUp){
		user.cred += relUp.rel
	}

	for _, relDown := range(user.relDown){
		user.cred -= relDown.rel
	}

	for _, post := range(user.posts){
		user.cred += post.rel
	}

	return user
}
