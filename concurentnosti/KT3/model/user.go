package model

type User struct {
	Id   int
	Name string
	Surname string
	Age  int

}

var Users []User

var NextId int

func GetAllUsers () (users []User, err error) {
	if len(Users) == 0 {
		Users = []User{
			{1, "Лев", "Млекопитающее", 4},
			{2, "Горилла", "Млекопетающее", 5},
			{3, "Змея", "Рептилия", 1},
			{4, "Орёл", "Птица", 2},
		}
	}
	return Users, nil
}

func AddUser (user *User) {
	user.Id = NextId
	NextId++

	Users = append(Users, *user)
}