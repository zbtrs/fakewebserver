package database

type User struct {
	UserName     string
	UserPassword string
	UserId       int
}

func AddUser(userName string, userPassword string) {
	MyDatabase.UserId++
	NewUser := new(User)
	NewUser.UserName = userName
	NewUser.UserPassword = userPassword
	NewUser.UserId = MyDatabase.UserId
	MyDatabase.Users[userName] = *NewUser
}
