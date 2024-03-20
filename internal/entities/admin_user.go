package entities

type AdminUserCreate struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AdminUser Нужна ли абстракция?
type AdminUser struct {
	ID int `json:"id"`
	AdminUserCreate
}
