package response

// UserResponse 用户响应
type UserResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	ProfileImageURL string  `json:"profile_image_url"`
	Phone           *string `json:"phone,omitempty"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
	LastActiveAt    int64   `json:"last_active_at"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expires_at"`
	User      *UserResponse `json:"user,omitempty"`
}
