package request

type RegisReq struct {
	Age      int    `json:"age" example:"12"`
	Email    string `json:"email" example:"admin@admin.com"`
	Password string `json:"password" example:"admin1"`
	Username string `json:"username" example:"admin"`
}

type LoginReq struct {
	Email    string `json:"email" example:"admin@admin.com"`
	Password string `json:"password" example:"admin1"`
}

type UpdateUserReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
