package response

import "finalassignment/variable"

type RegisResp struct {
	Age   int    `json:"age"`
	Email string `json:"email"`
	variable.Id
	Username string `json:"username"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type UpdateUserResp struct {
	variable.Id
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	variable.Updated
}

type DelUserResp struct {
	Message string `json:"message"`
}
