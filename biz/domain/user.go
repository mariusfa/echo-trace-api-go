package domain

type User struct {
	Id             int64
	Name           string
	HashedPassword string
	ApiToken       string
}
