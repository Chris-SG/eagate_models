package ddr_models

import (
	"github.com/chris-sg/eagate_models/user_models"
	"time"
)

type PlayerDetails struct {
	Code               int       `tag:"DDR-CODE_sougou" gorm:"column:code;primary_key"`
	Name               string    `tag:"ダンサーネーム_sougou" gorm:"column:name"`
	Prefecture         string    `tag:"所属都道府県_sougou" gorm:"column:location"`
	SingleRank         string    `tag:"段位(SINGLE)_sougou" gorm:"column:single_rank"`
	DoubleRank         string    `tag:"段位(DOUBLE)_sougou" gorm:"column:double_rank"`
	Affiliation        string    `tag:"所属クラス_sougou" gorm:"column:affiliation"`
	Playcount          int       `tag:"総プレー回数_sougou" gorm:"column:playcount"`
	LastPlayDate       time.Time `tag:"最終プレー日時_sougou" gorm:"column:last_play_date"`
	SinglePlaycount    int       `tag:"プレー回数_single" gorm:"column:single_playcount"`
	SingleLastPlayDate time.Time `tag:"最終プレー日時_single" gorm:"column:last_single_play_date"`
	DoublePlaycount    int       `tag:"プレー回数_double" gorm:"column:double_playcount"`
	DoubleLastPlayDate time.Time `tag:"最終プレー日時_double" gorm:"column:last_double_play_date"`

	EaGateUser user_models.User `gorm:"foreignkey:account_name"`
}

func (PlayerDetails) TableName() string {
	return "ddrPlayerDetails"
}


type Song struct {
	Id     string `gorm:"column:song_id;primary_key"`
	Name   string `gorm:"column:song_name"`
	Artist string `gorm:"column:song_artist"`
	Image  string `gorm:"column:song_image"`

	Difficulties []SongDifficulty `gorm:"foreignkey:SongId"`
}

func (Song) TableName() string {
	return "ddrSongs"
}

type SongDifficulty struct {
	SongId string `gorm:"column:song_id;primary_key"`
	DifficultyId int8 `gorm:"column:difficulty_id;primary_key"`
	DifficultyValue int8 `gorm:"column:difficulty_value"`
}

func (SongDifficulty) TableName() string {
	return "ddrSongDifficulties"
}

type SongStatistics struct {
	BestScore      int `tag:"ハイスコア" gorm:"column:score_record"`
	Lamp       int8		 `gorm:"column:clear_lamp"`
	PlayCount  int       `tag:"プレー回数" gorm:"column:playcount"`
	ClearCount int       `tag:"クリア回数" gorm:"column:clearcount"`
	MaxCombo   int       `tag:"最大コンボ数" gorm:"column:maxcombo"`
	LastPlayed time.Time `tag:"最終プレー時間" gorm:"column:lastplayed"`

	Chart SongDifficulty	`gorm:"primary_key"`
	Player PlayerDetails	`gorm:"primary_key"`
}

func (SongStatistics) TableName() string {
	return "ddrSongStatistics"
}

type Score struct {
	Score int `gorm:"column:score"`
	TimePlayed time.Time `gorm:"column:time_played;primary_key"`

	Chart SongDifficulty	`gorm:"primary_key"`
	Player PlayerDetails	`gorm:"primary_key"`
}