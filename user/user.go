package user

var Name = "user"

type User struct {
	Name string
	age int
}

func (u User) Hello() string {
    return "Hello " + u.Name
}

func (u User) GetAge() int { 
	return u.age
}

func (u *User) SetAge(age int) { 
	u.age = age
}
