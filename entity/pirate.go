package entity

//Pirate represents pirates table in database
type Pirate struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	DevilFruit string `gorm:"type:varchar(255)" json:"devil_fruit"`
	Crew       string `gorm:"type:varchar(255)" json:"crew"`
	Job        string `gorm:"type:varchar(255)" json:"job"`
	UserID     uint64 `gorm:"not null" json="-"`
	User       User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}
