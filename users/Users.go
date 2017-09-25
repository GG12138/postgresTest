package users

type Users struct {
	Id int64
	Name string
	Age int64
	Sex string
}
func NewUser(){

}

func(u Users) AddUsers(sex, name string, id, age int64) *Users{
	users := new(Users)
	users.Id = id
	users.Name = name
	users.Age = age
	users.Sex = sex
	return users
}

