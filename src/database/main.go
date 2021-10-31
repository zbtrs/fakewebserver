package database

type Database struct {
	Users  map[string]User
	UserId int
}

var MyDatabase Database
