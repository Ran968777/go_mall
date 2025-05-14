package api

import (
	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListIndexImgRequest 列表请求参数
type ListIndexImgRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	ShopID   uint64 `form:"shop_id"`
	Status   *int8  `form:"status"`
	Type     *int   `form:"type"`
}

// ListIndexImgResponse 列表响应
type ListIndexImgResponse struct {
	Total int64             `json:"total"`
	List  []models.IndexImg `json:"list"`
}

// ListIndexImg 获取轮播图列表
func ListIndexImg(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	query := db.Model(&models.IndexImg{})

	// 获取列表
	var list []models.IndexImg
	if err := query.Order("seq desc, upload_time desc").
		Find(&list).Error; err != nil {
		c.JSON(500, Response{
			Code:    "99999",
			Msg:     "获取列表失败",
			Fail:    true,
			Success: false,
		})
		return
	}

	// 遍历并处理每个图片URL
	for i := range list {
		if list[i].ImgURL != "" {
			list[i].ImgURL = "https://img.mall4j.com/" + list[i].ImgURL
		}
	}

	c.JSON(200, SuccessResponse(list))
}
