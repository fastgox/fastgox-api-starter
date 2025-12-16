package handle

import (
	"net/http"
	"strconv"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

// CheckBlacklist 检查黑名单
// @Summary 检查用户三要素是否在黑名单
// @Description 检查姓名、身份证号、手机号是否命中黑名单
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "姓名"
// @Param id_card query string false "身份证号"
// @Param phone query string false "手机号"
// @Success 200 {object} dto.Response{data=services.CheckResult}
// @Router /blacklist/check [get]
func CheckBlacklist(c *gin.Context) {
	name := c.Query("name")
	idCard := c.Query("id_card")
	phone := c.Query("phone")

	result, err := services.Blacklist.Check(name, idCard, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "检查失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}

// AddBlacklist 添加黑名单
// @Summary 添加黑名单记录
// @Description 添加用户三要素到黑名单，同一用户可多次拉黑
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.AddBlacklistRequest true "黑名单信息"
// @Success 200 {object} dto.Response
// @Router /blacklist [post]
func AddBlacklist(c *gin.Context) {
	var req services.AddBlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	if err := services.Blacklist.Add(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "添加成功",
	})
}

// RemoveBlacklist 解除黑名单
// @Summary 解除用户黑名单
// @Description 解除用户所有拉黑记录
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query int false "用户ID"
// @Param id_card query string false "身份证号"
// @Success 200 {object} dto.Response
// @Router /blacklist [delete]
func RemoveBlacklist(c *gin.Context) {
	userIDStr := c.Query("user_id")
	idCard := c.Query("id_card")

	var err error
	if userIDStr != "" {
		userID, _ := strconv.ParseInt(userIDStr, 10, 64)
		err = services.Blacklist.Remove(userID)
	} else if idCard != "" {
		err = services.Blacklist.RemoveByIDCard(idCard)
	} else {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "请提供user_id或id_card参数",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "解除失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "解除成功",
	})
}

// GetBlacklistDetail 获取黑名单用户详情
// @Summary 获取黑名单用户详情
// @Description 获取用户信息及所有拉黑记录
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query int true "用户ID"
// @Success 200 {object} dto.Response{data=services.BlacklistUserDetail}
// @Router /blacklist/detail [get]
func GetBlacklistDetail(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	detail, err := services.Blacklist.GetUserWithRecords(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "success",
		Data:    detail,
	})
}

// ListBlacklist 查询黑名单列表
// @Summary 分页查询黑名单用户
// @Description 分页查询黑名单用户列表
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param keyword query string false "搜索关键词"
// @Param status query int false "状态: 1-黑名单中 0-已解除" default(1)
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(20)
// @Success 200 {object} dto.Response
// @Router /blacklist/list [get]
func ListBlacklist(c *gin.Context) {
	keyword := c.Query("keyword")
	statusStr := c.DefaultQuery("status", "1")
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "20")

	status, _ := strconv.Atoi(statusStr)
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	list, total, err := services.Blacklist.List(page, size, keyword, int8(status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "查询失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "success",
		Data: gin.H{
			"list":  list,
			"total": total,
			"page":  page,
			"size":  size,
		},
	})
}

// BatchAddBlacklist 批量添加黑名单
// @Summary 批量添加黑名单
// @Description 批量添加多条黑名单记录
// @Tags 黑名单管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body []services.AddBlacklistRequest true "黑名单列表"
// @Success 200 {object} dto.Response
// @Router /blacklist/batch [post]
func BatchAddBlacklist(c *gin.Context) {
	var items []services.AddBlacklistRequest
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	successCount, errors := services.Blacklist.BatchAdd(items)

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "批量添加完成",
		Data: gin.H{
			"success_count": successCount,
			"fail_count":    len(errors),
			"errors":        errors,
		},
	})
}

func init() {
	router.AuthRouter.GET("/blacklist/check", CheckBlacklist)
	router.AuthRouter.GET("/blacklist/list", ListBlacklist)
	router.AuthRouter.GET("/blacklist/detail", GetBlacklistDetail)
	router.AuthRouter.POST("/blacklist", AddBlacklist)
	router.AuthRouter.POST("/blacklist/batch", BatchAddBlacklist)
	router.AuthRouter.DELETE("/blacklist", RemoveBlacklist)
}
