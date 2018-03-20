package users

import "time"

type Users struct {
	Id  	   int64 `xorm:"pk"`
	Name       string
	Password   string
	Age        int64
	Sex        string
	CreateTime time.Time
	EndTime    time.Time
}

func (u Users)AddUsers(sex, name ,password string, age int64) []*Users{
	users := make([]*Users,0)
	user := &Users{
		Name      : name,
		Password  : password,
		Age       : age,
		Sex       : sex,
		CreateTime: time.Now(),
		EndTime   :    time.Now(),
	}
	users = append(users,user)
	return users
}

