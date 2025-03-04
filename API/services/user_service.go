package services

import (
	"API/database"
	"API/models"
	"errors"
)

// Create a new user
func CreateUser(name, email string) models.User {
	user := models.User{Name: name, Email: email}
	database.DB.Create(&user)
	return user
}

// Get all users
func GetUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
func UpdateUser(id uint, updatedUser models.User) (*models.User, error) {
	var user models.User

	// Find the user by ID
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	// Update user fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	// Save changes
	database.DB.Save(&user)

	return &user, nil
}
func DeleteUser(id uint) error {
	var user models.User

	// Check if user exists
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return errors.New("user not found")
	}

	// Delete the user
	database.DB.Delete(&user)
	return nil
}
