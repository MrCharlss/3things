package dal

import "gorm.io/gorm"

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	// Todos    []Todo `gorm:"foreignKey:User"`
}

func CreateUser(user *User) *gorm.DB {
	return DB.Create(user)
}

func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return DB.Model(&User{}).Take(dest, conds...)
}
func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email=?", email)
}
