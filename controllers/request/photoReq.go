package request

type PostPhotoReq struct {
	Title    string `json:"title" example:"Foto Bagus"`
	Caption  string `json:"caption" example:"Estetik"`
	PhotoUrl string `json:"photo_url" example:"https://image.gambarpng.id/pngs/gambar-transparent-cute-cup-doodle-line-art-black-white-vector_46835.png"`
}

type UpdatePhotoReq struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}
