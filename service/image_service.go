package service

import (
	"fmt"
	"github.com/tosashimanto/iot-camera-web/controller/dto"
	"github.com/tosashimanto/iot-camera-web/db/gorm"
	"github.com/tosashimanto/iot-camera-web/db/orm"
	"github.com/tosashimanto/iot-camera-web/model/model_db"
)

func InsertImage(dto *dto.LocationDto) error {

	db := gorm.DBManager()

	fmt.Println(db.Value)

	image_manage := model_db.Image_manage{}
	image_manage.ImageData = dto.ImageData
	image_manage.ExifLat = dto.ExifLat
	image_manage.ExifLon = dto.ExifLon
	image_manage.ExifAlt = 0.0
	image_manage.UpdatedBy = 99999999
	image_manage.CreatedBy = 99999999
	err := orm.Create(&image_manage)
	if err != nil {
		fmt.Println("Error=", err)
		return err
	}
	orm.FindAll(&image_manage)
	fmt.Println("TEST=", &image_manage)

	var count = 0
	db.Debug().Find(&image_manage).Count(&count)
	if count == 0 {
		fmt.Println("該当レコードなし")
	} else {
		fmt.Println("RegisteredDate:" + (image_manage.CreatedAt).String())
	}
	return err
}

func GetImageList() []model_db.Image_manage {

	image_manage_array := []model_db.Image_manage{}

	orm.FindAll(&image_manage_array)
	fmt.Println("GetImageList len=", len(image_manage_array))
	fmt.Println("GetImageList=", &image_manage_array)
	//for i := 0; i < len(image_manage_array); i++ {
	//	util.CreateImageFile("test001" + strconv.Itoa(i) + ".jpg", image_manage_array[i].ImageData)
	//}
	return image_manage_array
}
