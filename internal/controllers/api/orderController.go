package api

import (
	"cyclopropane/internal/global"
	"cyclopropane/internal/models/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddOrder 创建订单
func AddOrder(c *gin.Context) {
	var o order.Order

	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检验订单
	if !o.CheckOrder() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "订单信息错误"})
		return
	}

	// TODO 将获取用户 email 的方法抽到一个统一的文件中
	id, _ := c.Get("email")
	creatorId, _ := id.(string)
	o.CreatorID = creatorId

	err := addOrder(o)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"creator_id": creatorId,
	})

	return
}

func addOrder(o order.Order) error {
	result := global.DB.Create(&o)
	if result.Error != nil {
		global.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil
}