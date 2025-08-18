package router

import (
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.UserSvc.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "获取用户失败",
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "获取成功",
		Data:    user,
	})
}

// UpdateUser 更新用户信息（需要认证）
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id") // 从认证中间件获取用户ID

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "更新成功",
		Data: gin.H{
			"id":      id,
			"user_id": userID,
		},
	})
}

// 使用 init 函数自动注册路由
func init() {
	// 直接使用路由组注册路由
	router.PublicRouter.GET("/users/:id", GetUser)
	router.AuthRouter.PUT("/users/:id", UpdateUser)
}
