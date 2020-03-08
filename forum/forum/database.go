package forum

type Database interface {
	GetUserFromUserId(string) (User, error)
	NewUserRegister(User) error
	AddPostUpVote(UserPost, User) error
	AddPostDownVote(UserPost, User) error
	AddUserPost(UserPost, User) error
	RemovePost(UserPost, User) error
	UpdateUserCredibility(User) error
	GetPostFromPostId(string) (Post, error)
	UpdatePostAfterAction(Post) error
	NewPost(Post) error
	Disconnect()
	GetUserFromUserName(username string) (User, interface{})
}
