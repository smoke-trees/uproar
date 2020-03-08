package user

type Database interface {
	GetUserFromUserId(string) (User, error)
	NewUserRegister(User) error
	AddPostUpVote(Post, User) error
	AddPostDownVote(Post, User) error
	AddPost(Post, User) error
	RemovePost(Post, User) error
	UpdateUserCredibility(User) error
	Disconnect()
}
