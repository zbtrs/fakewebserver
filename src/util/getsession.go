package util

import "math/rand"

func GetSession() string {
	var session = ""
	for i := 0; i < 10; i++ {
		x := rand.Intn(10)
		session += string(x)
	}
	return session
}
