package domain

type UserRegisterResp struct {
	UserId   int32  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Theme    string `json:"theme"`
	UseDay   int32  `json:"use_day"`
	Token    string `json:"token"`
}
type UserLoginResp struct {
	UserId      int32  `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Theme       string `json:"theme"`
	ChatCount   int32  `json:"chat_count"`
	MemoirCount int32  `json:"memoir_count"`
	UseDay      int32  `json:"use_day"`
	Token       string `json:"token"`
}
type UserGetInfoResp struct {
	UserId      int32  `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Theme       string `json:"theme"`
	ChatCount   int32  `json:"chat_count"`
	MemoirCount int32  `json:"memoir_count"`
	UseDay      int32  `json:"use_day"`
}
type UserUpdateResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type UserChangePasswordResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type UserLogoutResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
