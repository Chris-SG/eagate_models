package drs_models

import "time"

type PlayerDetails struct {
	Code int    `gorm:"column:code;primary_key"`
	Name string `gorm:"column:name"`

	EaGateUser *string `gorm:"column:eagate_user"`
}

func (PlayerDetails) TableName() string {
	return "drsPlayerDetails"
}

type PlayerProfileSnapshot struct {
	PlayCount   int `gorm:"column:play_count;primary_key"`
	PlaySeconds int `gorm:"column:play_seconds"`
	TotalStars  int `gorm:"column:total_stars"`
	UsedStars   int `gorm:"column:used_stars"`
	LastPlayed  time.Time `gorm:"column:last_played"`

	PlayerCode int `gorm:"column:player_code;primary_key"`
}

func (PlayerProfileSnapshot) TableName() string {
	return "drsPlayerProfileSnapshots"
}

type Song struct {
	SongId         string `gorm:"column:song_id;primary_key"`
	SongName       string `gorm:"column:name"`
	ArtistName     string `gorm:"column:artist"`
	MaxBpm         int    `gorm:"column:max_bpm"`
	MinBpm         int    `gorm:"column:min_bpm"`
	LimitationType int    `gorm:"column:limitation_type"`
	Genre          int    `gorm:"column:genre"`
	VideoFlags     int    `gorm:"column:video_flags"`
	License        string `gorm:"column:license"`
}

func (Song) TableName() string {
	return "drsSongs"
}

type Difficulty struct {
	Mode       string `gorm:"column:mode"`
	Difficulty string `gorm:"column:difficulty"`
	Level      int    `gorm:"column:level"`

	SongId string `gorm:"column:song_id;primary_key"`
}

func (Difficulty) TableName() string {
	return "drsDifficulties"
}

type PlayerSongStats struct {
	BestScore         int       `gorm:"column:best_score"`
	Combo             int       `gorm:"column:combo"`
	PlayCount         int       `gorm:"column:play_count"`
	Param             int       `gorm:"column:param"`
	BestScoreDateTime time.Time `gorm:"column:best_score_time"`
	LastPlayDateTime  time.Time `gorm:"column:last_play_time"`

	P1Code     int `gorm:"column:p1_code"`
	P1Score    int `gorm:"column:p1_score"`
	P1Perfects int `gorm:"column:p1_perfects"`
	P1Greats   int `gorm:"column:p1_greats"`
	P1Goods    int `gorm:"column:p1_goods"`
	P1Bads     int `gorm:"column:p1_bads"`

	P2Code     *int `gorm:"column:p2_code"`
	P2Score    *int `gorm:"column:p2_score"`
	P2Perfects *int `gorm:"column:p2_perfects"`
	P2Greats   *int `gorm:"column:p2_greats"`
	P2Goods    *int `gorm:"column:p2_goods"`
	P2Bads     *int `gorm:"column:p2_bads"`

	PlayerCode int    `gorm:"column:player_code;primary_key"`
	SongId     string `gorm:"column:song_id;primary_key"`
	Mode       string `gorm:"column:mode;primary_key"`
	Difficulty string `gorm:"column:difficulty;primary_key"`
}

func (PlayerSongStats) TableName() string {
	return "drsPlayerSongStats"
}

func (s1 PlayerSongStats) Equals(s2 PlayerSongStats) bool {
	return s1.SongId == s2.SongId &&
		s1.Mode == s2.Mode &&
		s1.Difficulty == s2.Difficulty &&
		s1.PlayerCode == s2.PlayerCode &&
		s1.BestScoreDateTime.String() == s2.BestScoreDateTime.String() &&
		s1.LastPlayDateTime.String() == s2.LastPlayDateTime.String() &&
		s1.BestScore == s2.BestScore &&
		s1.Combo == s2.Combo &&
		s1.PlayCount == s2.PlayCount
}

type PlayerScore struct {
	Shop     string    `gorm:"column:shop"`
	Score    int       `gorm:"column:score"`
	MaxCombo int       `gorm:"column:max_combo"`
	Param    int       `gorm:"column:param"`
	PlayTime time.Time `gorm:"column:play_time"`

	P1Code     int `gorm:"column:p1_code"`
	P1Score    int `gorm:"column:p1_score"`
	P1Perfects int `gorm:"column:p1_perfects"`
	P1Greats   int `gorm:"column:p1_greats"`
	P1Goods    int `gorm:"column:p1_goods"`
	P1Bads     int `gorm:"column:p1_bads"`

	P2Code     *int `gorm:"column:p2_code"`
	P2Score    *int `gorm:"column:p2_score"`
	P2Perfects *int `gorm:"column:p2_perfects"`
	P2Greats   *int `gorm:"column:p2_greats"`
	P2Goods    *int `gorm:"column:p2_goods"`
	P2Bads     *int `gorm:"column:p2_bads"`

	VideoUrl *string `gorm:"column:video_url"`

	PlayerCode int    `gorm:"column:player_code;primary_key"`
	SongId     string `gorm:"column:song_id;primary_key"`
	Mode       string `gorm:"column:mode;primary_key"`
	Difficulty string `gorm:"column:difficulty;primary_key"`
}

func (PlayerScore) TableName() string {
	return "drsPlayerScores"
}
