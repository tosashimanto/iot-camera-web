package api

type PutUploadURLResponse struct {
	ImageUplaodURLArray []ImageUploadURL `json:"imageUplaodURLArray"`
}

type ImageUploadURL struct {
	Index     int    `json:"index"`     // Index
	UploadUrl string `json:"uploadUrl"` // アップロードURL
}

type GetImageListResponse struct {
	ImageManegeArray []Image_manage `json:"imageManageArray"`
}

type Image_manage struct {
	ID           uint64  `json:"ID"`
	ImageData    string  `json:"ImageData"`
	ExifLat      float64 `json:"ExifLat"`
	ExifLon      float64 `json:"ExifLon"`
	ExifAlt      float64 `json:"ExifAlt"`
	ExifDateTime string  `json:"ExifDateTime"`
	CreatedBy    uint64  `json:"CreatedBy"`
	CreatedAt    string  `json:"CreatedAt"`
	UpdatedBy    uint64  `json:"UpdatedBy"`
	UpdatedAt    string  `json:"UpdatedAt"`
}
