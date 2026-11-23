package user

type User struct {
	ID    int
	Name  string
	Email string
}

var users []User //slice yo store users

func CreateUser(id int, name string, email string) User {
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	users = append(users, user)

	return user
}
