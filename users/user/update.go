package user

func updateRelAndUserOnUp(user User, post Post) User {
	user.RelUp = append(user.RelUp, post)

	return user
}

func updateRelAndUserOnDown(user User, post Post) User {
	user.RelDown = append(user.RelDown, post)

	return user
}

func updateCred(user User) User {
	for _, relUp := range user.RelUp {
		user.Cred += relUp.Rel
	}

	for _, relDown := range user.RelDown {
		user.Cred -= relDown.Rel
	}

	for _, post := range user.Posts {
		user.Cred += post.Rel
	}

	return user
}
