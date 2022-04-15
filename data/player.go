package data

type Player struct {
	ID       int64  `gorm:"column:id"`
	Nickname string `gorm:"column:nickname"`
	Battles  int64  `gorm:"column:battles"`
	Damage   int64  `gorm:"column:damage"`
	Wins     int64  `gorm:"column:wins"`
}

func (Player) TableName() string {
	return "players"
}
