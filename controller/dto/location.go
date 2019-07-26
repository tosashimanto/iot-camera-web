package dto

import "time"

type LocationDto struct {
	// 画像バイナリー
	ImageData []byte
	// GPS緯度
	ExifLat float64
	// GPS経度
	ExifLon float64
	// GPS高度
	ExifAlt float64
	// 撮影日時
	ExifDateTime time.Time
}
