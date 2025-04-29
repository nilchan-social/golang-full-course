package user

// public
type User struct {
	// public
	Name string

	// private (because of lower case of first letter)
	age int
}

// public
func (u *User) ChangeAge(newAge int) {
	u.age = newAge
}

// private (because of lower case of first letter)
func (u *User) changeName(newName string) {
	u.Name = newName
}
