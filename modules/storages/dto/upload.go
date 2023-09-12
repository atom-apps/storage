package dto

type UploadResponse struct {
	Name      string `json:"name,omitempty"`
	Mime      string `json:"mime,omitempty"`
	Ext       string `json:"ext,omitempty"`
	URL       string `json:"url,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
