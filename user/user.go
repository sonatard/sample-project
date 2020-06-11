package domain

type User struct {
	FirstName string
	LastName  string
}

func NewUser(first, last string) *User {
	return &User{
		FirstName: first,
		LastName:  last,
	}
}

func (u *User) Name() string {
	return u.LastName + " " +u.FirstName
}
