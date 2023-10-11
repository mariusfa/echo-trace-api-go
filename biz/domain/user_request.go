package domain

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
