package repo

type AppRepository interface {
	FindByID(id string) (UserDetail, error)
	FindByEmail(email string) (UserDetail, error)
	FindByUsername(username string) (UserDetail, error)
	InsertNewUser(user UserDetail) (bool, error)
	UpdatePassword(id string, newPassword string) (bool, error)
}