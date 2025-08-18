package request

// 基础分页请求
type PageRequest struct {
	Page int `json:"page" form:"page" binding:"min=1"`
	Size int `json:"size" form:"size" binding:"min=1,max=100"`
}
