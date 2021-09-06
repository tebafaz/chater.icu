package database

import "chat/models"

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := postgres.Get(&user, "SELECT id, username, COALESCE(password, '') AS password FROM users WHERE username=$1", username)
	return user, err
}

func AddUserGetID(user models.User) (int, error) {
	var userID models.ID
	if user.Password == "" {
		err := postgres.Get(&userID, "INSERT INTO users(username, password) VALUES($1, NULL) RETURNING id", user.Username)
		return userID.ID, err
	}
	err := postgres.Get(&userID, "INSERT INTO users(username, password) VALUES($1, $2) RETURNING id", user.Username, user.Password)
	return userID.ID, err
}
