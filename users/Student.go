package users

type Student struct {
	Id int64 `xorm:"pk"`
	Name string
	Email string
}
