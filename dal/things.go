package dal

import "gorm.io/gorm"

type Thing struct {
	gorm.Model
	Win        string `gorm:"not null"`
	Unhappy    string `gorm:"not null"`
	Lesson     string `gorm:"not null"`
	DayScore   uint   `gorm:"not null"`
	OtherScore uint   `json:"otherScore"`
	User       *uint  `gorm:"index, not null"`
}

func CreateThing(t *Thing) *gorm.DB {
	return DB.Create(t)
}

func FindThingsByUser(dest interface{}, userId interface{}) *gorm.DB  {
    return DB.Model(&Thing{}).Find(dest, "user = ?", userId)
    
}
