package user_models

type User struct {
	Name string `gorm:"column:account_name;primary_key"`
	Cookie string `gorm:"column:login_cookie"`
	Expiration int64 `gorm:"column:cookie_expiration"`
}

func (User) TableName() string {
	return "eaGateUser"
}