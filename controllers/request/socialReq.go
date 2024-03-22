package request

type PostSocialReq struct {
	Name      string `json:"name"`
	SocailUrl string `json:"social_media_url"`
}

type UpdateSocialReq struct {
	Name      string `json:"name"`
	SocailUrl string `json:"social_media_url"`
}
