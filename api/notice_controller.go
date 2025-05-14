package api

import (
	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListNotice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	query := db.Model(&models.Notice{})
	var res []models.Notice
	err := query.Order("publish_time desc").Where("status = ?", 1).Where("is_top", 1).Find(&res).Error
	if err != nil {
		c.JSON(500, Response{
			Code:    "99999",
			Msg:     "获取列表失败",
			Fail:    true,
			Success: false,
		})
		return
	}

	c.JSON(200, SuccessResponse(res))

}
