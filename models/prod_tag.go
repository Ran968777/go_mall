package models

import "time"

/**
-- auto-generated definition
create table tz_prod_tag
(
    id          bigint auto_increment comment '分组标签id'
        primary key,
    title       varchar(36) null comment '分组标题',
    shop_id     bigint      null comment '店铺Id',
    status      tinyint(1)  null comment '状态(1为正常,0为删除)',
    is_default  tinyint(1)  null comment '默认类型(0:商家自定义,1:系统默认)',
    prod_count  bigint      null comment '商品数量',
    style       int         null comment '列表样式(0:一列一个,1:一列两个,2:一列三个)',
    seq         int         null comment '排序',
    create_time timestamp   null comment '创建时间',
    update_time timestamp   null comment '修改时间',
    delete_time timestamp   null comment '删除时间'
)
    comment '商品分组表' charset = utf8;


*/

// ProdTag 商品分组
type ProdTag struct {
	ID         uint64    `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Title      string    `json:"title" gorm:"column:title;size:36"`
	ShopID     uint64    `json:"shopId" gorm:"column:shop_id"`
	Status     int8      `json:"status" gorm:"column:status;comment:状态(1为正常,0为删除)"`
	IsDefault  int8      `json:"isDefault" gorm:"column:is_default;comment:默认类型(0:商家自定义,1:系统默认)"`
	ProdCount  int64     `json:"prodCount" gorm:"column:prod_count;comment:商品数量"`
	Style      int       `json:"style" gorm:"column:style;comment:列表样式(0:一列一个,1:一列两个,2:一列三个)"`
	Seq        int       `json:"seq" gorm:"column:seq;comment:排序"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
	DeleteTime time.Time `json:"deleteTime" gorm:"column:delete_time"`
}

// TableName 指定表名
func (n ProdTag) TableName() string {
	return "tz_prod_tag"
}
