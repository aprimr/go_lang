package store

import "learn-jwt/models"

var Users = map[string]models.User{}

func AddUser(user models.User) {
	Users[user.Id] = user
}

func GetUserByid(id string) (models.User, bool) {
	user, exists := Users[id]
	return user, exists
}

func GetUserByEmail(email string) (models.User, bool) {
	for _, user := range Users {
		if user.Email == email {
			return user, true
		}
	}
	return models.User{}, false
}
