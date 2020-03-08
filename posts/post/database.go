package post

type Database interface {
	GetPostFromPostId(string)(Post, error)
	UpdatePostAfterAction(Post)(error)
	NewPost(Post)(error)
	Disconnect()
}





