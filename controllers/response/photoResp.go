package response

import "finalassignment/variable"

type PostPhotoResp struct {
	variable.Id
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	variable.UserId
	variable.Created
}

type GetPhotoResp struct {
	variable.Id
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	variable.UserId
	variable.Created
	variable.Updated
	User GetUserItemPhoto `json:"User"`
}

type GetUserItemPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhotoResp struct {
	variable.Id
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	variable.UserId
	variable.Updated
}

type DelPhotoResp struct {
	Message string `json:"message"`
}
