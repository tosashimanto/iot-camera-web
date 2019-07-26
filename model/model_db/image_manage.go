package model_db

import "time"

type Image_manage struct {
	ID           uint64    `gorm:"primary_key;column:id"`
	ImageData    []byte    `gorm:"column:image_data" sql:"type:bytea"`                   // 画像バイナリー
	ExifLat      float64   `gorm:"column:exif_lat" sql:"not null;type:double precision"` // GPS緯度
	ExifLon      float64   `gorm:"column:exif_lon" sql:"not null;type:double precision"` // GPS経度
	ExifAlt      float64   `gorm:"column:exif_alt" sql:"not null;type:double precision"` // GPS高度
	ExifDateTime time.Time `gorm:"column:exif_date_time" sql:"not null;type:timestamp"`  // 撮影日時
	CreatedBy    uint64    `gorm:"column:created_by;" sql:"not null;type:bigint"`
	CreatedAt    time.Time `gorm:"column:created_at;" sql:"not null;type:timestamp"`
	UpdatedBy    uint64    `gorm:"column:updated_by;" sql:"not null;type:bigint"`
	UpdatedAt    time.Time `gorm:"column:updated_at;" sql:"not null;type:timestamp"`
}
