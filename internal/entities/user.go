package entities

type UserBase struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type User struct {
	UserBase
	ID int `json:"id"`
}
