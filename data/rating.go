package data

type Rating struct {
	Number int64   `gorm:"column:number"`
	SpaID  int64   `gorm:"column:player_id"`
	Score  int64   `gorm:"column:score"`
	MMR    float64 `gorm:"column:mmr"`
}

func (Rating) TableName() string {
	return "rating"
}
