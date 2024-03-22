package response

import "finalassignment/variable"

type PostCommentResp struct {
	variable.Id
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
	variable.UserId
	variable.Created
}

type GetCommentResp struct {
	variable.Id
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
	variable.UserId
	variable.Updated
	variable.Created
	User  GetUserItemComment
	Photo GetPhotoItemComment
}
type GetUserItemComment struct {
	variable.Id
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetPhotoItemComment struct {
	variable.Id
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	variable.UserId
}

type UpdateCommentResp struct {
	variable.Id
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	variable.UserId
	variable.Updated
}

type DelCommentResp struct {
	Message string `json:"message"`
}
