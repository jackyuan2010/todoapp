package model

type BaseModel struct{
	id string `gorm:"primaryKey" json:"id"`
	create_at int64 `gorm:"autoCreateTime:nano" json:"create_at"`
	update_at int64 `gorm:"autoUpdateTime:nano" json:"update_at"`
}