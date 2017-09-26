package users

import "time"

type Users struct {
	Id int64
	Name string
	Age int64
	Sex string
	CreateTime time.Time
}

func (u Users)AddUsers(sex, name string, age int64) []*Users{
	users := make([]*Users,0)
	user := &Users{
		Name  : name,
		Age   : age,
		Sex   : sex,
		CreateTime: time.Now(),
	}
	users = append(users,user)
	return users
}

