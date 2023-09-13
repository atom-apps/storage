package dto

type UploadResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	RealName  string `json:"real_name"`
	Mime      string `json:"mime"`
	Ext       string `json:"ext"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}
