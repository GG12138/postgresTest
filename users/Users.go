package users

type Users struct {
	Id int64 `xorm:"pk"`
	Name string `xorm:"unique"`
	Age int64
	Sex string
}

