package obj

type User struct {
	id string
	name string
}

func NewUser(id string) *User {
	return &User{id: id}
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) SetName(v string) {
	u.name = v
}