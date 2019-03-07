package repo

type AppRepository interface {
	FindByID(id string) (UserDetail, error)
	FindByEmail(email string) (UserDetail, error)
	InsertNewUser(user UserDetail) (bool, error)
	UpdatePassword(id string, newPassword string) (bool, error)
	GetGenres() ([]GenreList, error)
	GetInstruments() ([]InstrumentList, error)
	InsertProfile(name string, gender string, birthdate string, bio string, avatarurl string, id string, genre UserGenre, instrument UserInstrument) (bool, error)
	UpdateProfile(name string, gender string, birthdate string, bio string, avatarurl string, id string, genre string, instrument string) (bool, error)
	InsertNewPost(newPost UserPost) (bool, error)
	DeletePost(postId string) (bool, error)
	InsertFollow(userFollow UserFollow) (bool, error)
	GetFollower(id string) ([]UserDetail, error)
	GetFollowing(id string) ([]UserDetail, error)
	GetUserPost(id string) ([]UserPost, error)
	GetUserProfile(id string) (UserDetail, error)
	GetFollowedId(id string) (string, error)
	GetRelatedPost(id string) ([]UserPost, error)
}