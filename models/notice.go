package models

import (
	"encoding/json"
	"time"
)

/*
create table tz_notice
(

	id           bigint auto_increment comment '公告id'
	    primary key,
	shop_id      bigint                              null comment '店铺id',
	title        varchar(36)                         null comment '公告标题',
	content      text                                null comment '公告内容',
	status       tinyint(1)                          null comment '状态(1:公布 0:撤回)',
	is_top       tinyint                             null comment '是否置顶',
	publish_time timestamp                           null comment '发布时间',
	update_time  timestamp default CURRENT_TIMESTAMP null comment '更新时间'

)

	charset = utf8;*
*/

// Notice 公告
type Notice struct {
	ID          uint64    `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	ShopID      uint64    `json:"shopId" gorm:"column:shop_id"`
	Title       string    `json:"title" gorm:"column:title;size:36"`
	Content     string    `json:"content" gorm:"column:content;type:text"`
	Status      int8      `json:"status" gorm:"column:status;comment:状态(1:公布 0:撤回)"`
	IsTop       int8      `json:"isTop" gorm:"column:is_top;comment:是否置顶"`
	PublishTime time.Time `json:"publishTime" gorm:"column:publish_time"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName 指定表名
func (Notice) TableName() string {
	return "tz_notice"
}

// MarshalJSON 自定义JSON序列化
func (n Notice) MarshalJSON() ([]byte, error) {
	type Alias Notice
	return json.Marshal(&struct {
		Alias
		PublishTime string `json:"publishTime"`
		UpdateTime  string `json:"updateTime"`
	}{
		Alias:       Alias(n),
		PublishTime: n.PublishTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  n.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}
