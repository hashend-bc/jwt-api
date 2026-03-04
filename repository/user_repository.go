package repository

import "user-auth-api/models"

var users = make(map[string]models.User)
var idCounter = 1

func CreateUser(username, password string) models.User {
	user := models.User{
		ID:       idCounter,
		Username: username,
		PasswordHash: password,
	}

	users[username] = user
	idCounter++

	return user
}

func GetUserByUsername(username string) (models.User, bool) {
	user, exists := users[username]
	return user, exists
}