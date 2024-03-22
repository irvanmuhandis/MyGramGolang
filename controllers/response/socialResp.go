package response

import "finalassignment/variable"

type PostSocialResp struct {
	variable.Id
	Name      string `json:"name"`
	SocialUrl string `json:"socail_media_url"`
	variable.UserId
	variable.Created
}

type GetSocialResp struct {
	SocialMedia []SocialMediaItem `json:"social_medias"`
}

type SocialMediaItem struct {
	variable.Id
	Name      string `json:"name"`
	SocialUrl string `json:"socail_media_url"`
	variable.UserId
	variable.Created
	variable.Updated
	User GetUserSocial
}

type GetUserSocial struct {
	variable.Id
	Username string `json:"username"`
	Profile  string `json:"profile_image_url"`
}

type UpdateSocialResp struct {
	variable.Id
	Name      string `json:"name"`
	SocialUrl string `json:"socail_media_url"`
	variable.UserId
	variable.Updated
}

type DelSocialResp struct {
	Message string `json:"message"`
}
