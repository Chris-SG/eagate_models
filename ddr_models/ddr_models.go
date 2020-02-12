package ddr_models

type Song struct {
	Id     string `gorm:"column:song_id;primary_key"`
	Name   string `gorm:"column:song_name"`
	Artist string `gorm:"column:song_artist"`
	Image  string `gorm:"column:song_image"`
}

func (Song) TableName() string {
	return "ddrSongs"
}

type SongDifficulty struct {
	Song Song `gorm:"foreignkey:song_id;association_foreignkey:song_id"`
	SongId string `gorm:"column:song_id;primary_key"`
	DifficultyId int8 `gorm:"column:difficulty_id;primary_key"`
	DifficultyValue int8 `gorm:"column:difficulty_value"`
}

func (SongDifficulty) TableName() string {
	return "ddrSongDifficulties"
}