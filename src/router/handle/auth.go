package handle

import (
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

// Login 登录
// @Summary 用户登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "登录请求"
// @Success 200 {object} dto.Response{data=response.LoginResponse}
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 400, Message: err.Error()})
		return
	}
	result, err := services.AuthSvc.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{Code: 401, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: 200, Message: "登录成功", Data: result})
}

func init() {
	router.PublicRouter.POST("/auth/login", Login)
}
