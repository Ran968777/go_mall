package api

import (
	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTag(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	tx := db.Model(&models.ProdTag{})

	var res []models.ProdTag
	err := tx.Order("seq desc").Where("status = ?", 1).Find(&res).Error
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
