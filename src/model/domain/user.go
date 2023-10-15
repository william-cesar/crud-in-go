package domain

type tUser struct {
	id       string
	name     string
	email    string
	password string
	age      int8
	active   bool
}

func (u *tUser) GetId() string {
	return u.id
}

func (u *tUser) GetName() string {
	return u.name
}

func (u *tUser) GetEmail() string {
	return u.email
}

func (u *tUser) GetPassword() string {
	return u.password
}

func (u *tUser) GetAge() int8 {
	return u.age
}

func (u *tUser) IsActive() bool {
	return u.active
}

func (u *tUser) SetId(id string) {
	u.id = id
}

func (u *tUser) SetName(name string) {
	u.name = name
}

func (u *tUser) SetEmail(email string) {
	u.email = email
}

func (u *tUser) SetPassword(password string) {
	u.password = password
}

func (u *tUser) SetAge(age int8) {
	u.age = age
}

func (u *tUser) SetActive(active bool) {
	u.active = active
}
