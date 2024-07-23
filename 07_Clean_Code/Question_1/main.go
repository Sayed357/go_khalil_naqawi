package main

import (
	"errors"
	"fmt"
)

type User struct {
	Email    string
	Password string
}

// Renaming User Repository to UserRepo to follow naming conventions for types.
type UserRepo struct {
	DB []User
}

// Register method should return an error if registration fails to handle it.
func (r *UserRepo) Register(u User) error {

	// Check for empty email or password and return an error.
	if u.Email == "" || u.Password == "" {
		return errors.New("registration failed: email or password cannot be empty")
	}

	r.DB = append(r.DB, u)
	return nil
}

// Login method should return an error or a token and an error if login fails.
func (r *UserRepo) Login(u User) (string, error) {

	if u.Email == "" || u.Password == "" {
		return "", errors.New("login failed: email or password cannot be empty")
	}

	// Check if the user exists in the database.
	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token", nil
		}
	}

	// Return an error if the user is not found.
	return "", errors.New("login failed: invalid email or password")
}

func main() {

	repo := UserRepo{}

	// User registration
	err := repo.Register(User{Email: "khalil.naqawi@gmail.com", Password: "codingcourse2024"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Attempt to log in with the registered user.
	token, err := repo.Login(User{Email: "khalil.naqawi@gmail.com", Password: "codingcourse2024"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Login successful, token:", token)
}
