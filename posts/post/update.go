package post

func UpdatePostOnUp(Post Post, User PostUser) Post{
	Post.PostUp += User.Cred

	Post.Rel = Post.PostUp - Post.PostDown

	return Post
}

func UpdatePostOnDown(Post Post, User PostUser) Post{
	Post.PostDown += User.Cred

	Post.Rel = Post.PostUp - Post.PostDown

	return Post
}
