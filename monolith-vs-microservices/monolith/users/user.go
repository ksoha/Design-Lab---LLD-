package users

type User struct {
	ID    int
	Name  string
	Email string
}

var users []User //slice yo store users

func CreateUser(name string, email string) User {
	user := User{
		ID:    len(users) + 1,
		Name:  name,
		Email: email,
	}

	users = append(users, user)

	return user
}
