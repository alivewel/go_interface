package users

type User struct {
	ID   int
	Name string
	Age  int
}

type Storage interface {
	Users() ([]User, error)
	UsersByAge(age int) ([]User, error)
	User(id int) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int) error
	// другие методы...
}
