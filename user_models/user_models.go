package user_models

type User struct {
	Name       string `gorm:"column:account_name;primary_key"`
	NickName   string `gorm:"column:community_name"`
	Cookie     string `gorm:"column:login_cookie"`
	Expiration int64  `gorm:"column:cookie_expiration"`
}

func (User) TableName() string {
	return "eaGateUser"
}
