package user

import "strconv"

// User ...
type User struct {
	ID      string `json:"Id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

func Get(id int, name string, surname string, age int) User {
	return User{ID: strconv.Itoa(id), Name: name, Surname: surname, Age: strconv.Itoa(age)}
}
