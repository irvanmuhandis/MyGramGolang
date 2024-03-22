package request

type PostCommentReq struct {
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
}

type UpdateCommentReq struct {
	Message string `json:"message"`
}
