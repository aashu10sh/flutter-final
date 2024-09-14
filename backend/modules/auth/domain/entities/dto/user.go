package auth_entities

type User struct {
	Id       string
	Username string
	Name     string
}

type UserWithPassword struct {
	Id       string
	Name     string
	Username string
	Password string
}

func (u *User) NewUser(name string) User {
	return User{
		Name: name,
	}
}

func (u *User) NewUserWithId(Id string) {
	u.Id = Id
}
