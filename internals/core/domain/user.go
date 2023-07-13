package domain

type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null" json:"-"`
}

func CreateUser(id int, email string, name string, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: password,
	}
}

func (u *User) GetEmail() string {
	return u.Email
}
