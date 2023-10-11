package domain




// type User struct {
// 	Id int64
// 	Name string
// 	HashedPassword string
// 	ApiToken string
// }
type UserRequest struct {
	Username string
	Password string
}

func (ur *UserRequest) ToUser() User {
	return User{
		Name: ur.Username,
		HashedPassword: ur.Password,
	}
}
