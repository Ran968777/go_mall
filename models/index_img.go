package models

import "time"

// IndexImg 主页轮播图
type IndexImg struct {
	ImgID      uint64    `json:"-" gorm:"primaryKey;column:img_id"`
	ShopID     uint64    `json:"-" gorm:"column:shop_id"`
	ImgURL     string    `json:"imgUrl" gorm:"column:img_url;not null"`
	Des        string    `json:"-" gorm:"column:des;default:''"`
	Title      string    `json:"-" gorm:"column:title"`
	Link       string    `json:"-" gorm:"column:link"`
	Status     int8      `json:"-" gorm:"column:status;default:0"`
	Seq        int       `json:"seq" gorm:"column:seq;default:0"`
	UploadTime time.Time `json:"uploadTime" gorm:"column:upload_time"`
	Relation   uint64    `json:"relation" gorm:"column:relation"`
	Type       int       `json:"type" gorm:"column:type"`
}

// TableName 指定表名
func (IndexImg) TableName() string {
	return "tz_index_img"
}
