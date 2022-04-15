package data

type Rating struct {
	Number   int64   `gorm:"column:number"`
	PlayerID int64   `gorm:"column:player_id"`
	Score    int64   `gorm:"column:score"`
	MMR      float64 `gorm:"column:mmr"`
}

func (Rating) TableName() string {
	return "rating"
}

func RatingIDMap(r Rating) int64 {
	return r.PlayerID
}
